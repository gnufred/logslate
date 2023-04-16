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

func (p *Piper) ReadLine() (string, error) {
	if scanned := p.scanner.Scan(); scanned {
		line := p.scanner.Text()
		return line, nil
	} else {
		if err := p.scanner.Err(); err == nil {
			// if Scan() is false and Err() is nil, we reached io.EOF
			return "", io.EOF
		} else {
			return "", err
		}
	}
}
