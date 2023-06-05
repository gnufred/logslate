package args

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
	nginx "gitlab.com/gnufred/nxparse"
)

type Args struct {
	Help       bool
	Debug      int
	CpuProf    string
	MemProf    string
	Db         string
	LogFrmt    string
	ReportPath string
	ReportFrmt string

	// Time arguments are managed in time.go
	// Set all 4 time vars .SetSinceUntil()
	since     string    // Geter .Since()
	until     string    // Geter .Until()
	startTime time.Time // Geter .StartTime()
	endTime   time.Time // Geter .EndTime()
}

var args = Args{}

func Init() {
	if !pflag.Parsed() {
		parse()
		validate()
	}
}

func Get() Args {
	return args
}

func validate() {
	showHelpAndExit := func() {
		fmt.Printf("Usage:\n")
		pflag.PrintDefaults()
		os.Exit(2)
	}

	// User asked for help.
	if args.Help {
		showHelpAndExit()
	}

	// --debug argument is invalid.
	if args.Debug < 0 || args.Debug > 3 {
		fmt.Printf("Error: debug args must be -v or --debug=[1,3]\n\n")
		showHelpAndExit()
	}

	// There are extra unknown arguments or params.
	if narg := pflag.NArg(); narg != 0 {
		fmt.Printf("Error: unknown argument(s) found\n\n")
		showHelpAndExit()
	}

	// Report format is invalid.
	if args.ReportPath != "" {
		if f := args.ReportFrmt; f != "json" && f != "text" {
			fmt.Printf("Error: --format value must end with .json or .txt\n\n")
			showHelpAndExit()
		}
	}
}

func parse() {
	// Get the current pflag SetFlag
	flags := pflag.CommandLine
	flags.SortFlags = false
	hideFlags := true

	// Help flags.
	helpFlag := flags.BoolP("help", "h", false, "Show help message.")
	if hideFlags {
		flags.MarkHidden("help")
	}

	// Debug flags.
	flags.CountP("debug", "v", "Debug: verbosity")
	if hideFlags {
		flags.MarkHidden("debug")
	}

	// pprof profiling flags.
	cpuprofFlag := flags.StringP("profcpu", "", "", "Profile CPU usage.")
	if hideFlags {
		flags.MarkHidden("profcpu")
	}
	flags.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	memprofFlag := flags.StringP("profmem", "", "", "Profile memory usage.")
	if hideFlags {
		flags.MarkHidden("profmem")
	}
	flags.Lookup("profmem").NoOptDefVal = "mem.prof"

	// Time flags.
	sinceFlag := flags.StringP("since", "s", "", "Parse entries since specified date.")
	untilFlag := flags.StringP("until", "u", "", "Parse entries until specified date.")

	// Database flags.
	dbFlag := flags.StringP("db", "d", "", "SQLite3 DB file path.")

	// NGINX `log_format` flags.
	logfrmtFlag := flags.StringP("log_format", "f", nginx.Combined, "NGINX log_format directive.")

	// Report flags.
	reportFlag := flags.StringP("report", "r", "", "Report file path. Must end with .txt or .json.")

	// Parse and save CLI args.
	args = Args{}
	pflag.Parse()

	args.Help = bool(*helpFlag)

	debugCnt, _ := flags.GetCount("debug")
	args.Debug = debugCnt

	args.CpuProf = string(*cpuprofFlag)

	args.MemProf = string(*memprofFlag)

	args.SetSinceUntil(string(*sinceFlag), string(*untilFlag))

	args.Db = string(*dbFlag)

	args.LogFrmt = string(*logfrmtFlag)

	args.ReportPath = string(*reportFlag)

	args.ReportFrmt = parseReportPath(string(*reportFlag))
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
