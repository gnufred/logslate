package args

import (
	"log"
	"time"

	"github.com/gnufred/go-systemd-time/pkg/systemdtime"

	"github.com/araddon/dateparse"
)

func parseTimestamp(timestamp string, unset time.Time) time.Time {
	// return default time.Time if string is empty
	if timestamp == "" {
		return unset
	}

	// try and understand english
	translatedTime, err := systemdtime.TranslateWords(timestamp)
	if err == nil {
		return translatedTime
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
