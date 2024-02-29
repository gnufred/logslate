package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/debug"
	"gitlab.com/gnufred/logslate/root"
)

// Changed at compile/runtime by `-ldflags`.
var Version = "dev"

func main() {
	err := Init()
	if err != nil {
		status := Exit(err)
		os.Exit(status)
	}

	a := args.Get()
	switch a.Subcmd {
	case args.VERSION:
		err = PrintVersion()
	case args.PARSE:
		err = Parse()
	}

	status := Exit(err)
	if status > 0 {
		os.Exit(status)
	}
}

func Init() error {
	err := args.Init()
	if err != nil {
		return err
	}

	debug.Init()
	debug.Log(debug.VVV, "os.Args", os.Args)
	debug.Log(debug.VVV, "args.Get", args.Get())

	return nil
}

func PrintVersion() error {
	fmt.Println(Version)
	return nil
}

func Parse() error {
	a := args.Get()
	a.Version = Version
	debug.Log(debug.VV, "start time", a.StartTime())
	debug.Log(debug.VV, "end time", a.EndTime())

	_, err := root.ParseStdin(a)
	return err
}

func Exit(err error) int {
	if err == nil {
		return 0
	}

	if errors.Is(err, args.ErrHelpShown) {
		return 1
	}

	stderr := log.New(os.Stderr, "", 0)
	fmtErr := "\nfatal: %s."

	if errors.Is(err, bufio.ErrTooLong) {
		err = errors.New("input data might be corrupted or invalid")
		stderr.Printf(fmtErr, err)
		return 2
	}

	if errors.Is(err, root.ErrLimitReach) {
		stderr.Printf(fmtErr, err)
		fmt.Println("--log_format is probably incorrect.")
		return 3
	}

	stderr.Printf(fmtErr, err)

	return 1
}
