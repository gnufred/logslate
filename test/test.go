// A test helper package.
// All functions in this package sould only be used in _test.go files.
// panic() is used to handle unexpected failure because it need to panic
// the test run.
package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var defaultStdout *os.File
var defaultStderr *os.File

func init() {
	defaultStdout = os.Stdout
	defaultStderr = os.Stderr
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
	os.Stderr = defaultStderr
}

func ErrorExpected(t *testing.T, err error, exp string) {
	if err == nil {
		t.Errorf("got an nil error\nexp: %s", exp)
		return
	}
	got := err.Error()
	if !strings.Contains(got, exp) {
		t.Errorf(
			"did not get the expected error\ngot: %s\nexp: %s",
			got, exp,
		)
	}
}

func MockStdin(input []byte) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(fmt.Sprintf("mocking Stdin failed: %s", err))
	}

	// https://gitlab.com/gnufred/logslate/-/issues/99
	// The code hang here when the input is 65536 character long, so
	// greater than a uint16.
	//
	// https://linux.die.net/man/7/pipe
	// Since Linux 2.6.11, the pipe capacity is 65536 bytes.
	//
	// To solve this issue we need to change the capacity of the pipe at runtime.
	//
	// Here's an example on how to do this with Python
	// https://www.golinuxhub.com/2018/05/how-to-view-and-increase-default-pipe-size-buffer/
	//
	// There's a constant for this in the syscall package but I don't understand how
	// to interact with it.
	// https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/syscall/zerrors_linux_amd64.go;l=11
	//
	// The best I could come up with is the following but it panic on a nil pointer.
	/*
		var p *[2]C.int
		syscall.RawSyscall(
			syscall.SYS_PIPE2,
			uintptr(unsafe.Pointer(p)),
			uintptr(syscall.F_SETPIPE_SZ),
			1048576,
		)
		w := os.NewFile(uintptr(p[1]), "|1")
		w.Write(input)
	*/
	_, err = w.Write(input)
	w.Close()

	if err != nil {
		panic("Error while mocking Stdin")
	}

	// restore stdin
	os.Stdin = r
}
