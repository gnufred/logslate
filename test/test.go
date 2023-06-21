// A test helper package.
// All functions in this package sould only be used in _test.go files.
// panic() is used to handle unexpected failure because it need to panic
// the test run.
package test

import (
	"fmt"
	"os"
)

var defaultStdout *os.File

func init() {
	defaultStdout = os.Stdout
}

func DisableOutput() {
	devnull, err := os.Open(os.DevNull)
	os.Stdout = devnull
	if err != nil {
		panic("cannot disable output")
	}
}

func EnableOutput() {
	os.Stdout = defaultStdout
}

func MockStdin(input []byte) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(fmt.Sprintf("mocking Stdin failed: %s", err))
	}

	_, err = w.Write(input)
	defer w.Close()
	if err != nil {
		panic("Error while mocking Stdin")
	}

	// restore stdin
	os.Stdin = r
}
