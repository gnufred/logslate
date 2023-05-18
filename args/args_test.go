package args

import (
	"log"
	"testing"
	"time"
)

func TestDefaultArgs(t *testing.T) {
	args := Get()

	if args.V {
		log.Fatal("v flag is not false")
	}
	if args.Vv {
		log.Fatal("vv flag is not false")
	}
	if args.Vvv {
		log.Fatal("vvv flag is not false")
	}

	if args.Db != "" {
		log.Fatal("db flag is not an empty string")
	}

	if args.since != "" {
		log.Fatal("db flag is not an empty string")
	}
	if args.until != "" {
		log.Fatal("db flag is not an empty string")
	}

	nilTime := time.Time{}
	if args.startTime != nilTime {
		log.Fatal("db flag is not an empty string")
	}
	if args.endTime != nilTime {
		log.Fatal("db flag is not an empty string")
	}
}
