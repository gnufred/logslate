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

type Args struct {
	Help    bool
	Debug   int
	CpuProf string
	MemProf string
	Db      string
	LogFrmt string

	// Time arguments (time.go)
	// Set with .SetSinceUntil()
	since     string    // Geter .Since()
	until     string    // Geter .Until()
	startTime time.Time // Geter .StartTime()
	endTime   time.Time // Geter .EndTime()

	// Report arguments (report.go)
	ReportPath   string
	ReportFormat string
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

func Help() {
	usage := `
PARSING
  -f, --log_format
    Need to match NGINX log_format directive from your nginx.conf.
    Default to NGINX default combined format.

  -s, --since
  -u, --until
    Parse entries using specified date(s).
    Can use timestamps or system.d style wording.

DATABASE
  -d, --db
    SQLite3 DB file path.

REPORT
  -r, --report
    Report file path. Must end with .txt or .json.

  -l, --layout
    Custom report layout for screen and text reports.

MORE
  You should read the detailed usage guide at:
  https://gitlab.com/gnufred/logslate/-/blob/main/.docs/usage.md
`

	fmt.Fprint(os.Stderr, usage)
}

func validate(narg int) error {
	showHelp := func() error {
		Help()
		return ErrHelpShown
	}

	// User asked for help.
	if args.Help {
		return showHelp()
	}

	// --debug argument is invalid.
	if args.Debug < 0 || args.Debug > 3 {
		fmt.Printf("Error: debug args must be -v or --debug=[1,3]\n\n")
		return showHelp()
	}

	// --layout contains only valid characters.
	regex := "[^a-z0-9:,\\_]"
	match, _ := regexp.MatchString(regex, args.ReportLayout)
	if match {
		fmt.Printf("Error: invalid characters in layout format.\n\n")
		return showHelp()
	}

	// --layout is not used with JSON format reports.
	if args.ReportFormat == "json" && args.ReportLayout != "" {
		fmt.Printf("Error: report layout is only available for screen and text reports.\n\n")
		return showHelp()
	}

	// There are extra unknown arguments or params.
	if narg != 0 {
		fmt.Printf("Error: unknown argument(s) found\n\n")
		return showHelp()
	}

	// Report format is invalid.
	if args.ReportPath != "" {
		if f := args.ReportFormat; f != "json" && f != "text" {
			fmt.Printf("Error: --format value must end with .json or .txt\n\n")
			return showHelp()
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
	reportFlag := flags.StringP("report", "r", "", "")

	// Report layout flag.
	reportLayoutFlag := flags.StringP("layout", "l", "", "")

	// Parse and save CLI args.
	args = Args{}
	pflag.Parse()

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

	args.ReportPath = string(*reportFlag)

	args.ReportFormat = parseReportPath(string(*reportFlag))

	args.ReportLayout = string(*reportLayoutFlag)

	return nil
}

// parseReportPath check the path file extension to determine if it's
// in a supported format. Supported formats are "txt", and "json".
func parseReportPath(path string) string {
	lpath := strings.ToLower(path)
	if path == "screen" {
		return "text"
	} else if strings.HasSuffix(lpath, ".txt") {
		return "text"
	} else if strings.HasSuffix(lpath, ".json") {
		return "json"
	}
	return ""
}
