package root

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
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
	Ad   uint
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

	// Initialize database
	d := db.Singleton()
	if a.Db != "" {
		err := d.Create(a.Db)
		if err != nil {
			return par, err
		}
		d.ApplyDefaults()

		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	// Create runtime progress bar
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
	// Initialize the threat log.
	var tl []translate.ThreatLog
	// Initialize map length with highest attack type constant.
	atkTypes := make(map[threat.AttackType]uint, threat.XSS)
	attackers := make(map[string]uint)
	attacks := make([]report.AttackRD, 0)
	ads := make(map[string]uint)
	bots := make(map[string]uint)
	// "Zero" values for first/last entry must be, infinitely in the future/past.
	first, _ := time.Parse("2006-01-02", "9999-01-02")
	last := time.Time{}

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

		sl, err := translate.NginxToSlate(nx)
		if err != nil {
			par.Err++
			debug.Log(debug.VVV, "translateError", err)
			// For now there's a single error that can bubble up and we can safely
			// ignore it. Counting and logging is fine for now.
		}
		tl = append(tl, translate.SlateToThreat(sl))
		z := len(tl) - 1

		if !timeMatch(sl.Time, bool(a.Since() != ""), bool(a.Until() != ""), a.StartTime(), a.EndTime()) {
			par.Skip++
			continue
		}

		if a.Db != "" {
			err = d.Insert(sl)
			if err != nil {
				return par, err
			}
			// Commit every "x" successfully parsed lines.
			// Testing with db/BenchmarkDbInsert shows that "10" is the best number of inserts to batch.
			if par.Ok%10 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		if sl.Time.Before(first) {
			first = sl.Time
		}
		if sl.Time.After(last) {
			last = sl.Time
		}

		if tl[z].Attack != threat.Nil {
			par.Atk++

			debug.Log(debug.VVV, "flagged", sl.Attack.Name+" | "+sl.Path+sl.Query)

			// Report data
			ip := tl[z].Ip.String()
			atkTypes[tl[z].Attack]++
			attackers[ip]++
			attacks = append(attacks,
				report.AttackRD{
					Type:  threat.String(tl[z].Attack),
					Time:  tl[z].Time.Format("2006-01-02 15:04:05"),
					Ip:    ip,
					Path:  tl[z].Path,
					Query: tl[z].Query,
				},
			)
		}

		if tl[z].Ad != "" {
			par.Ad++
			ads[tl[z].Ad]++
		}

		if tl[z].Bot != "" {
			par.Bot++
			bots[tl[z].Bot]++
		}

		par.Ok++
	}

	// Need to print a return character after the last progress bar terminal update.
	fmt.Print("\n")

	// Print parsed line data.
	fmt.Printf("Parsed: ok [%d] skip [%d] error [%d]\n", par.Ok, par.Skip, par.Err)
	fmt.Printf("Threats: attack [%d], ad [%d], bots [%d]\n", par.Atk, par.Ad, par.Bot)

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
		rep := report.New(
			a.ReportPath,
			a.ReportFormat,
			a.ReportLayout,
		)

		rep.Runtime.Start = runStartedAt
		rep.Runtime.End = runEndedAt
		rep.Size.Byte = bytesTotal

		rep.Parsed.Ok = par.Ok
		rep.Parsed.Skip = par.Skip
		rep.Parsed.Error = par.Err
		// "Zero" value for report first/last is "0001-01-01 00:00:00 UTC"
		if first.Equal(a.EndTime()) || first.Before(a.EndTime()) {
			rep.Parsed.First = first
		}
		rep.Parsed.Last = last

		// If the time entries didn't change, we set them to
		// what the user specified.
		rep.TimeArgs.Since = a.Since()
		rep.TimeArgs.Until = a.Until()
		rep.TimeArgs.Start = a.StartTime()
		rep.TimeArgs.End = a.EndTime()

		rep.Threats.Attacks.Its = par.Atk
		rep.Threats.Attacks.Prct = report.Prct(par.Atk, par.Ok)
		rep.Threats.Ads.Its = par.Ad
		rep.Threats.Ads.Prct = report.Prct(par.Ad, par.Ok)
		rep.Threats.Bots.Its = par.Bot
		rep.Threats.Bots.Prct = report.Prct(par.Bot, par.Ok)

		rep.AttackTypes = report.AttackTypesRD{
			Unk: atkTypes[threat.Unk],
			SQL: atkTypes[threat.SQL],
			CMD: atkTypes[threat.CMD],
			Dir: atkTypes[threat.Dir],
			XSS: atkTypes[threat.XSS],
			Scn: atkTypes[threat.Scn],
		}

		for ip, cnt := range attackers {
			rep.Attackers[ip] = report.ListRD{
				Cnt: cnt,
			}
		}

		for name, cnt := range ads {
			rep.Ads[name] = report.ListRD{
				Cnt: cnt,
			}
		}

		for name, cnt := range bots {
			rep.Bots[name] = report.ListRD{
				Cnt: cnt,
			}
		}

		sort.Slice(attacks, func(i, j int) bool {
			return attacks[i].Time < attacks[j].Time
		})
		rep.Attacks = attacks

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
