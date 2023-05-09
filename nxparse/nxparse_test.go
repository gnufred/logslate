package nxparse

import (
	"bufio"
	"errors"
	"os"
	"testing"
	"time"
)

// Test an empty line, which should return zero value for all fields
func TestMainEmptyStuff(t *testing.T) {
	_, err := Parse("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = Parse(Combined, "")
	if err != nil && !errors.Is(err, ErrSepNotFound) {
		t.Fatalf("got error: %s\nexpected: %s", err, ErrSepNotFound)
	}
}

// Test a single line of `combined` log format for the expected values
func TestMainCombinedOneLine(t *testing.T) {
	checkMatch := func(t *testing.T, got, exp interface{}) {
		if got != exp {
			t.Fatalf("bad log variable value: got %v, exp %v", got, exp)
		}
	}

	file, err := os.Open("testdata/combined_oneline.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	entries := []Entry{}
	for scanner.Scan() {
		entry, err := Parse(Combined, scanner.Text())
		if err != nil {
			t.Fatal(err)
		}
		entries = append(entries, entry)
	}
	e := entries[0]

	checkMatch(t, e.RemoteAddr, "216.245.221.84")
	checkMatch(t, e.RemoteUser, "-")
	tm, _ := time.Parse(timeLayoutLocal, "27/Apr/2023:09:49:32 +0000")
	checkMatch(t, e.TimeLocal, tm)
	checkMatch(t, e.Request, "HEAD /?monitoring=uptimerobot HTTP/1.1")
	checkMatch(t, e.Status, int64(200))
	checkMatch(t, e.BodyBytesSent, int64(0))
	checkMatch(t, e.Http["referer"], "https://example.com/?monitoring=uptimerobot")
	checkMatch(t, e.Http["user_agent"], "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)")
}

func TestMainCombinedNormal(t *testing.T) {
	file, err := os.Open("testdata/combined_normal.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		_, err := Parse(Combined, scanner.Text())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestMainCombinedTruncated(t *testing.T) {
	file, err := os.Open("testdata/combined_truncated.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txt := scanner.Text()
		_, err := Parse(Combined, txt)
		if err == nil {
			t.Fatalf("truncated line did not generate an error\n%s\n", txt)
		}
	}
}
