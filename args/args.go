package args

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/pflag"
	nginx "gitlab.com/gnufred/nxparse"
)

type SubCmd byte
type LogFormat byte

const (
	NILSUB SubCmd = iota
	PARSE
	VERSION

	SCREEN LogFormat = iota
	TEXT
	JSON
)

type Args struct {
	// Subcommand.
	Subcmd SubCmd

	// Global arguments.
	Version string
	Help    bool
	Debug   int
	CpuProf string
	MemProf string
	Db      string
	LogFrmt string

	// Time arguments.
	// Set with .SetSinceUntil()
	since     string    // Geter .Since()
	until     string    // Geter .Until()
	startTime time.Time // Geter .StartTime()
	endTime   time.Time // Geter .EndTime()

	// Threat log arguments.
	ThreatPath   string
	ThreatFormat LogFormat

	// Report arguments.
	ReportPath   string
	ReportFormat LogFormat
	ReportLayout string
}

var args = Args{}

var ErrHelpShown = errors.New("help shown")

func Init() error {
	if !pflag.Parsed() {
		err := parse()
		if err != nil {
			return err
		}

		narg := pflag.NArg()
		err = validate(narg)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get() Args {
	return args
}

func Help(err string) {
	if err != "" {
		fmt.Fprint(os.Stderr, "\nError: "+err+".\n\n")
	}

	usage := `
USAGE
  logslate [SUBCOMMAND] [OPTIONS]

SUB COMMANDS
  p, parse (default)
    Parse NGINX access.log.

  v, version
    Print version.

PARSE
  -f, --log_format <NGINX LOG_FORMAT>
    Need to match NGINX log_format directive from your nginx.conf.
    Default to NGINX default combined format.

  -s, --since <TIME>
  -u, --until <TIME>
    Parse entries using specified date(s).
    Can use timestamps or system.d style wording.

  -d, --db <FILE>
    Convert NGINX access.log to an SQLite3 DB.

  -t, --threats (screen|<FILE>)
    Output threat log. ".json" files will be formatted as such.

  -r, --report (screen|<FILE>)
    Output report. ".json" files will be formatted as such.

  -l, --layout <LAYOUT>
    Custom report layout for screen and text reports.

MORE
  For detailed usage instructions read:
  https://gitlab.com/gnufred/logslate/-/blob/main/.docs/usage.md
`

	fmt.Fprint(os.Stderr, usage)
}

func validate(narg int) error {
	// Show help function.
	showHelp := func(err string) error {
		Help(err)
		return ErrHelpShown
	}

	// Return true if invalid.
	cHelp := func(a Args) bool { return args.Help }
	cManySubC := func(a Args) bool { return narg > 1 }
	cSubC := func(a Args) bool { return args.Subcmd == NILSUB }
	cDebug := func(a Args) bool { return a.Debug < 0 || a.Debug > 3 }
	cLayout := func(a Args) bool {
		match, _ := regexp.MatchString("[^a-z0-9:,\\_]", a.ReportLayout)
		return match
	}
	cReportFormat := func(a Args) bool { return a.ReportFormat == JSON && a.ReportLayout != "" }

	checks := []struct {
		F func(a Args) bool
		E string
	}{
		{cHelp, ""},
		{cManySubC, "multiple sub commands"},
		{cSubC, "invalid sub command"},
		{cDebug, "debug args must be -v or --debug=[1,3]"},
		{cLayout, "invalid characters in layout format"},
		{cReportFormat, "report layout is only available for screen and text reports"},
	}

	for _, c := range checks {
		if invalid := c.F(args); invalid {
			return showHelp(c.E)
		}
	}

	return nil
}

func parse() error {
	// Get the current pflag SetFlag
	flags := pflag.CommandLine
	flags.SortFlags = false
	hideFlags := true

	// Help flags.
	helpFlag := flags.BoolP("help", "h", false, "")
	if hideFlags {
		flags.MarkHidden("help")
	}

	// Debug flags.
	flags.CountP("debug", "v", "")
	if hideFlags {
		flags.MarkHidden("debug")
	}

	// pprof profiling flags.
	cpuprofFlag := flags.StringP("profcpu", "", "", "")
	if hideFlags {
		flags.MarkHidden("profcpu")
	}
	flags.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	memprofFlag := flags.StringP("profmem", "", "", "")
	if hideFlags {
		flags.MarkHidden("profmem")
	}
	flags.Lookup("profmem").NoOptDefVal = "mem.prof"

	// Time flags.
	sinceFlag := flags.StringP("since", "s", "", "")
	untilFlag := flags.StringP("until", "u", "", "")

	// Database flags.
	dbFlag := flags.StringP("db", "d", "", "")

	// NGINX `log_format` flags.
	logfrmtFlag := flags.StringP("log_format", "f", nginx.Combined, "")

	// Report flags.
	threatLogFlag := flags.StringP("threats", "t", "", "")

	// Report flags.
	reportFlag := flags.StringP("report", "r", "", "")

	// Report layout flag.
	reportLayoutFlag := flags.StringP("layout", "l", "", "")

	// Parse and save CLI args.
	args = Args{}
	pflag.Parse()

	// Get subcommand
	args.Subcmd = parseSubCmd(flags.Arg(0))

	args.Help = bool(*helpFlag)

	debugCnt, _ := flags.GetCount("debug")
	args.Debug = debugCnt

	args.CpuProf = string(*cpuprofFlag)

	args.MemProf = string(*memprofFlag)

	err := args.SetSinceUntil(string(*sinceFlag), string(*untilFlag))
	if err != nil {
		return err
	}

	args.Db = string(*dbFlag)

	args.LogFrmt = string(*logfrmtFlag)

	args.ThreatPath = string(*threatLogFlag)
	args.ThreatFormat = parseLogFormat(string(*threatLogFlag))

	args.ReportPath = string(*reportFlag)
	args.ReportFormat = parseLogFormat(string(*reportFlag))
	args.ReportLayout = string(*reportLayoutFlag)

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
	default:
		return NILSUB
	}
}
