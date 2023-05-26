package args

import (
	"log"
	"time"

	"gitlab.com/gnufred/logslate/args/sysdt"

	"github.com/araddon/dateparse"
)

func parseTimestamp(timestamp string, unset time.Time) time.Time {
	// Return default value if no timestamp is specified.
	if timestamp == "" {
		return unset
	}

	// Return time based on known English words.
	wordedTime, err := sysdt.Worded(timestamp)
	if err == nil {
		return wordedTime
	}

	// Return time parsed from systemd format.
	now := time.Now()
	adjustedTime, err := sysdt.Adjust(timestamp, now)
	if err == nil {
		return adjustedTime
	}

	// Return time parse from an unknown layout.
	parsedTime, err := dateparse.ParseLocal(timestamp)
	if err != nil {
		log.Fatalf("parsing %s error %s", timestamp, err)
	}
	return parsedTime
}
