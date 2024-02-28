package report

import (
	"strings"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
)

func TestStringifyErrStruct(t *testing.T) {
	_, err := Stringify("")
	test.ErrorExpected(t, err, "struct data")
}

func TestLayoutInvalidOption(t *testing.T) {
	r := newTestReport("bots:count_invalid")
	_, err := Stringify(*r)
	test.ErrorExpected(t, err, "invalid option")
}

func TestStringifyErrLayout(t *testing.T) {
	r := newTestReport("bots:count_invalid")
	_, err := Stringify(*r)
	test.ErrorExpected(t, err, "layout specification")
}

func TestStringify(t *testing.T) {
	r := newTestReport("ads")
	got, err := Stringify(*r)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	exp := `=== [[ Ads ]] ===
Pinterest: 10
`

	if got != exp {
		t.Errorf("stringify miss match, got %s, exp %s", got, exp)
	}
}

func TestValidTextOutput(t *testing.T) {
	testMatch := func(t *testing.T, layout, exp string) {
		r := New(layout)
		r.Threats.Bots.Its = 100
		r.Threats.Bots.Prct = 1
		tr, _ := prepareData(r)
		got, err := createTextOutput(tr)
		if err != nil {
			t.Error(err)
		}
		gs := got.String()
		if !strings.Contains(gs, exp) {
			t.Errorf("text output didn't contain %s\noutput:\n%s", exp, gs)
		}
	}

	data := []struct {
		Layout, Exp string
	}{
		{"parsed", header("parsed")},
		{"parsed", "Runtime:"},
		{"time", header("time")},
		{"time", "Start: 0001-01-01 00:00:00 UTC"},
		{"threats", header("threats")},
		{"threats", "Bots: 100 (1%)"},
	}

	for _, d := range data {
		testMatch(t, d.Layout, d.Exp)
	}
}

func TestLayoutDoesNotExist(t *testing.T) {
	r := New("banana!")
	tr := NewText(r)
	_, err := createTextOutput(tr)
	test.ErrorExpected(t, err, "does not exist")
}

func newTestReport(layout string) *Report {
	r := New(layout)
	r.Runtime = RuntimeRD{
		time.Now(),
		time.Now().Add(time.Second * 10),
	}
	r.Parsed = ParsedRD{5000, 0, 0, time.Time{}, time.Time{}}
	r.Size.Byte = 50000000
	r.Ads = map[string]ListRD{
		"Pinterest": {Cnt: 10},
	}
	r.Bots = map[string]ListRD{
		"Google": {Cnt: 100},
	}
	r.Attackers = map[string]ListRD{
		"1.2.3.4": {Cnt: 1000},
	}
	return r
}

func TestPrepareData(t *testing.T) {
	r := newTestReport("ads:order_count,bots:order_count,attackers:order_count")

	tr, err := prepareData(r)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	if tr.Ads["Pinterest"].Cnt != uint(10) {
		t.Errorf("filtered ads: Pinterest != 10")
	}

	if tr.Bots["Google"].Cnt != uint(100) {
		t.Errorf("filtered bots: Google != 100")
	}

	if tr.Attackers["1.2.3.4"].Cnt != uint(1000) {
		t.Errorf("filtered attackers: 1.2.3.4 != 1000")
	}
}

func TestHeader(t *testing.T) {
	testMatch := func(t *testing.T, in, exp string) {
		got := header(in)
		if got != exp {
			t.Errorf("header doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		In, Exp string
	}{
		{"parsed", "=== [[ Parsed ]] ==="},
		{"test", "=== [[ Test ]] ==="},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
