package args

import (
	"log"
	"testing"
	"time"
)

func TestDefaultArgs(t *testing.T) {
	args := Get()

	if args.V {
		log.Fatal("v flag is not false")
	}
	if args.Vv {
		log.Fatal("vv flag is not false")
	}
	if args.Vvv {
		log.Fatal("vvv flag is not false")
	}

	if args.Db != "" {
		log.Fatal("db flag is not an empty string")
	}

	if args.since != "" {
		log.Fatal("db flag is not an empty string")
	}
	if args.until != "" {
		log.Fatal("db flag is not an empty string")
	}

	if args.startTime != UnsetTime {
		log.Fatal("db flag is not an empty string")
	}
	if args.endTime != UnsetTime {
		log.Fatal("db flag is not an empty string")
	}
}

func TestNullTimestamp(t *testing.T) {
	tm := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	testSingleTimestamp(t, "", tm)
}

func TestTrstringerGosystemdtime(t *testing.T) {
	tm := time.Now().Truncate(time.Minute)
	testSingleSystemdtime(t, "0 day", tm)
	testSingleSystemdtime(t, "0 hour", tm)
	testSingleSystemdtime(t, "0 minute", tm)

	tm = time.Now().Add(-24 * time.Hour).Truncate(time.Minute)
	testSingleSystemdtime(t, "-1 day", tm)
	testSingleSystemdtime(t, "-24 hour", tm)
	testSingleSystemdtime(t, "-24 hours", tm)
	testSingleSystemdtime(t, "-1440 minute", tm)
	testSingleSystemdtime(t, "-1440 minutes", tm)

	tm = time.Now().Add(-24 * 2 * time.Hour).Truncate(time.Minute)
	testSingleSystemdtime(t, "-2 day", tm)
	tm = time.Now().Add(-24 * 3 * time.Hour).Truncate(time.Minute)
	testSingleSystemdtime(t, "-3 day", tm)
	tm = time.Now().Add(-24 * 4 * time.Hour).Truncate(time.Minute)
	testSingleSystemdtime(t, "-4 day", tm)
}

func testSingleSystemdtime(t *testing.T, timestamp string, expected time.Time) {
	got := parseTimestamp(timestamp, UnsetTime).Truncate(time.Minute)
	if !got.Equal(expected) {
		t.Fatalf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}

func TestAraddonDateparse(t *testing.T) {
	// test random timestamp format
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testSingleTimestamp(t, "2023", tm)
	testSingleTimestamp(t, "1st Jan 2023", tm)
	testSingleTimestamp(t, "jan 1 2023", tm)
	testSingleTimestamp(t, "2023-01-01", tm)

	// test NGINX format
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testSingleTimestamp(t, "01/Jan/2023", tm)
	testSingleTimestamp(t, "01/Jan/2023:00:00:00", tm)
	testSingleTimestamp(t, "01/Jan/2023:00:00:00 +0000", tm)

	// test NGINX format in UTC
	tm, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.UTC)
	testSingleTimestamp(t, "01/Jan/2023:00:00:00 +0000", tm)
}

func testSingleTimestamp(t *testing.T, timestamp string, expected time.Time) {
	got := parseTimestamp(timestamp, UnsetTime)
	if !got.Equal(expected) {
		t.Fatalf("parseTimestamp() error: got %v, expected %v", got, expected)
	}
}
