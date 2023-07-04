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
	Banner()
	err := Init()
	if err != nil {
		status := Exit(err)
		os.Exit(status)
	}

	err = Run()
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

func Banner() {
	fmt.Printf("Logslate %s @ https://gitlab.com/gnufred/logslate\n", Version)
}

func Run() error {
	a := args.Get()
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
