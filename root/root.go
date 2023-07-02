package root

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/debug"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/prof"
	"gitlab.com/gnufred/logslate/report"
	"gitlab.com/gnufred/logslate/threat"
	"gitlab.com/gnufred/logslate/translate"

	progbar "github.com/schollz/progressbar/v3"
	nginx "gitlab.com/gnufred/nxparse"
)

type Parsed struct {
	Ok   uint
	Skip uint
	Err  uint
	Atk  uint
	Bot  uint
}

var (
	ErrNoInputData = errors.New("no input data")
	ErrReadPipe    = errors.New("error while reading piped io.Stdin")
	ErrLimitReach  = errors.New("parsing error limit reached")
)

func ParseStdin(a args.Args) (Parsed, error) {
	pi := pipe.New()
	par := Parsed{}

	if !pi.IsPiped() {
		return par, ErrNoInputData
	}

	d := db.Singleton()
	if a.Db != "" {
		err := d.Create(a.Db)
		if err != nil {
			return par, err
		}
		d.ApplyDefaults()
	}

	if a.Db != "" {
		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	bar := createBar()

	// Reporting data starts here.
	runStartedAt := time.Now()

	// Start CPU profiling if argument is set.
	cpuf, err := prof.StartCpuProfile(a)
	if err != nil {
		return par, err
	}
	// The defer stop profiling in case the function exit before the profiling
	// stops when specifically requested.
	defer prof.StopCpuProfile(a, cpuf)

	// Initialize parser loop variables.
	var bytesTotal uint
	atkTypes := make(map[byte]uint, threat.XSS)
	attackers := make(map[string]uint)
	bots := make(map[string]uint)

	// Parser loop.
	for {
		if par.Ok < 1 && par.Err > 100 {
			return par, ErrLimitReach
		}

		line, byteCnt, err := pi.ReadLine()
		bytesTotal = bytesTotal + uint(byteCnt)
		bar.Add(byteCnt)

		if err == io.EOF {
			break
		}

		if err != nil {
			return par, fmt.Errorf("%s: %w", ErrReadPipe, err)
		}

		if line == "" {
			par.Err++
			continue
		}

		nx, err := nginx.Parse(a.LogFrmt, line)
		if err != nil {
			par.Err++
			debug.Log(debug.VVV, "parseError", line)
			continue
		}

		ls, err := translate.Nginx(nx)
		if err != nil {
			par.Err++
			debug.Log(debug.VVV, "translateError", err)
			// For now there's a single error that can bubble up and we can safely
			// ignore it. Counting and logging is fine for now.
		}

		if !timeMatch(ls.Time, bool(a.Since() != ""), bool(a.Until() != ""), a.StartTime(), a.EndTime()) {
			par.Skip++
			continue
		}

		if a.Db != "" {
			d.InsertLog(ls)
			// Commit every "x" successfully parsed lines.
			// Testing with db/BenchmarkDbInsert shows that "10" is the best number of inserts to batch.
			if par.Ok%10 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		if ls.Attack.Score > 0 {
			debug.Log(debug.VVV, "flagged", ls.Attack.Name+" | "+ls.Path+ls.Query)

			// Compile attack type data.
			if a.ReportPath != "" {
				atkTypes[ls.Attack.Id]++
				attackers[ls.Ip.String()]++
			}
			par.Atk++
		}

		if ls.Bot != "" {
			if a.ReportPath != "" {
				bots[ls.Bot]++
			}
			par.Bot++
		}

		par.Ok++
	}

	// Need to print a return character after the last progress bar terminal update.
	fmt.Print("\n")

	// Print parsed line data.
	fmt.Printf("Parsed: ok [%d] skip [%d] error [%d]\n", par.Ok, par.Skip, par.Err)
	fmt.Printf("Threats: attack [%d], bots [%d]\n", par.Atk, par.Bot)

	// Stop CPU profiling.
	prof.StopCpuProfile(a, cpuf)

	// Write memory heap if argument is set.
	err = prof.MemProfile(a)
	if err != nil {
		return par, err
	}

	// Reporting data stops here.
	runEndedAt := time.Now()

	// Process report data if requested.
	if a.ReportPath != "" {
		// Aggregate report data.
		rep := report.New()
		rep.Path = a.ReportPath
		rep.Frmt = a.ReportFrmt
		rep.Run.Start = runStartedAt
		rep.Run.End = runEndedAt
		rep.Size.Byte = bytesTotal
		rep.Line.Ok = par.Ok
		rep.Line.Skip = par.Skip
		rep.Line.Error = par.Err
		rep.Time.Since = a.Since()
		rep.Time.Until = a.Until()
		rep.Time.Start = a.StartTime()
		rep.Time.End = a.EndTime()
		rep.Db.Path = a.Db

		rep.Thrt.Atks.Its = par.Atk
		rep.Thrt.Atks.Prct = report.Prct(par.Atk, par.Ok)
		rep.Thrt.Bots.Its = par.Bot
		rep.Thrt.Bots.Prct = report.Prct(par.Bot, par.Ok)

		rep.AtkTy = report.AtkTypes{
			Unk: atkTypes[threat.Unk],
			SQL: atkTypes[threat.SQL],
			CMD: atkTypes[threat.CMD],
			Dir: atkTypes[threat.Dir],
			XSS: atkTypes[threat.XSS],
		}

		for ip, cnt := range attackers {
			rep.Atkr[ip] = report.Atkrs{
				Cnt: cnt,
			}
		}

		for name, cnt := range bots {
			rep.Bot[name] = report.Bots{
				Cnt: cnt,
			}
		}

		// Print report.
		rep.Print()
	}

	return par, nil
}

func createBar() *progbar.ProgressBar {
	bar := progbar.NewOptions(-1,
		progbar.OptionSetDescription("parsing"),
		progbar.OptionSetWriter(os.Stdout),
		progbar.OptionShowBytes(true),
		progbar.OptionSetWidth(10),
		progbar.OptionThrottle(166*time.Millisecond),
		progbar.OptionShowCount(),
		progbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stdout, "\n")
		}),
		progbar.OptionSpinnerType(14),
		progbar.OptionSetRenderBlankState(true),
	)
	return bar
}

func timeMatch(ts time.Time, since bool, until bool, start time.Time, end time.Time) bool {
	// First possibility: no arguments.
	if !since && !until {
		return true
	}

	// Second possibility: since is specified.
	if since && !until {
		if ts.Equal(start) || ts.After(start) {
			return true
		}
	}

	// Third possibility: until is specified.
	if !since && until {
		if ts.Before(end) || ts.Equal(end) {
			return true
		}
	}

	// Fourth possibility: both are specified.
	if since && until {
		if (ts.Equal(start) || ts.After(start)) &&
			(ts.Before(end) || ts.Equal(end)) {
			return true
		}
	}

	return false
}
