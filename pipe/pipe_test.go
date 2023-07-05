package pipe

import (
	"io"
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

// Tests are based on the 2nd answer from this guide:
// https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
// The `defer` is REQUIRED, otherwise, the real io.Stdin will be read and tests will fail.

type mockedStdinExpect struct {
	err error
	len int
}

func TestReadLine_OneLineBasic(t *testing.T) {
	// Mock Stdin.
	input, _ := os.ReadFile("testdata/oneline.log")
	test.MockStdin(input)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)

	// Prepare data.
	data := []mockedStdinExpect{
		{nil, 214},
	}

	// Run tests.
	pipe := New()
	if pipe.IsPiped() != true {
		t.Errorf("data not from pipe")
	}
	//checkBytes(t, pipe, 215)
	checkData(t, pipe, data)
}

func TestReadLine_TwoLinesPlusEOF(t *testing.T) {
	// Mock stdin.
	input, _ := os.ReadFile("testdata/twolinespluseof.log")
	test.MockStdin(input)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)

	// Prepare data.
	data := []mockedStdinExpect{
		{nil, 224},
		{nil, 166},
		{nil, 0},
		{io.EOF, 0},
	}

	// Run tests.
	pipe := New()
	if pipe.IsPiped() != true {
		t.Errorf("data not from pipe")
	}
	checkData(t, pipe, data)
}

func checkData(t *testing.T, pipe *Piper, data []mockedStdinExpect) {
	for _, d := range data {
		text, _, err := pipe.ReadLine()
		if d.err != nil {
			test.ErrorExpected(t, err, d.err.Error())
		}
		checkLine(t, text, d.len)
	}
}

func checkLine(t *testing.T, got string, expected int) {
	gotLen := len(got)
	if gotLen != expected {
		t.Errorf("Got %d text, expeced %d", gotLen, expected)
	}
}
