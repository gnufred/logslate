package args

import (
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestDefaultArgs(t *testing.T) {
	Init()
	a := Singleton()

	if a.Help {
		t.Error("debug flag is not false")
	}

	if a.Debug.Verbosity != 0 {
		t.Error("debug flag is not 0")
	}

	if a.Debug.CpuProf != "" {
		t.Error("cpuprof flag is not an empty string")
	}

	if a.Debug.MemProf != "" {
		t.Error("memprof flag is not an empty string")
	}

	if a.Time.since != "" {
		t.Error("since flag is not an empty string")
	}
	if a.Time.until != "" {
		t.Error("until flag is not an empty string")
	}

	if a.Parse.Db != "" {
		t.Error("db flag is not an empty string")
	}

	if a.Parse.LogFrmt != nginx.Combined {
		t.Error("log_format flag is different than nxparse.Combined")
	}

	nilTime := time.Time{}
	if a.Time.startTime != nilTime {
		t.Error("startTime flag is not time.Time{}")
	}
	if a.Time.endTime == nilTime {
		t.Error("endTime flag is time.Time{}")
	}

	if a.Threat.Path != "" {
		t.Error("ThreatPath is not an empty string")
	}
	if a.Threat.Format != TEXT {
		t.Error("ThreatFormat is not TEXT")
	}

	if a.Report.Path != "" {
		t.Error("ReportPath is not an empty string")
	}
	if a.Report.Format != TEXT {
		t.Error("ReportFormat is not TEXT")
	}
	if a.Report.Layout != "" {
		t.Error("ReportLayout is not an empty string")
	}
}

func TestValidateHelp(t *testing.T) {
	a := Singleton()
	narg := 0
	a.Help = true
	test.DisableOutput()
	err := validate(narg)
	test.EnableOutput()
	test.ErrorExpected(t, err, "help shown")
	a.Help = false
}

func TestInvalidDebugArg(t *testing.T) {
	a := Singleton()
	narg := 0
	a.Command = PARSE
	a.Debug.Verbosity = 999
	test.DisableOutput()
	err := validate(narg)
	test.EnableOutput()
	test.ErrorExpected(t, err, "debug args")
	a.Debug.Verbosity = 0
}

func TestInvalidNargs(t *testing.T) {
	narg := 2
	test.DisableOutput()
	err := validate(narg)
	test.EnableOutput()
	test.ErrorExpected(t, err, "multiple sub")
}

func TestInvalidLayout(t *testing.T) {
	a := Singleton()
	narg := 1
	a.Command = PARSE
	a.Report.Layout = "banana!"
	test.DisableOutput()
	err := validate(narg)
	test.EnableOutput()
	test.ErrorExpected(t, err, "layout format")
	a.Report.Layout = ""
}

func TestReparseInvalidDB(t *testing.T) {
	a := Singleton()
	narg := 1
	a.Command = REPARSE
	a.Parse.Db = "testdata/nill.db"
	test.DisableOutput()
	err := validate(narg)
	test.EnableOutput()
	test.ErrorExpected(t, err, "cannot use")
	a.Parse.Db = ""
}

func TestCustomReportWithJsonFormat(t *testing.T) {
	a := Singleton()
	nargs := 0
	a.Command = PARSE
	a.Report.Path = "report.json"
	a.Report.Format = JSON
	a.Report.Layout = "parsed"
	test.DisableOutput()
	validate(nargs)
	test.EnableOutput()
	a.Report.Path = ""
	a.Report.Layout = ""
}

func TestParseLogFormat(t *testing.T) {
	// Test screen report
	frmt := parseLogFormat("screen")
	if frmt != SCREEN {
		t.Errorf("unexpected format %d", frmt)
	}
	// Test .json report format.
	frmt = parseLogFormat("/user/run/1000/example.com/report.json")
	if frmt != JSON {
		t.Errorf("unexpected format %d", frmt)
	}
	// Test .txt report format.
	frmt = parseLogFormat("/user/run/1000/example.com/report.txt")
	if frmt != TEXT {
		t.Errorf("unexpected format %d", frmt)
	}
	frmt = parseLogFormat("/user/run/1000/example.com/report.log")
	if frmt != TEXT {
		t.Errorf("unexpected format %d", frmt)
	}
}

func TestParseSubcmd(t *testing.T) {
	testMatch := func(t *testing.T, in string, exp SubCmd) {
		got := parseSubCmd(in)
		if got != exp {
			t.Errorf("sub command miss match, got %d, exp %d", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp SubCmd
	}{
		{"", PARSE},
		{"p", PARSE},
		{"parse", PARSE},
		{"v", VERSION},
		{"version", VERSION},
		{"r", REPARSE},
		{"reparse", REPARSE},
		{"invalid", NILSUB},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
