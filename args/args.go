package args

import (
	"flag"
	"fmt"
	"os"
	"time"

	nginx "gitlab.com/gnufred/logslate/nxparse"
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
	// https://pkg.go.dev/flag#NFlag
	// NArg is the number of arguments remaining after flags have been processed.
	if narg := flag.NArg(); narg != 0 {
		fmt.Printf("unknown argument(s) found\nUsage of logslate:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func parse() {
	// debug flags
	vFlag := flag.Bool("v", false, "Debug: verbose")
	vvFlag := flag.Bool("vv", false, "Debug: very verbose")
	vvvFlag := flag.Bool("vvv", false, "Debug: very very verbose")

	// pprof flags
	cpuprofFlag := flag.String("cpuprof", "", "pprof cpu profile output file")
	memprofFlag := flag.String("memprof", "", "pprof cpu profile output file")

	// db flags
	dbFlag := flag.String("db", "", "SQLite3 DB file path")

	// time flags
	sinceFlag := flag.String("since", "", "Parse entries since specified date")
	untilFlag := flag.String("until", "", "Parse entries until specified date")

	// NGINX `log_format` flags
	logfrmtFlag := flag.String("log_format", nginx.Combined, "NGINX log_format directive")

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
