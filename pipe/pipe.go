package pipe

import (
	"bufio"
	"io"
	"os"
	"unicode/utf8"
)

type Piper struct {
	scanner *bufio.Scanner
}

func New() *Piper {
	return &Piper{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Source: https://stackoverflow.com/a/26567513/6900541
func (p *Piper) IsPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func (p *Piper) ReadLine() (string, int, error) {

	if scanned := p.scanner.Scan(); scanned {
		line := p.scanner.Text()
		// We need to add +2 to account for "\n" end line characters
		byteCnt := utf8.RuneCountInString(line) + 2
		return line, byteCnt, nil
	} else {
		// Error handling.
		if err := p.scanner.Err(); err == nil {
			// if Scan() is false and Err() is nil, we reached io.EOF
			return "", 0, io.EOF
		} else {
			return "", 0, err
		}
	}
}
