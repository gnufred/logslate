package report

import (
	"os"
	"strings"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
)

func newTestReport(path, format string) *Report {
	r := New(path, format, defaultLayout)
	r.Run = Runtimes{
		time.Now(),
		time.Now().Add(time.Second * 10),
	}
	r.Line = Lines{5000, 0, 0}
	r.Size.Byte = 50000000
	return r
}

func TestPrint(t *testing.T) {
	printFrmt(t, "screen")
	printFrmt(t, "text")
	printFrmt(t, "json")
}

func printFrmt(t *testing.T, frmt string) {
	var path, format string

	switch frmt {
	case "screen":
		path = "screen"
		format = "text"
	case "text":
		path = "testreport.txt"
		format = "text"
	case "json":
		path = "testreport.json"
		format = "json"
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
	r := New("screen", "text", "banana!")
	err := printText(r)
	test.ErrorExpected(t, err, "report error")

	r = New("/invalid/path/fail/please", "text", "summary")
	err = printText(r)
	test.ErrorExpected(t, err, "no such file")

	r = New("/invalid/path/fail/please", "json", "summary")
	err = printJson(r)
	test.ErrorExpected(t, err, "no such file")
}

func TestCreateTextOutput(t *testing.T) {
	validTextOutput(t)
	invalidTextOutput(t)
}

func validTextOutput(t *testing.T) {
	testMatch := func(t *testing.T, layout, exp string) {
		r := New("screen", "text", layout)
		r.Thrt.Bots.Its = 100
		r.Thrt.Bots.Prct = 1
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
		{"summary", header("summary")},
		{"summary", "Runtime:"},
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
	r := New("screen", "text", "banana!")
	_, err := createTextOutput(r)
	test.ErrorExpected(t, err, "does not exist")

	r = New("screen", "text", "bots:banana!_lots!")
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
		{"summary", "=== [[ Summary ]] ==="},
		{"test", "=== [[ Test ]] ==="},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
