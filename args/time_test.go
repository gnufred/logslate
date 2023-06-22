package args

import (
	"strings"
	"testing"
	"time"
)

func TestSetSinceUntil(t *testing.T) {
	validSinceUntil(t)
	invalidSince(t)
	invalidUntil(t)
}

func validSinceUntil(t *testing.T) {
	sinceArg := "-24 hours ago"
	untilArg := "now"

	a := Get()
	err := a.SetSinceUntil(sinceArg, untilArg)
	if err != nil {
		t.Error(err)
	}

	since := a.Since()
	if since != sinceArg {
		t.Errorf("since: bad value, got %s, exp %s", since, sinceArg)
	}

	until := a.Until()
	if until != untilArg {
		t.Errorf("until: bad value, got %s, exp %s", until, untilArg)
	}

	start := a.StartTime().Add(-24 * time.Hour).Truncate(time.Hour)
	exp := time.Now().Truncate(time.Hour)
	if start != exp {
		t.Errorf("start: bad value, got %v, exp %v", start, exp)
	}

	end := a.EndTime().Truncate(time.Hour)
	exp = time.Now().Truncate(time.Hour)
	if end != exp {
		t.Errorf("end: bad value, got %v, exp %v", end, exp)
	}
}

func invalidSince(t *testing.T) {
	sinceArg := "invalid_time_string"
	untilArg := "now"

	a := Get()
	err := a.SetSinceUntil(sinceArg, untilArg)
	if !strings.Contains(err.Error(), "unable to parse") {
		t.Errorf("did not get expected error: %s", err)
	}
}

func invalidUntil(t *testing.T) {
	sinceArg := "now"
	untilArg := "invalid_time_string"

	a := Get()
	err := a.SetSinceUntil(sinceArg, untilArg)
	if !strings.Contains(err.Error(), "unable to parse") {
		t.Errorf("did not get expected error: %s", err)
	}
}

func TestNullTimestamp(t *testing.T) {
	tm := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	testStamp(t, "", tm)
}

func TestBadTimestamp(t *testing.T) {
	_, err := parseTimestamp("invalid_time_string", time.Time{})
	if !strings.Contains(err.Error(), "unable to parse") {
		t.Errorf("did not get expected error: %s", err)
	}
}

func TestTranslatedTimeNow(t *testing.T) {
	tm, err := parseTimestamp("now", time.Time{})
	if err != nil {
		t.Error(err)
	}
	got := tm.Truncate(time.Second)
	exp := time.Now()
	exp = exp.Truncate(time.Second)
	testTime(t, got, exp)
}

func TestTranslatedTimeToday(t *testing.T) {
	tm, err := parseTimestamp("today", time.Time{})
	if err != nil {
		t.Error(err)
	}
	got := tm
	exp := time.Now()
	exp = exp.Truncate(time.Hour * 24)
	testTime(t, got, exp)
}

func TestTranslatedTimeYesterday(t *testing.T) {
	tm, err := parseTimestamp("yesterday", time.Time{})
	if err != nil {
		t.Error(err)
	}
	got := tm
	exp := time.Now()
	exp = exp.Add(time.Hour * -24).Truncate(time.Hour * 24)
	testTime(t, got, exp)
}

func testTime(t *testing.T, got time.Time, expected time.Time) {
	if !got.Equal(expected) {
		t.Errorf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}

func TestTrstringerGosystemdtime(t *testing.T) {
	testSDT := func(t *testing.T, timestamp string, expected time.Time) {
		tm, err := parseTimestamp(timestamp, time.Time{})
		if err != nil {
			t.Error(err)
		}
		got := tm.Truncate(time.Minute)
		if !got.Equal(expected) {
			t.Errorf("parseTimestamp() error: got %v, expected %v", got, expected)
		}
	}

	tm := time.Now().Truncate(time.Minute)
	testSDT(t, "0 day", tm)
	testSDT(t, "0 hour", tm)
	testSDT(t, "0 minute", tm)

	tm = time.Now().Add(-24 * time.Hour).Truncate(time.Minute)
	testSDT(t, "-1 day", tm)
	testSDT(t, "-24 hour", tm)
	testSDT(t, "-24 hours", tm)
	testSDT(t, "-1440 minute", tm)
	testSDT(t, "-1440 minutes", tm)

	tm = time.Now().Add(-24 * 2 * time.Hour).Truncate(time.Minute)
	testSDT(t, "-2 day", tm)
	testSDT(t, "2 days ago", tm)
	tm = time.Now().Add(-24 * 3 * time.Hour).Truncate(time.Minute)
	testSDT(t, "-3 day", tm)
	testSDT(t, "3 days ago", tm)
	tm = time.Now().Add(-24 * 4 * time.Hour).Truncate(time.Minute)
	testSDT(t, "-4 day", tm)
	testSDT(t, "4 days ago", tm)
}

func TestAraddonDateparse(t *testing.T) {
	// Test random timestamp format.
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "2023", tm)
	testStamp(t, "1st Jan 2023", tm)
	testStamp(t, "jan 1 2023", tm)
	testStamp(t, "2023-01-01", tm)

	// Test NGINX format.
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "01/Jan/2023", tm)
	testStamp(t, "01/Jan/2023:00:00:00", tm)
	testStamp(t, "01/Jan/2023:00:00:00 +0000", tm)

	// Test NGINX format in UTC.
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "01/Jan/2023:00:00:00 +0000", tm)
}

func testStamp(t *testing.T, timestamp string, expected time.Time) {
	tm, err := parseTimestamp(timestamp, time.Time{})
	if err != nil {
		t.Error(err)
	}
	got := tm
	if !got.Equal(expected) {
		t.Errorf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}
