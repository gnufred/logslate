package args

import (
	"fmt"
	"testing"
	"time"

	nginx "gitlab.com/gnufred/nxparse"
)

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
	// Test .json report format.
	frmt := parseReportPath("/user/run/1000/example.com/report.json")
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
