package root

import (
	"testing"
	"time"
)

func TestWithinRange(t *testing.T) {
	testMatch := func(t *testing.T, log, since, until time.Time, exp bool) {
		got := withinRange(log, since, until)
		if got != exp {
			t.Errorf("within time range miss match, got %t, exp %t", got, exp)
		}
	}
	tParse := func(t *testing.T, ts string) time.Time {
		if ts == "" {
			return time.Time{}
		}
		if tm, err := time.Parse("2006", ts); err != nil {
			t.Fatalf("cannot parse test time data, this really shouldn't happen")
			return time.Time{}
		} else {
			return tm
		}
	}
	data := []struct {
		Log   time.Time
		Since time.Time
		Until time.Time
		Exp   bool
	}{
		{tParse(t, "2000"), tParse(t, "1000"), tParse(t, "3000"), true},
		{tParse(t, "2000"), tParse(t, "2000"), tParse(t, "3000"), true},
		{tParse(t, "2000"), tParse(t, "1000"), tParse(t, "2000"), true},
		{tParse(t, "2000"), tParse(t, "2000"), tParse(t, "2000"), true},
		{tParse(t, "2000"), tParse(t, "3000"), tParse(t, "3000"), false},
		{tParse(t, "2000"), tParse(t, "1000"), tParse(t, "1000"), false},
		{tParse(t, "2000"), tParse(t, ""), tParse(t, ""), false},
		{tParse(t, "2000"), tParse(t, ""), tParse(t, "3000"), true},
	}

	for _, d := range data {
		testMatch(t, d.Log, d.Since, d.Until, d.Exp)
	}
}
