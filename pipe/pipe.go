package pipe

import (
	"bufio"
	"io"
	"os"
)

type Piper struct {
	scanner *bufio.Scanner
}

func New() *Piper {
	return &Piper{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (pi *Piper) ReadLine() (string, error) {
	if scanned := pi.scanner.Scan(); scanned {
		line := pi.scanner.Text()
		return line, nil
	} else {
		if err := pi.scanner.Err(); err == nil {
			// if Scan() is false and Err() is nil, we reached io.EOF
			return "", io.EOF
		} else {
			return "", err
		}
	}
}
