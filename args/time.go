package args

import (
	"fmt"
	"time"

	"gitlab.com/gnufred/logslate/args/sysdt"

	"github.com/araddon/dateparse"
)

func (t *timeArgs) Since() string {
	return t.since
}

func (t *timeArgs) Until() string {
	return t.until
}

func (t *timeArgs) StartTime() time.Time {
	return t.startTime
}

func (t *timeArgs) EndTime() time.Time {
	return t.endTime
}

func (t *timeArgs) SetSince(since string) error {
	var err error
	t.since = since
	t.startTime, err = parseTimestamp(since, time.Time{})
	return err
}

func (t *timeArgs) SetUntil(until string) error {
	var err error
	t.until = until
	t.endTime, err = parseTimestamp(until, time.Now())
	return err
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
