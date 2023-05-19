package args

import (
	"testing"
	"time"
)

func TestNullTimestamp(t *testing.T) {
	tm := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	testStamp(t, "", tm)
}

func TestTranslatedTimeNow(t *testing.T) {
	got := parseTimestamp("now", time.Time{})
	got = got.Truncate(time.Second)
	exp := time.Now()
	exp = exp.Truncate(time.Second)
	testTime(t, got, exp)
}

func TestTranslatedTimeToday(t *testing.T) {
	got := parseTimestamp("today", time.Time{})
	exp := time.Now()
	exp = exp.Truncate(time.Hour * 24)
	testTime(t, got, exp)
}

func TestTranslatedTimeYesterday(t *testing.T) {
	got := parseTimestamp("yesterday", time.Time{})
	exp := time.Now()
	exp = exp.Add(time.Hour * -24).Truncate(time.Hour * 24)
	testTime(t, got, exp)
}

func testTime(t *testing.T, got time.Time, expected time.Time) {
	if !got.Equal(expected) {
		t.Fatalf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}

func TestTrstringerGosystemdtime(t *testing.T) {
	testSDT := func(t *testing.T, timestamp string, expected time.Time) {
		got := parseTimestamp(timestamp, time.Time{}).Truncate(time.Minute)
		if !got.Equal(expected) {
			t.Fatalf("parseTimestamp() error: got %v, expected %v", got, expected)
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
	// test random timestamp format
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "2023", tm)
	testStamp(t, "1st Jan 2023", tm)
	testStamp(t, "jan 1 2023", tm)
	testStamp(t, "2023-01-01", tm)

	// test NGINX format
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "01/Jan/2023", tm)
	testStamp(t, "01/Jan/2023:00:00:00", tm)
	testStamp(t, "01/Jan/2023:00:00:00 +0000", tm)

	// test NGINX format in UTC
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testStamp(t, "01/Jan/2023:00:00:00 +0000", tm)
}

func testStamp(t *testing.T, timestamp string, expected time.Time) {
	got := parseTimestamp(timestamp, time.Time{})
	if !got.Equal(expected) {
		t.Fatalf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}
