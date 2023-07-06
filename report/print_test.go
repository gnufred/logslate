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

func TestAtkrOptions(t *testing.T) {
	rep := New("screen", "text", "attackers")
	data := NewText(rep)
	err := atkrOptions(data, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")
}

func TestAtkrOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Atkr["127.0.0.1"] = Atkrs{10}
	rep.Atkr["127.0.0.2"] = Atkrs{100}
	data := NewText(rep)

	err := atkrOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.AtkrCnt != nil {
		t.Error("attacker count is not nil")
	}

	err = atkrOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.AtkrCnt[0].Item != "127.0.0.2" {
		t.Error("attacker order by count data is incorrect")
	}
}

func TestBotOptions(t *testing.T) {
	rep := New("screen", "text", "bots")
	data := NewText(rep)
	err := botOptions(data, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")
}

func TestBotOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Bot["one"] = Bots{10}
	rep.Bot["two"] = Bots{100}
	data := NewText(rep)

	err := botOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.BotCnt != nil {
		t.Error("bot count is not nil")
	}

	err = botOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.BotCnt[0].Item != "two" {
		t.Error("bot ordered by count data is incorrect")
	}
}
