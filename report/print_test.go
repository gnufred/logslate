package report

import (
	"os"
	"strings"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/test"
)

func newTestReport(path string, format args.LogFormat) *Report {
	r := New(path, format, defaultLayout)
	r.Runtime = RuntimeRD{
		time.Now(),
		time.Now().Add(time.Second * 10),
	}
	r.Parsed = ParsedRD{5000, 0, 0, time.Time{}, time.Time{}}
	r.Size.Byte = 50000000
	return r
}

func TestPrint(t *testing.T) {
	printFrmt(t, "screen")
	printFrmt(t, "text")
	printFrmt(t, "json")
}

func printFrmt(t *testing.T, frmt string) {
	var path string
	var format args.LogFormat

	switch frmt {
	case "screen":
		path = "screen"
		format = args.SCREEN
	case "text":
		path = "testreport.txt"
		format = args.TEXT
	case "json":
		path = "testreport.json"
		format = args.JSON
	}

	r := newTestReport(path, format)
	test.DisableOutput()
	r.Print()
	test.EnableOutput()

	if path != "screen" {
		err := os.Remove(path)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPrintTest(t *testing.T) {
	r := New("screen", args.TEXT, "banana!")
	err := printText(r)
	test.ErrorExpected(t, err, "report error")

	r = New("/invalid/path/fail/please", args.TEXT, "parsed")
	err = printText(r)
	test.ErrorExpected(t, err, "no such file")

	r = New("/invalid/path/fail/please", args.TEXT, "parsed")
	err = printJson(r)
	test.ErrorExpected(t, err, "no such file")
}

func TestCreateTextOutput(t *testing.T) {
	validTextOutput(t)
	invalidTextOutput(t)
}

func validTextOutput(t *testing.T) {
	testMatch := func(t *testing.T, layout, exp string) {
		r := New("screen", args.TEXT, layout)
		r.Threats.Bots.Its = 100
		r.Threats.Bots.Prct = 1
		got, err := createTextOutput(r)
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

func invalidTextOutput(t *testing.T) {
	r := New("screen", args.TEXT, "banana!")
	_, err := createTextOutput(r)
	test.ErrorExpected(t, err, "does not exist")

	r = New("screen", args.TEXT, "bots:banana!_lots!")
	_, err = createTextOutput(r)
	test.ErrorExpected(t, err, "layout specification")
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
