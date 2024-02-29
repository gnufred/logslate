package args

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
	nginx "gitlab.com/gnufred/nxparse"
)

type SubCmd byte
type LogFormat byte

const (
	NILSUB SubCmd = iota
	VERSION
	PARSE
	REPARSE

	SCREEN LogFormat = iota
	TEXT
	JSON
)

var ErrHelpShown = errors.New("help shown")

func Help() {
	fmt.Fprint(os.Stderr, `
USAGE
  logslate [SUBCOMMAND] [OPTIONS]

SUB COMMANDS
  p, parse (default)
    Parse NGINX access.log.

  r, reparse
    Parse LOGSLATE threat.json.

  v, version
    Print version.

PARSE
  -f, --log_format <NGINX LOG_FORMAT>
    Need to match NGINX log_format directive from your nginx.conf.
    Default to NGINX default combined format.

  -d, --db <FILE>
    Convert NGINX access.log to an SQLite3 DB.

PARSE, REPARSE
  -s, --since <TIME>
  -u, --until <TIME>
    Parse entries using specified date(s).
    Can use timestamps or system.d style wording.

  -t, --threats (screen|<FILE>)
    Output threat log. ".json" files will be formatted as such.

  -r, --report (screen|<FILE>)
    Output report. ".json" files will be formatted as such.

  -l, --layout <LAYOUT>
    Custom report layout for screen and text reports.

MORE
  For detailed usage instructions read:
  https://gitlab.com/gnufred/logslate/-/blob/main/.docs/usage.md
`)

}

func Init() error {
	if !pflag.Parsed() {
		if err := parse(); err != nil {
			return err
		}
		if err := validate(pflag.NArg()); err != nil {
			return err
		}
	}
	return nil
}

func validate(nargs int) error {
	a := Singleton()
	if a.Help {
		Help()
		return ErrHelpShown
	}

	// Return true if invalid.
	cManySubC := func(a *ArgsV2) bool { return nargs > 1 }
	cSubC := func(a *ArgsV2) bool { return a.Command == NILSUB }
	cDebug := func(a *ArgsV2) bool { return a.Debug.Verbosity < 0 || a.Debug.Verbosity > 3 }
	cLayout := func(a *ArgsV2) bool {
		match, _ := regexp.MatchString("[^a-z0-9:,\\_]", a.Report.Layout)
		return match
	}
	cReportFormat := func(a *ArgsV2) bool { return a.Report.Format == JSON && a.Report.Layout != "" }
	cReparseDb := func(a *ArgsV2) bool { return a.Parse.Db != "" && a.Command != PARSE }
	cReparseLogFmt := func(a *ArgsV2) bool { return a.Parse.LogFrmt != "" && a.Command != PARSE }

	checks := []struct {
		F func(a *ArgsV2) bool
		E string
	}{
		{cManySubC, "multiple sub commands"},
		{cSubC, "invalid sub command"},
		{cDebug, "debug args must be -v or --debug=[1,3]"},
		{cLayout, "invalid characters in layout format"},
		{cReportFormat, "report layout is only available for screen and text reports"},
		{cReparseDb, "cannot use -d, --db with reparse"},
		{cReparseLogFmt, "cannot use -f, --log_format with reparse"},
	}

	for _, c := range checks {
		if invalid := c.F(a); invalid {
			return errors.New(c.E)
		}
	}

	return nil
}

func parse() error {
	flags := pflag.CommandLine
	flags.SortFlags = false

	f := struct {
		Help         *bool
		CpuProf      *string
		MemProf      *string
		Db           *string
		LogFrmt      *string
		Since        *string
		Until        *string
		ThreatLog    *string
		Report       *string
		ReportLayout *string
	}{
		Help:         flags.BoolP("help", "h", false, ""),
		CpuProf:      flags.StringP("profcpu", "", "", ""),
		MemProf:      flags.StringP("profmem", "", "", ""),
		Db:           flags.StringP("db", "d", "", ""),
		LogFrmt:      flags.StringP("log_format", "f", nginx.Combined, ""),
		Since:        flags.StringP("since", "s", "", ""),
		Until:        flags.StringP("until", "u", "", ""),
		ThreatLog:    flags.StringP("threats", "t", "", ""),
		Report:       flags.StringP("report", "r", "", ""),
		ReportLayout: flags.StringP("layout", "l", "", ""),
	}

	flags.CountP("debug", "v", "")
	debugCnt, _ := flags.GetCount("debug")

	flags.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	flags.Lookup("profmem").NoOptDefVal = "mem.prof"

	flags.MarkHidden("help")
	flags.MarkHidden("debug")
	flags.MarkHidden("profcpu")
	flags.MarkHidden("profmem")

	pflag.Parse()
	arg0 := flags.Arg(0)

	argInst = &ArgsV2{
		Command: parseSubCmd(arg0),
		Help:    bool(*f.Help),
		Debug: debugArgs{
			Verbosity: debugCnt,
			CpuProf:   string(*f.CpuProf),
			MemProf:   string(*f.MemProf),
		},
		Parse: parseArgs{
			Db:      string(*f.Db),
			LogFrmt: string(*f.LogFrmt),
		},
		Threat: threatArgs{
			Path:   string(*f.ThreatLog),
			Format: parseLogFormat(string(*f.ThreatLog)),
		},
		Report: reportArgs{
			Path:   string(*f.Report),
			Format: parseLogFormat(string(*f.Report)),
			Layout: string(*f.ReportLayout),
		},
	}

	if err := argInst.Time.SetSince(string(*f.Since)); err != nil {
		return err
	}
	if err := argInst.Time.SetUntil(string(*f.Until)); err != nil {
		return err
	}
	return nil
}

func parseLogFormat(path string) LogFormat {
	lpath := strings.ToLower(path)
	if path == "screen" {
		return SCREEN
	} else if strings.HasSuffix(lpath, ".json") {
		return JSON
	}
	return TEXT
}

func parseSubCmd(s string) SubCmd {
	switch s {
	case "":
		return PARSE
	case "p":
		return PARSE
	case "parse":
		return PARSE
	case "v":
		return VERSION
	case "version":
		return VERSION
	case "r":
		return REPARSE
	case "reparse":
		return REPARSE
	default:
		return NILSUB
	}
}
