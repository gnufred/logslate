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

func TestMainPrintVersion(t *testing.T) {
	a := args.Singleton()
	a.Command = args.VERSION
	main()
}

func TestParseOk(t *testing.T) {
	a := args.Singleton()
	a.Command = args.PARSE
	file := "root/testdata/20ok.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)
	Exec(a)
}

func TestParseREPORT(t *testing.T) {
	a := args.Singleton()
	a.Command = args.REPARSE
	file := "root/testdata/threat.json"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)
	Exec(a)
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

func TestExecNoCmd(t *testing.T) {
	a := args.Singleton()
	a.Command = args.NILSUB
	err := Exec(a)
	test.ErrorExpected(t, err, "no command")
}

func TestExecVersion(t *testing.T) {
	execVersion()
}

func TestExitNoErr(t *testing.T) {
	testMatch := func(t *testing.T, in error, exp int) {
		got := Exit(in)
		if got != exp {
			t.Errorf("bad exit status, got %d, exp %d", got, exp)
		}
	}

	data := []struct {
		In  error
		Exp int
	}{
		{nil, 0},
		{errors.New("testing.error"), 1},
		{args.ErrHelpShown, 2},
		{bufio.ErrTooLong, 3},
		{root.ErrLimitReach, 4},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
