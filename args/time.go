package args

import (
	"fmt"
	"time"

	"gitlab.com/gnufred/logslate/args/sysdt"

	"github.com/araddon/dateparse"
)

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

func (a *Args) SetSinceUntil(sincestamp string, untilstamp string) error {
	a.since = sincestamp
	a.until = untilstamp

	tm, err := parseTimestamp(sincestamp, time.Time{})
	if err != nil {
		return err
	}
	a.startTime = tm

	tm, err = parseTimestamp(untilstamp, time.Now())
	if err != nil {
		return err
	}
	a.endTime = tm

	return nil
}

func parseTimestamp(timestamp string, unset time.Time) (time.Time, error) {
	// Return default value if no timestamp is specified.
	if timestamp == "" {
		return unset, nil
	}

	// Return time based on known English words.
	wordedTime, err := sysdt.Worded(timestamp)
	if err == nil {
		return wordedTime, nil
	}

	// Return time parsed from systemd format.
	now := time.Now()
	adjustedTime, err := sysdt.Adjust(timestamp, now)
	if err == nil {
		return adjustedTime, nil
	}

	// Return time parse from an unknown layout.
	parsedTime, err := dateparse.ParseLocal(timestamp)
	if err != nil {
		return unset, fmt.Errorf("unable to parse \"%s\" into a usable time", timestamp)
	}
	return parsedTime, nil
}
