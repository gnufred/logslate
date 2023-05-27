package args

import (
	"fmt"
	"log"
	"testing"
	"time"

	nginx "gitlab.com/gnufred/nxparse"
)

func TestDefaultArgs(t *testing.T) {
	Init()
	args := Get()

	if args.Help {
		log.Fatal("debug flag is not false")
	}

	if args.Debug != 0 {
		log.Fatal("debug flag is not 0")
	}

	if args.CpuProf != "" {
		log.Fatal("cpuprof flag is not an empty string")
	}

	if args.MemProf != "" {
		log.Fatal("memprof flag is not an empty string")
	}

	if args.since != "" {
		log.Fatal("since flag is not an empty string")
	}
	if args.until != "" {
		log.Fatal("until flag is not an empty string")
	}

	if args.Db != "" {
		log.Fatal("db flag is not an empty string")
	}

	if args.LogFrmt != nginx.Combined {
		fmt.Println(args)
		log.Fatal("log_format flag is different then nxparse.Combined")
	}

	nilTime := time.Time{}
	if args.startTime != nilTime {
		log.Fatal("startTime flag is not time.Time{}")
	}
	if args.endTime == nilTime {
		log.Fatal("endTime flag is time.Time{}")
	}
}
