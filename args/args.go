package args

import (
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	nginx "gitlab.com/gnufred/nxparse"
)

type Args struct {
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
	if flag.Parsed() {
		return
	}

	parse()
	validate()
}

func validate() {
	// NArg is the number of arguments remaining after flags have been processed.
	if narg := flag.NArg(); narg != 0 {
		fmt.Printf("unknown argument(s) found\nUsage of logslate:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func parse() {
	// debug flags
	vFlag := flag.BoolP("v", "", false, "Debug: verbose")
	vvFlag := flag.BoolP("vv", "", false, "Debug: very verbose")
	vvvFlag := flag.BoolP("vvv", "", false, "Debug: very very verbose")

	// pprof flags
	cpuprofFlag := flag.StringP("profcpu", "", "", "Profile CPU usage")
	flag.Lookup("profcpu").NoOptDefVal = "cpu.prof"
	memprofFlag := flag.StringP("profmem", "", "", "Profile memory usage")
	flag.Lookup("profmem").NoOptDefVal = "mem.prof"

	// db flags
	dbFlag := flag.StringP("db", "d", "", "SQLite3 DB file path")

	// time flags
	sinceFlag := flag.StringP("since", "s", "", "Parse entries since specified date")
	untilFlag := flag.StringP("until", "u", "", "Parse entries until specified date")

	// NGINX `log_format` flags
	logfrmtFlag := flag.StringP("log_format", "f", nginx.Combined, "NGINX log_format directive")

	// process CLIS flags
	flag.Parse()

	// return structured arguments
	args = Args{
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
