package pipe

import (
	"os"
	"testing"
)

func mockStdin(input []byte) error {
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	_, err = w.Write(input)
	if err != nil {
		return err
	}
	w.Close()

	// restore stdin
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	return nil
}

// https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
func (p *Piper) TestReadLine(t *testing.T) {
	input := []byte("anything")
	err := mockStdin(input)
	if err != nil {
		t.Fatalf("Fatal error while mocking Stdin")
	}
	text, err := p.ReadLine()
	if text != string(input) {
		t.Fatalf("io.stding read failure. got %s, expected %s", text, input)
	}
	t.Log("Its alive")
}
