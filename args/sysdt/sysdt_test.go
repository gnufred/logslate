package sysdt

import (
	"strings"
	"testing"
	"time"
)

func TestWorded(t *testing.T) {
	testWord := func(t *testing.T, raw string, exp time.Time) {
		got, err := Worded(raw)
		if raw == "now" {
			got = got.Truncate(time.Second)
		}
		if err != nil {
			t.Errorf("got %v, exp %v", got, exp)
		}
		if !got.Equal(exp) {
			t.Errorf("got %v, exp %v", got, exp)
		}
	}

	testWord(t, "now", time.Now().Truncate(time.Second))
	testWord(t, "today", time.Now().Truncate(time.Hour*24))
	testWord(t, "yesterday", time.Now().Add(time.Hour*-24).Truncate(time.Hour*24))
	testWord(t, "tomorrow", time.Now().Add(time.Hour*24).Truncate(time.Hour*24))
	testInvalidWord(t, "invalid_word")
}

func testInvalidWord(t *testing.T, raw string) {
	_, err := Worded(raw)
	if !strings.Contains(err.Error(), "no match") {
		t.Errorf("did not get expected error: %s", err)
	}
}

func TestAdjust(t *testing.T) {
	testAdjust := func(t *testing.T, ts string, ti time.Time, exp time.Time) {
		got, err := Adjust(ts, ti)
		if err != nil {
			t.Error(err)
		}
		if got != exp {
			t.Errorf("got %v, exp %v", got, exp)
		}
	}

	testAdjust(t, "4 days 2 hr",
		time.Date(2023, 5, 12, 5, 0, 0, 0, time.UTC),
		time.Date(2023, 5, 16, 7, 0, 0, 0, time.UTC),
	)

	testInvalidAdjust(t)
}

func testInvalidAdjust(t *testing.T) {
	nilTime := time.Time{}
	tm, err := Adjust("invalid_duration", nilTime)
	if !strings.Contains(err.Error(), "incorrect format") {
		t.Errorf("did not get expected error: %s", err)
	}
	if tm != nilTime {
		t.Errorf("time not nil")
	}
}

func TestUnitParsingValue(t *testing.T) {
	data := []struct {
		u   string
		exp time.Duration
	}{
		{"usec", time.Microsecond},
		{"us", time.Microsecond},
		{"msec", time.Millisecond},
		{"ms", time.Millisecond},
		{"s", time.Second},
		{"sec", time.Second},
		{"second", time.Second},
		{"seconds", time.Second},
		{"m", time.Minute},
		{"min", time.Minute},
		{"minute", time.Minute},
		{"minutes", time.Minute},
		{"h", time.Hour},
		{"hr", time.Hour},
		{"hour", time.Hour},
		{"hours", time.Hour},
		{"d", 24 * time.Hour},
		{"day", 24 * time.Hour},
		{"days", 24 * time.Hour},
		{"w", 7 * 24 * time.Hour},
		{"week", 7 * 24 * time.Hour},
		{"weeks", 7 * 24 * time.Hour},
		{"M", time.Duration(30.44 * float64(24) * float64(time.Hour))},
		{"month", time.Duration(30.44 * float64(24) * float64(time.Hour))},
		{"months", time.Duration(30.44 * float64(24) * float64(time.Hour))},
		{"y", time.Duration(365.25 * float64(24) * float64(time.Hour))},
		{"year", time.Duration(365.25 * float64(24) * float64(time.Hour))},
		{"years", time.Duration(365.25 * float64(24) * float64(time.Hour))},
	}

	testParse := func(t *testing.T, u string, exp time.Duration) {
		got, err := unitToDuration(u)
		if err != nil {
			t.Error(err)
		}

		if got == exp {
			t.Logf("good input %s: got %d, exp %d", u, got, exp)
		} else {
			t.Errorf("bad input %s: got %d, exp %d", u, got, exp)
		}
	}

	for _, d := range data {
		testParse(t, d.u, d.exp)
	}
}

func TestUnitParsingSuccess(t *testing.T) {

	data := []struct {
		u   string
		exp bool
	}{
		{"usecs", false},
		{"usec", true},
		{"us", true},
		{"msec", true},
		{"ms", true},
		{"xmsecs", false},
		{"s", true},
		{"sec", true},
		{"second", true},
		{"sonds", false},
		{"seconds", true},
		{"m", true},
		{"min", true},
		{"minute", true},
		{"minutes", true},
		{"minutos", false},
		{"h", true},
		{"hr", true},
		{"hour", true},
		{"hours", true},
		{"hrour", false},
		{"d", true},
		{"day", true},
		{"days", true},
		{"dayz", false},
		{"w", true},
		{"week", true},
		{"weeks", true},
		{"wek", false},
		{"M", true},
		{"month", true},
		{"months", true},
		{"moths", false},
		{"y", true},
		{"year", true},
		{"years", true},
		{"yars", false},
	}

	testParse := func(t *testing.T, u string, exp bool) {
		_, err := unitToDuration(u)
		if (err == nil) == exp {
			t.Logf("good input %s: exp %t", u, exp)
		} else {
			t.Errorf("bad input %s: err %v', exp %t", u, err, exp)
		}
	}

	for _, d := range data {
		testParse(t, d.u, d.exp)
	}
}

func TestToDuration(t *testing.T) {
	testDuration := func(t *testing.T, ts string, exp time.Duration) {
		got, err := parseDuration(ts)
		if err != nil {
			t.Error(err)
		}
		if got != exp {
			t.Errorf("got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		ts  string
		exp time.Duration
	}{
		{"2 min 23sec", 2*time.Minute + 23*time.Second},
		{"1sec", 1 * time.Second},
		{"10 minutes", 10 * time.Minute},
		{"-10 minutes", -10 * time.Minute},
		{"-10 minutes ago", 10 * time.Minute},
		{"1 hour", 1 * time.Hour},
		{"-1hour", -1 * time.Hour},
		{"-1hour ago", 1 * time.Hour},
		{"-2day", -2 * 24 * time.Hour},
		{"2day ago", -2 * 24 * time.Hour},
		{"-2day ago", 2 * 24 * time.Hour},
	}

	for _, d := range data {
		testDuration(t, d.ts, d.exp)
	}
}
