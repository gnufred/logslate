package report

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	got := New("", "", "")
	if got.Line.Ok != 0 {
		t.Error("empty report has non empty values")
	}
}

func TestNewText(t *testing.T) {
	got := NewText(New("", "", ""))
	if got.Line.Ok != 0 {
		t.Error("empty report has non empty values")
	}
}

func TestLinesTotal(t *testing.T) {
	l := Lines{Ok: 1, Skip: 1, Error: 1}
	got := l.Total()
	exp := uint(3)
	if got != exp {
		t.Errorf("line total doesn't match, got %d, exp %d", got, exp)
	}
}

func TestTimeTotalValid(t *testing.T) {
	now := time.Now()
	tm := Runtimes{
		Start: now,
		End:   now.Add(1 * time.Hour),
	}
	got := tm.Total()
	exp, _ := time.ParseDuration("1h")
	if got != exp {
		t.Errorf("time total doesn't match, got %d, exp %d", got, exp)
	}
}

func TestAllPerSecSmallerThanZero(t *testing.T) {
	now := time.Now()
	r := Report{
		Run: Runtimes{
			Start: now,
			End:   now.Add(10 * time.Millisecond),
		},
		Line: Lines{Ok: 1},
		Size: Sizes{Byte: 1},
	}
	if got := r.BytesPerSec(); got != 1 {
		t.Errorf("bytes per sec doesn't match, exp 0, got %d", got)
	}
	if got := r.LinesPerSec(); got != 1 {
		t.Errorf("lines per sec doesn't match, exp 0, got %d", got)
	}
}

func TestAllPerSecNil(t *testing.T) {
	r := Report{}
	if got := r.BytesPerSec(); got != 0 {
		t.Errorf("bytes per sec doesn't match, exp 0, got %d", got)
	}
	if got := r.LinesPerSec(); got != 0 {
		t.Errorf("lines per sec doesn't match, exp 0, got %d", got)
	}
}

func TestAllPerSecValid(t *testing.T) {
	now := time.Now()
	r := Report{
		Run: Runtimes{
			Start: now,
			End:   now.Add(10 * time.Second),
		},
		Line: Lines{Ok: 50},
		Size: Sizes{Byte: 500},
	}
	exp := uint(50)
	if got := r.BytesPerSec(); got != exp {
		t.Errorf("bytes per sec doesn't match, exp %d, got %d", exp, got)
	}
	exp = uint(5)
	if got := r.LinesPerSec(); got != exp {
		t.Errorf("lines per sec doesn't match, exp %d, got %d", exp, got)
	}
}

func TestAttackTotal(t *testing.T) {
	at := AtkTypes{1, 1, 1, 1, 1}
	got := at.Total()
	exp := uint(5)
	if got != exp {
		t.Errorf("total doesn't match, got %d, exp %d", got, exp)
	}
}
