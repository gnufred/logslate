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
	"gitlab.com/gnufred/logslate/log"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/printer"
	"gitlab.com/gnufred/logslate/prof"
	"gitlab.com/gnufred/logslate/report"
	"gitlab.com/gnufred/logslate/threat"

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

	// Initialize progress bar.
	var bytesTotal uint
	showProgress := true
	var bar *progbar.ProgressBar

	if showProgress {
		fmt.Printf("Logslate %s @ https://gitlab.com/gnufred/logslate\n", a.Version)
	}

	if !pi.IsPiped() {
		return par, ErrNoInputData
	}

	// Start CPU profiling if argument is set.
	cpuf, err := prof.StartCpuProfile(a)
	if err != nil {
		return par, err
	}
	// The defer stop profiling in case the function exit before the profiling
	// stops when specifically requested.
	defer prof.StopCpuProfile(a, cpuf)

	// Initialize database.
	d := db.Singleton()
	outputToDB := false
	if a.Db != "" {
		outputToDB = true
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

	// Initialize debug output.
	outputToDebug := false
	if a.Debug >= int(debug.VVV) {
		outputToDebug = true
	}

	// Initialize the threat log.
	outputToThreatLog := false
	var tlPrinter printer.Printer
	if a.ThreatPath != "" {
		outputToThreatLog = true
		tlPrinter = printer.Printer{
			Path:        a.ThreatPath,
			Format:      a.ThreatFormat,
			NewLine:     "\n",
			Stringifier: log.StringifyThreatLog,
		}
		if a.ThreatFormat == args.SCREEN {
			showProgress = false
		} else {
			err := tlPrinter.OpenFile()
			if err != nil {
				return par, err
			}
			defer tlPrinter.CloseFile()
		}
	}

	// Initialize reporting.
	outputToReport := false
	var rep *report.Report
	var repTL report.ReportTL
	var repPrinter printer.Printer
	if a.ReportPath != "" {
		outputToReport = true
		repTL = report.NewTL()
		rep = report.New(
			a.ReportLayout,
		)
		repPrinter = printer.Printer{
			Path:        a.ReportPath,
			Format:      a.ReportFormat,
			Stringifier: report.Stringify,
		}

		if a.ReportFormat != args.SCREEN {
			err := repPrinter.OpenFile()
			if err != nil {
				return par, err
			}
			defer repPrinter.CloseFile()
		}
	}

	// Runtime counter start here.
	runStartedAt := time.Now()

	if showProgress {
		bar = createBar()
	}

	// Parser loop.
	for {
		if par.Ok < 1 && par.Err > 100 {
			return par, ErrLimitReach
		}

		line, byteCnt, err := pi.ReadLine()
		bytesTotal = bytesTotal + uint(byteCnt)
		if showProgress {
			bar.Add(byteCnt)
		}

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

		sl, err := log.NginxToSlate(nx)
		if err != nil {
			par.Err++
			debug.Log(debug.VVV, "translateError", err)
			// For now there's a single error that can bubble up and we can safely
			// ignore it. Counting and logging is fine for now.
		}

		if !timeMatch(sl.Time, bool(a.Since() != ""), bool(a.Until() != ""), a.StartTime(), a.EndTime()) {
			par.Skip++
			continue
		}

		if sl.IsAttack() {
			par.Atk++
			debug.Log(debug.VVV, "flagged", sl.Attack.Name+" | "+sl.Path+sl.Query)
		}

		if sl.IsAd() {
			par.Ad++
		}

		if sl.IsBot() {
			par.Bot++
		}

		par.Ok++

		// === OUTPUT SECTION == //

		if outputToDB {
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

		if sl.IsThreat() {
			line := log.SlateToThreat(sl)

			if outputToThreatLog {
				tlPrinter.Print(line)
			}

			if outputToReport {
				repTL.Parse(line)
			}

			if outputToDebug {
				if sl.IsAttack() {
					par.Atk++
					debug.Log(debug.VVV, "flagged", line.Path+" "+line.Query)
				}
			}
		}

		if outputToReport {
			rep.Parsed.ParseFirstLast(sl.Time)
		}
	}

	// Runtime counter stop here.
	runEndedAt := time.Now()

	if showProgress {
		// Need to print a return character after the last progress bar terminal update.
		fmt.Print("\n")

		// Print parsed line data.
		fmt.Printf("Parsed: ok [%d] skip [%d] error [%d]\n", par.Ok, par.Skip, par.Err)
		fmt.Printf("Threats: attack [%d], ad [%d], bots [%d]\n", par.Atk, par.Ad, par.Bot)
	}

	// Process report data if requested.
	if outputToReport {
		rep.ThreatLog = repTL

		rep.Runtime.Start = runStartedAt
		rep.Runtime.End = runEndedAt
		rep.Size.Byte = bytesTotal

		rep.Parsed.Ok = par.Ok
		rep.Parsed.Skip = par.Skip
		rep.Parsed.Error = par.Err

		// If the time entries didn't change, we set them to
		// what the user specified.
		rep.TimeArgs.Since = a.Since()
		rep.TimeArgs.Until = a.Until()
		rep.TimeArgs.Start = a.StartTime()
		rep.TimeArgs.End = a.EndTime()

		// Compile advertising data.
		rep.Threats.Ads.Its = par.Ad
		rep.Threats.Ads.Prct = report.Prct(par.Ad, par.Ok)

		for name, cnt := range rep.ThreatLog.Ads {
			rep.Ads[name] = report.ListRD{
				Cnt: cnt,
			}
		}

		// Compile bot data.
		rep.Threats.Bots.Its = par.Bot
		rep.Threats.Bots.Prct = report.Prct(par.Bot, par.Ok)

		for name, cnt := range rep.ThreatLog.Bots {
			rep.Bots[name] = report.ListRD{
				Cnt: cnt,
			}
		}

		// Compile attack data.
		rep.Threats.Attacks.Its = par.Atk
		rep.Threats.Attacks.Prct = report.Prct(par.Atk, par.Ok)

		rep.Attacks = report.AttacksRD{
			UNK: rep.ThreatLog.Attacks[threat.UNK],
			SQL: rep.ThreatLog.Attacks[threat.SQL],
			CMD: rep.ThreatLog.Attacks[threat.CMD],
			DIR: rep.ThreatLog.Attacks[threat.DIR],
			XSS: rep.ThreatLog.Attacks[threat.XSS],
			SCN: rep.ThreatLog.Attacks[threat.SCN],
		}

		for ip, cnt := range rep.ThreatLog.Attackers {
			rep.Attackers[ip] = report.ListRD{
				Cnt: cnt,
			}
		}

		repPrinter.Print(*rep)
	}

	// Stop CPU profiling.
	prof.StopCpuProfile(a, cpuf)

	// Write memory heap if argument is set.
	err = prof.MemProfile(a)
	if err != nil {
		return par, err
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
