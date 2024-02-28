package args

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestValidate(t *testing.T) {
	test.DisableOutput()
	validateHelp(t)
	invalidDebugArg(t)
	invalidThreatPath(t)
	invalidReportPath(t)
	invalidNargs(t)
	invalidLayout(t)
	customReportWithJsonFormat(t)
	test.EnableOutput()
}

func validateHelp(t *testing.T) {
	narg := 0
	args.Help = true
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.Help = false
}

func invalidDebugArg(t *testing.T) {
	narg := 0
	args.Debug = 999
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.Debug = 0
}

func invalidThreatPath(t *testing.T) {
	narg := 0
	args.ThreatPath = "report.invalid"
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.ThreatPath = ""
}

func invalidReportPath(t *testing.T) {
	narg := 0
	args.ReportPath = "report.invalid"
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
	args.ReportPath = ""
}

func invalidNargs(t *testing.T) {
	narg := 1
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
}

func invalidLayout(t *testing.T) {
	narg := 0
	args.ReportLayout = "banana!"
	err := validate(narg)
	test.ErrorExpected(t, err, "help shown")
}

func customReportWithJsonFormat(t *testing.T) {
	args.ReportPath = "report.json"
	args.ReportFormat = JSON
	args.ReportLayout = "parsed"
	validate(0)
	args.ReportPath = ""
	args.ReportLayout = ""
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
