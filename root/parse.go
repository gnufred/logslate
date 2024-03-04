package root

import (
	"fmt"
	"io"
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

func ParseStdin(a *args.ArgsV2, version string) (Parsed, error) {
	pi := pipe.New()
	par := Parsed{}

	// Initialize output processing variables.
	outTo := initOutput()

	// Banner.
	if outTo.Progress {
		fmt.Printf("Logslate %s @ https://gitlab.com/gnufred/logslate\n", version)
	}

	if !pi.IsPiped() {
		return par, ErrNoInputData
	}

	// Start CPU profiling if argument is set.
	cpuf, err := prof.StartCpuProfile(a.Debug.CpuProf)
	if err != nil {
		return par, err
	}
	// The defer stop profiling in case the function exit before the profiling
	// stops when specifically requested.
	defer prof.StopCpuProfile(a.Debug.CpuProf, cpuf)

	// Initialize database.
	d := db.Singleton()
	if outTo.DB {
		if err := d.Create(a.Parse.Db); err != nil {
			return par, err
		}
		d.ApplyDefaults()

		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	// Initialize progress bar.
	var bytesTotal uint
	var bar *progbar.ProgressBar
	if outTo.Progress {
		bar = createBar(a.Command)
	}

	// Initialize the threat log.
	var tlPrinter printer.Printer
	if outTo.ThreatLog {
		tlPrinter = printer.Printer{
			Path:        a.Threat.Path,
			Format:      a.Threat.Format,
			NewLine:     "\n",
			Stringifier: log.StringifyThreatLog,
		}
		if a.Threat.Format != args.SCREEN {
			if err := tlPrinter.OpenFile(); err != nil {
				return par, err
			}
			defer tlPrinter.CloseFile()
		}
	}

	// Initialize reporting.
	var rep *report.Report
	var repTL report.ReportTL
	var repPrinter printer.Printer
	if outTo.Report {
		repTL = report.NewTL()
		rep = report.New(
			a.Report.Layout,
		)
		repPrinter = printer.Printer{
			Path:        a.Report.Path,
			Format:      a.Report.Format,
			Stringifier: report.Stringify,
		}

		if a.Report.Format != args.SCREEN {
			if err := repPrinter.OpenFile(); err != nil {
				return par, err
			}
			defer repPrinter.CloseFile()
		}
	}

	// Runtime counter start here.
	runStartedAt := time.Now()

	// Parser loop.
	for {
		if par.Ok < 1 && par.Err > 100 {
			return par, ErrLimitReach
		}

		line, byteCnt, err := pi.ReadLine()
		bytesTotal = bytesTotal + uint(byteCnt)
		if outTo.Progress {
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

		var nx nginx.Entry
		if a.Command == args.PARSE {
			nx, err = nginx.Parse(a.Parse.LogFrmt, line)
			if err != nil {
				par.Err++
				debug.Log(debug.VVV, "parseError", line)
				continue
			}
		}

		var sl log.SlateLog
		if a.Command == args.PARSE {
			sl, err = log.NginxToSlate(nx)
			if err != nil {
				par.Err++
				debug.Log(debug.VVV, "translateError", err)
				// For now there's a single error that can bubble up and we can safely
				// ignore it. Counting and logging is fine for now.
			}
		}

		var tl log.ThreatLog
		if a.Command == args.PARSE {
			tl = log.SlateToThreat(sl)
		}

		if a.Command == args.REPARSE {
			tl, err = log.ParseThreat(line)
			if err != nil {
				par.Err++
				debug.Log(debug.VVV, "unmarshalError", line)
				continue
			}

		}

		if !withinRange(tl.Time, a.Time.StartTime(), a.Time.EndTime()) {
			par.Skip++
			continue
		}

		if outTo.DB {
			if err = d.Insert(sl); err != nil {
				return par, err
			}
			// Commit every "x" successfully parsed lines.
			// Testing with db/BenchmarkDbInsert shows that "10" is the best number of inserts to batch.
			if par.Ok%10 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		if outTo.ThreatLog {
			tlPrinter.Print(tl)
		}

		if tl.IsThreat() {
			if outTo.Report {
				repTL.Parse(tl)
			}

			if tl.IsAttack() {
				par.Atk++
				debug.Log(debug.VVV, "flagged", tl.Attack.Name+" | "+tl.Path+tl.Query)
			}

			if tl.IsAd() {
				par.Ad++
			}

			if tl.IsBot() {
				par.Bot++
			}
		}

		par.Ok++

		if outTo.Report {
			rep.Parsed.ParseFirstLast(tl.Time)
		}
	}

	// Runtime counter stop here.
	runEndedAt := time.Now()

	if outTo.Progress {
		// Need to print a return character after the last progress bar terminal update.
		fmt.Print("\n")

		// Print parsed line data.
		fmt.Printf("Entries: ok [%d] skip [%d] error [%d]\n", par.Ok, par.Skip, par.Err)
		fmt.Printf("Threats: attack [%d], ad [%d], bots [%d]\n", par.Atk, par.Ad, par.Bot)
	}

	// Process report data if requested.
	if outTo.Report {
		rep.ThreatLog = repTL

		rep.Runtime.Start = runStartedAt
		rep.Runtime.End = runEndedAt
		rep.Size.Byte = bytesTotal

		rep.Parsed.Ok = par.Ok
		rep.Parsed.Skip = par.Skip
		rep.Parsed.Error = par.Err

		// If the time entries didn't change, we set them to
		// what the user specified.
		rep.TimeArgs.Since = a.Time.Since()
		rep.TimeArgs.Until = a.Time.Until()
		rep.TimeArgs.Start = a.Time.StartTime()
		rep.TimeArgs.End = a.Time.EndTime()

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

	// Write memory heap if argument is set.
	if err := prof.MemProfile(a.Debug.MemProf); err != nil {
		return par, err
	}

	return par, nil
}
