package args

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestValidate(t *testing.T) {
	test.DisableOutput()
	validateHelp(t)
	invalidDebugArg(t)
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
	if !strings.Contains(err.Error(), "help shown") {
		t.Errorf("did not get the expected error, got %s", err)
	}
	args.Help = false
}

func invalidDebugArg(t *testing.T) {
	narg := 0
	args.Debug = 999
	err := validate(narg)
	if !strings.Contains(err.Error(), "help shown") {
		t.Errorf("did not get the expected error, got %s", err)
	}
	args.Debug = 0
}

func invalidReportPath(t *testing.T) {
	narg := 0
	args.ReportPath = "report.invalid"
	args.ReportFormat = "invalid"
	err := validate(narg)
	if !strings.Contains(err.Error(), "help shown") {
		t.Errorf("did not get the expected error, got %s", err)
	}
	args.ReportPath = ""
	args.ReportFormat = ""
}

func invalidNargs(t *testing.T) {
	narg := 1
	err := validate(narg)
	if !strings.Contains(err.Error(), "help shown") {
		t.Errorf("did not get the expected error, got %s", err)
	}
}

func invalidLayout(t *testing.T) {
	narg := 0
	args.ReportLayout = "banana!"
	err := validate(narg)
	if !strings.Contains(err.Error(), "help shown") {
		t.Errorf("did not get the expected error, got %s", err)
	}
}

func customReportWithJsonFormat(t *testing.T) {
	args.ReportPath = "report.json"
	args.ReportFormat = "json"
	args.ReportLayout = "summary"
	validate(0)
	args.ReportPath = ""
	args.ReportFormat = ""
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
		t.Error("log_format flag is different then nxparse.Combined")
	}

	nilTime := time.Time{}
	if args.startTime != nilTime {
		t.Error("startTime flag is not time.Time{}")
	}
	if args.endTime == nilTime {
		t.Error("endTime flag is time.Time{}")
	}
}

func TestParseReportPath(t *testing.T) {
	// Test screen report
	frmt := parseReportPath("screen")
	if frmt != "text" {
		t.Errorf("unexpected format %s", frmt)
	}
	// Test .json report format.
	frmt = parseReportPath("/user/run/1000/example.com/report.json")
	if frmt != "json" {
		t.Errorf("unexpected format %s", frmt)
	}
	// Test .txt report format.
	frmt = parseReportPath("/user/run/1000/example.com/report.txt")
	if frmt != "text" {
		t.Errorf("unexpected format %s", frmt)
	}
	// Test unknown report format.
	frmt = parseReportPath("/user/run/1000/example.com/report.csv")
	if frmt != "" {
		t.Errorf("unexpected format %s", frmt)
	}
}
