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
	V         bool
	Vv        bool
	Vvv       bool
	Db        string
	CpuProf   string
	MemProf   string
	LogFrmt   string
	since     string
	until     string
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

	if args.Help {
		showHelpAndExit()
	}
	// NArg is the number of arguments remaining after flags have been processed.
	if narg := pflag.NArg(); narg != 0 {
		fmt.Printf("Error: unknown argument(s) found\n\n")
		showHelpAndExit()
	}
}

func parse() {
	flags := pflag.CommandLine

	// help flag
	helpFlag := flags.BoolP("help", "h", false, "Show help message")

	// debug flags
	vFlag := flags.BoolP("v", "", false, "Debug: verbose")
	vvFlag := flags.BoolP("vv", "", false, "Debug: very verbose")
	vvvFlag := flags.BoolP("vvv", "", false, "Debug: very very verbose")

	// pprof flags
	cpuprofFlag := flags.StringP("profcpu", "", "", "Profile CPU usage")
	flags.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	memprofFlag := flags.StringP("profmem", "", "", "Profile memory usage")
	flags.Lookup("profmem").NoOptDefVal = "mem.prof"

	// db flags
	dbFlag := flags.StringP("db", "d", "", "SQLite3 DB file path")

	// time flags
	sinceFlag := flags.StringP("since", "s", "", "Parse entries since specified date")
	untilFlag := flags.StringP("until", "u", "", "Parse entries until specified date")

	// NGINX `log_format` flags
	logfrmtFlag := flags.StringP("log_format", "f", nginx.Combined, "NGINX log_format directive")

	flags.SortFlags = false
	flags.MarkHidden("help")
	flags.MarkHidden("v")
	flags.MarkHidden("vv")
	flags.MarkHidden("vvv")
	flags.MarkHidden("profcpu")
	flags.MarkHidden("profmem")

	// process CLIS flags
	pflag.Parse()

	// return structured arguments
	args = Args{
		Help:    bool(*helpFlag),
		V:       bool(*vFlag),
		Vv:      bool(*vvFlag),
		Vvv:     bool(*vvvFlag),
		CpuProf: string(*cpuprofFlag),
		MemProf: string(*memprofFlag),
		Db:      string(*dbFlag),
		LogFrmt: string(*logfrmtFlag),
	}

	args.SetSinceUntil(string(*sinceFlag), string(*untilFlag))
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
