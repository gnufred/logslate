package pipe

import (
	"os"
	"testing"
)

// Based on https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
// Needs refactoring
// Need to create a mockStding() function to reuse within tests
func (p *Piper) Test_ReadLine(t *testing.T) {
	input := []byte("anything")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	w.Close()

	// restore stdin
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	text, err := p.ReadLine()
	if text != "anything" {
		t.Fatalf("io.stding read failure. got %s, expected %s", text, input)
	}
}
