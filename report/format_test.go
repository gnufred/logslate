package report

import (
	"testing"
	"time"
)

func TestPrct(t *testing.T) {
	test := func(t *testing.T, divident, divisor uint, exp float64) {
		r := New()
		got := r.Prct(divident, divisor)
		if got != exp {
			t.Errorf("percentage doesn't match, got %f, exp %f", got, exp)
		}
	}

	data := []struct {
		Divident uint
		Divisor  uint
		Exp      float64
	}{
		{0, 999, 0},
		{999, 0, 0},
		{1, 100, 1},
		{10, 100, 10},
		{50, 100, 50},
		{1, 2, 50},
		{550, 17144, 3.2},
	}

	for _, d := range data {
		test(t, d.Divident, d.Divisor, d.Exp)
	}
}

func TestHumanTime(t *testing.T) {
	testTD := func(t *testing.T, tm time.Time, exp string) {
		r := New()
		if got := r.HumanTime(tm); got != exp {
			t.Errorf("human time didn't match, got %s, exp %s", got, exp)
		}
	}

	tm1, _ := time.Parse("2006-01-02 15:04:05 MST", "2023-06-05 12:58:03 UTC")
	data := []struct {
		Tim time.Time
		Exp string
	}{
		{time.Time{}, "0001-01-01 00:00:00 UTC"},
		{tm1, "2023-06-05 12:58:03 UTC"},
	}

	for _, d := range data {
		testTD(t, d.Tim, d.Exp)
	}
}

func TestHumanDuration(t *testing.T) {
	testHD := func(t *testing.T, d time.Duration, exp string) {
		r := New()
		if got := r.HumanDuration(d); got != exp {
			t.Errorf("human duration didn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		Dur time.Duration
		Exp string
	}{
		{123456789 * time.Nanosecond, "123ms"},
		{999999999 * time.Nanosecond, "999ms"},
		{123456789 * time.Millisecond, "34h17m36s"},
		{1233848 * time.Millisecond, "20m33s"},
		{12345 * time.Millisecond, "12.3s"},
		{10 * time.Second, "10s"},
	}

	for _, d := range data {
		testHD(t, d.Dur, d.Exp)
	}
}

func TestBytesToMB(t *testing.T) {
	testBTM := func(t *testing.T, b uint, m float64) {
		r := New()
		if got := r.BytesToMB(b); got != m {
			t.Errorf("bytes to MB didn't match, got %f, exp %f", got, m)
		}
	}

	data := []struct {
		Bytes uint
		MB    float64
	}{
		{1048576, 1},
		{1048576 * 2, 2},
		{1048576 * 4, 4},
	}

	for _, d := range data {
		testBTM(t, d.Bytes, d.MB)
	}
}

func TestAddCommas(t *testing.T) {
	testCommas := func(t *testing.T, n uint, exp string) {
		r := New()
		if got := r.AddCommas(n); got != exp {
			t.Errorf("add commas didn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		Num uint
		Exp string
	}{
		{0, "0"},
		{1234, "1,234"},
		{123456789, "123,456,789"},
	}

	for _, d := range data {
		testCommas(t, d.Num, d.Exp)
	}
}
