package args

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestValidateHelp(t *testing.T) {
	test.DisableOutput()
	narg := 0
	args.Help = true
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.Help = false
	test.EnableOutput()
}

func TestInvalidDebugArg(t *testing.T) {
	test.DisableOutput()
	narg := 0
	args.Subcmd = PARSE
	args.Debug = 999
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.Debug = 0
	test.EnableOutput()
}

func TestInvalidNargs(t *testing.T) {
	test.DisableOutput()
	narg := 2
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	test.EnableOutput()
}

func TestInvalidLayout(t *testing.T) {
	test.DisableOutput()
	narg := 1
	args.Subcmd = PARSE
	args.ReportLayout = "banana!"
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	test.EnableOutput()
}

func TestCustomReportWithJsonFormat(t *testing.T) {
	test.DisableOutput()
	nargs := 0
	args.Subcmd = PARSE
	args.ReportPath = "report.json"
	args.ReportFormat = JSON
	args.ReportLayout = "parsed"
	validate(nargs)
	args.ReportPath = ""
	args.ReportLayout = ""
	test.EnableOutput()
}

func TestDefaultArgs(t *testing.T) {
	Init()
	args := Get()

	if args.Help {
		t.Error("debug flag is not false")
	}

	if args.Debug != 0 {
		t.Error("debug flag is not 0")
	}

	if args.CpuProf != "" {
		t.Error("cpuprof flag is not an empty string")
	}

	if args.MemProf != "" {
		t.Error("memprof flag is not an empty string")
	}

	if args.since != "" {
		t.Error("since flag is not an empty string")
	}
	if args.until != "" {
		t.Error("until flag is not an empty string")
	}

	if args.Db != "" {
		t.Error("db flag is not an empty string")
	}

	if args.LogFrmt != nginx.Combined {
		fmt.Println(args)
		t.Error("log_format flag is different than nxparse.Combined")
	}

	nilTime := time.Time{}
	if args.startTime != nilTime {
		t.Error("startTime flag is not time.Time{}")
	}
	if args.endTime == nilTime {
		t.Error("endTime flag is time.Time{}")
	}

	if args.ThreatPath != "" {
		t.Error("ThreatPath is not an empty string")
	}
	if args.ThreatFormat != TEXT {
		t.Error("ThreatFormat is not TEXT")
	}

	if args.ReportPath != "" {
		t.Error("ReportPath is not an empty string")
	}
	if args.ReportFormat != TEXT {
		t.Error("ReportFormat is not SCREEN")
	}
	if args.ReportLayout != "" {
		t.Error("ReportLayout is not an empty string")
	}
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
		{"parse", PARSE},
		{"invalid", NILSUB},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
