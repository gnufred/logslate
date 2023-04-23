package args

import (
	"log"
	"testing"
)

func TestDefaultArgs(t *testing.T) {
	args := Get()
	if args.Db != "" { log.Fatal("db flag is not an empty string") }
    if args.V { log.Fatal("v flag is not false") }
    if args.Vv { log.Fatal("vv flag is not false") }
    if args.Vvv { log.Fatal("vvv flag is not false") }
}
