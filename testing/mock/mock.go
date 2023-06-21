package mock

import (
	"log"
	"os"
)

func Stdin(input []byte) {
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
