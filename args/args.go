package args

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	V   bool
	Vv  bool
	Vvv bool
	Db  string
}

var args = Args{}

func Init() {
	Parse()
	Validate(args)
}

func Validate(args Args) {
	// https://pkg.go.dev/flag#NFlag
	// NArg is the number of arguments remaining after flags have been processed.
	if narg := flag.NArg(); narg == 0 {
		return
	}

	fmt.Printf("unknown argument(s) found\nUsage of logslate:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func Parse() {
	vFlag := flag.Bool("v", false, "Debug: verbose")
	vvFlag := flag.Bool("vv", false, "Debug: very verbose")
	vvvFlag := flag.Bool("vvv", false, "Debug: very very verbose")
	dbFlag := flag.String("db", "", "SQLite3 DB file path")
	flag.Parse()

	args = Args{
		V:   bool(*vFlag),
		Vv:  bool(*vvFlag),
		Vvv: bool(*vvvFlag),
		Db:  string(*dbFlag),
	}
}

func Get() Args {
	return args
}
