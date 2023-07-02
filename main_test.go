package main

import (
	"bufio"
	"errors"
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/root"
	"gitlab.com/gnufred/logslate/test"
)

func TestMain(t *testing.T) {
	file := "root/testdata/20ok.log"
	// TODO: fix test.MockStdin() to get 100% coverage.
	// file := "root/testdata/toolong.log"
	input, _ := os.ReadFile(file)
	test.MockStdin(input)
	main()
}

func TestInit(t *testing.T) {
	Init()
}

func TestBanner(t *testing.T) {
	Banner()
}

func TestRun(t *testing.T) {
	testRunOk(t)
	// TODO: fix test.MockStdin() to get 100% coverage.
	// testRunTooLong(t)
}

func testRunOk(t *testing.T) {
	file := "root/testdata/20ok.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)
	Run()
}

// TODO: fix test.MockStdin() to get 100% coverage.
/* func testRunTooLong(t *testing.T) {
	file := "root/testdata/toolong.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)
	Run()
} */

func TestExit(t *testing.T) {
	exitErrHelp(t)
	exitErrAny(t)
	exitErrTooLong(t)
	exitErrLimitReach(t)
}

func exitErrHelp(t *testing.T) {
	got := Exit(args.ErrHelpShown)
	exp := 1
	if got != exp {
		t.Errorf("bad exit status, got %d, exp %d", got, exp)
	}
}

func exitErrAny(t *testing.T) {
	got := Exit(errors.New("testing.error"))
	exp := 1
	if got != exp {
		t.Errorf("bad exit status, got %d, exp %d", got, exp)
	}
}

func exitErrTooLong(t *testing.T) {
	got := Exit(bufio.ErrTooLong)
	exp := 2
	if got != exp {
		t.Errorf("bad exit status, got %d, exp %d", got, exp)
	}
}

func exitErrLimitReach(t *testing.T) {
	got := Exit(root.ErrLimitReach)
	exp := 3
	if got != exp {
		t.Errorf("bad exit status, got %d, exp %d", got, exp)
	}
}
