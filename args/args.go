package args

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/trstringer/go-systemd-time/pkg/systemdtime"
)

type Args struct {
	V         bool
	Vv        bool
	Vvv       bool
	Db        string
	Since     string
	Until     string
	StartTime time.Time
	EndTime   time.Time
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

	// process CLIS flags
	flag.Parse()

	// return structured arguments
	args = Args{
		V:         bool(*vFlag),
		Vv:        bool(*vvFlag),
		Vvv:       bool(*vvvFlag),
		Db:        string(*dbFlag),
		Since:     string(*sinceFlag),
		Until:     string(*untilFlag),
		StartTime: parseTimestamp(string(*sinceFlag), UnsetTime),
		EndTime:   parseTimestamp(string(*untilFlag), time.Now()),
	}
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
	if err == nil {
		log.Fatalf("parsing %s error %s", timestamp, err)
	}
	return parsedTime
}

func Get() Args {
	return args
}
