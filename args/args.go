package args

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	nginx "gitlab.com/gnufred/logslate/nxparse"

	"github.com/araddon/dateparse"
	"github.com/trstringer/go-systemd-time/pkg/systemdtime"
)

type Args struct {
	V         bool
	Vv        bool
	Vvv       bool
	Db        string
	LogFrmt   string
	since     string
	until     string
	startTime time.Time
	endTime   time.Time
}

var args = Args{}

// Default time value for unset or unparsed time values
// Use this variable to compare to time.Time arguments
// if sometime == args.UnsetTime { fmt.Print("not set") }
var UnsetTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

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
		Db:      string(*dbFlag),
		LogFrmt: string(*logfrmtFlag),
	}

	args.SetSinceUntil(string(*sinceFlag), string(*untilFlag))
}

func (a *Args) SetSinceUntil(sincestamp string, untilstamp string) {
	a.since = sincestamp
	a.until = untilstamp
	a.startTime = parseTimestamp(sincestamp, UnsetTime)
	a.endTime = parseTimestamp(untilstamp, time.Now())
}

func parseTimestamp(timestamp string, unset time.Time) time.Time {
	// return default time.Time if string is empty
	if timestamp == "" {
		return unset
	}

	// return time.Time parsed from systemd.time format
	now := time.Now()
	adjustedTime, err := systemdtime.AdjustTime(now, timestamp)
	if err == nil {
		return adjustedTime
	}

	// return time.Time parsed from unknown layout
	parsedTime, err := dateparse.ParseLocal(timestamp)
	if err != nil {
		log.Fatalf("parsing %s error %s", timestamp, err)
	}
	return parsedTime
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
