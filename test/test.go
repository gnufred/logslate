package test

import (
	"log"
	"os"
)

var stdout = os.Stdout

func DisableOutput() {
	stdout = os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
}

func EnableOutput() {
	os.Stdout = stdout
}

func MockStdin(input []byte) {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("Mocking Stdin failed: %s", err)
	}

	_, err = w.Write(input)
	defer w.Close()
	if err != nil {
		log.Fatalf("Error while mocking Stdin")
	}

	// restore stdin
	os.Stdin = r
}
