package test

import (
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/pipe"
)

func TestDisableOutput(t *testing.T) {
	DisableOutput()
	got := os.Stdout
	EnableOutput()
	if got == defaultStdout {
		t.Errorf("output still point to os.Stdout, got %v, os.Stdout %v", got, defaultStdout)
	}
}

func TestEnableOutput(t *testing.T) {
	if defaultStdout != os.Stdout {
		t.Errorf("output doesn't point to os.Stdout, got %v, os.Stdout %v", defaultStdout, os.Stdout)
	}
}

func TestMockStdin(t *testing.T) {
	// Mock the pipe with data from test file.
	file := "testdata/mockstdin.data"
	input, _ := os.ReadFile(file)
	MockStdin(input)

	// Read the data and validate it's content
	pi := pipe.New()
	got, _, err := pi.ReadLine()
	if err != nil {
		t.Error(err)
	}
	exp := "ABCDE"
	if exp != got {
		t.Errorf("mocked data doesn't match, got %s, exp %s", got, exp)
	}
}
