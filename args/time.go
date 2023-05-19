package args

import (
	"log"
	"time"

	"gitlab.com/gnufred/logslate/args/sysdt"

	"github.com/araddon/dateparse"
)

func parseTimestamp(timestamp string, unset time.Time) time.Time {
	// return default time.Time if string is empty
	if timestamp == "" {
		return unset
	}

	// try and understand english
	wordedTime, err := sysdt.Worded(timestamp)
	if err == nil {
		return wordedTime
	}

	// return time.Time parsed from systemd.time format
	now := time.Now()
	adjustedTime, err := sysdt.Adjust(timestamp, now)
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
