package report

import (
	"os"
	"strings"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
)

func newTestReport() *Report {
	r := New()
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
	r := newTestReport()

	switch frmt {
	case "screen":
		r.Path = "screen"
		r.Frmt = "text"
	case "text":
		r.Path = "testreport.txt"
		r.Frmt = "text"
	case "json":
		r.Path = "testreport.json"
		r.Frmt = "json"
	}

	test.DisableOutput()
	r.Print()
	test.EnableOutput()

	if r.Path != "screen" {
		err := os.Remove(r.Path)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCreateTextOutput(t *testing.T) {
	testOut := func(t *testing.T, gs, exp string) {
		if !strings.Contains(gs, exp) {
			t.Errorf("text output didn't contain %s", exp)
		}
	}

	r := New()
	got, err := createTextOutput(r)
	if err != nil {
		t.Error(err)
	}
	gs := got.String()

	data := []string{
		"=== [[ Summary ]] ===",
		"=== [[ Time ]] ===",
		"=== [[ Threats ]] ===",
		"=== [[ Attacks ]] ===",
	}

	for _, s := range data {
		testOut(t, gs, s)
	}
}
