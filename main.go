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
	if err := Init(); err != nil {
		os.Exit(Exit(err))
	}
	a := args.Singleton()
	if err := Exec(a); err != nil {
		os.Exit(Exit(err))
	}
}

func Init() error {
	if err := args.Init(); err != nil {
		return err
	}
	debug.Init()
	return nil
}

func Exec(a *args.ArgsV2) error {
	switch a.Command {
	case args.VERSION:
		return execVersion()
	case args.PARSE:
		return execParse(a)
	case args.REPARSE:
		return execParse(a)
	default:
		return errors.New("no command")
	}
}

func execVersion() error {
	fmt.Println(Version)
	return nil
}

func execParse(a *args.ArgsV2) error {
	_, err := root.ParseStdin(a, Version)
	return err
}

func Exit(err error) int {
	if err == nil {
		return 0
	}

	if errors.Is(err, args.ErrHelpShown) {
		return 2
	}

	stderr := log.New(os.Stderr, "", 0)
	fmtErr := "\nfatal: %s."

	if errors.Is(err, bufio.ErrTooLong) {
		err = errors.New("input data might be corrupted or invalid")
		stderr.Printf(fmtErr, err)
		return 3
	}

	if errors.Is(err, root.ErrLimitReach) {
		stderr.Printf(fmtErr, err)
		fmt.Println("--log_format is probably incorrect.")
		return 4
	}

	stderr.Printf(fmtErr, err)

	return 1
}
