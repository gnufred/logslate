package pipe

import (
	"io"
	"os"
	"testing"
)

/* ============
Tests are based on the 2nd answer of this guide
https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
the `defer` is REQUIRED, otherwise the real io.stdin will be read and tests will fail
*/

type mockedStdinExpect struct {
	err error
	len int
}

func TestReadLine_OneLineBasic(t *testing.T) {
	// mock stding
	input, _ := os.ReadFile("testdata/oneline.log")
	mockStdin(t, input)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)

	// prepare data
	data := []mockedStdinExpect{
		{nil, 214},
	}

	// do the actual test
	pipe := New()
    if pipe.IsPiped() != true {
        t.Fatalf("data not from pipe")
    }
	//checkBytes(t, pipe, 215)
	checkData(t, pipe, data)
}

func TestReadLine_TwoLinesPlusEOF(t *testing.T) {
	// mock stding
	input, _ := os.ReadFile("testdata/twolinespluseof.log")
	mockStdin(t, input)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)

	// prepare data
	data := []mockedStdinExpect{
		{nil, 224},
		{nil, 166},
		{nil, 0},
		{io.EOF, 0},
	}

	// run tests
	pipe := New()
    if pipe.IsPiped() != true {
        t.Fatalf("data not from pipe")
    }
	//checkBytes(t, pipe, 393)
	checkData(t, pipe, data)
}

func checkData(t *testing.T, pipe *Piper, data []mockedStdinExpect) {
	for _, d := range data {
		text, _, err := pipe.ReadLine()
		checkError(t, err, d.err)
		checkLine(t, text, d.len)
	}
}

func checkBytes(t *testing.T, pipe *Piper, expected uint64) {
	len, err := pipe.BytesLen()
	if err != nil {
		t.Fatalf("BytesLen() error %s", err)
	}
	if len != expected {
		t.Fatalf("Got %d bytes, expeted %d", len, expected)
	}
}

func checkLine(t *testing.T, got string, expected int) {
	gotLen := len(got)
	if gotLen != expected {
		t.Fatalf("Got %d text, expeced %d", gotLen, expected)
	}
}

func checkError(t *testing.T, got error, expected error) {
	if got != expected {
		t.Fatalf("Got error %s, expected %s", got, expected)
	}
}

func mockStdin(t *testing.T, input []byte) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Mocking Stdin failed: %s", err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Fatalf("Fatal error while mocking Stdin")
	}
	w.Close()

	// restore stdin
	os.Stdin = r
}
