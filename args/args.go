package args

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
	nginx "gitlab.com/gnufred/nxparse"
)

type Args struct {
	Help      bool
	Debug     int
	CpuProf   string
	MemProf   string
	since     string
	until     string
	Db        string
	LogFrmt   string
	startTime time.Time
	endTime   time.Time
}

var args = Args{}

func Init() {
	if !pflag.Parsed() {
		parse()
		validate()
	}
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
}

func parse() {
	// Get the current pflag SetFlag
	flags := pflag.CommandLine
	flags.SortFlags = false
	hideFlags := true

	// Help flags.
	helpFlag := flags.BoolP("help", "h", false, "Show help message")
	if hideFlags {
		flags.MarkHidden("help")
	}

	// Debug flags.
	flags.CountP("debug", "v", "Debug: verbosity")
	if hideFlags {
		flags.MarkHidden("debug")
	}

	// pprof profiling flags.
	cpuprofFlag := flags.StringP("profcpu", "", "", "Profile CPU usage")
	if hideFlags {
		flags.MarkHidden("profcpu")
	}
	flags.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	memprofFlag := flags.StringP("profmem", "", "", "Profile memory usage")
	if hideFlags {
		flags.MarkHidden("profmem")
	}
	flags.Lookup("profmem").NoOptDefVal = "mem.prof"

	// Time flags.
	sinceFlag := flags.StringP("since", "s", "", "Parse entries since specified date")
	untilFlag := flags.StringP("until", "u", "", "Parse entries until specified date")

	// Database flags.
	dbFlag := flags.StringP("db", "d", "", "SQLite3 DB file path")

	// NGINX `log_format` flags.
	logfrmtFlag := flags.StringP("log_format", "f", nginx.Combined, "NGINX log_format directive")

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
}

func (a *Args) SetSinceUntil(sincestamp string, untilstamp string) {
	a.since = sincestamp
	a.until = untilstamp
	a.startTime = parseTimestamp(sincestamp, time.Time{})
	a.endTime = parseTimestamp(untilstamp, time.Now())
}
func Get() Args {
	return args
}

func (a *Args) Since() string {
	return a.since
}

func (a *Args) Until() string {
	return a.until
}

func (a *Args) StartTime() time.Time {
	return a.startTime
}

func (a *Args) EndTime() time.Time {
	return a.endTime
}
