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

func (p *Piper) BytesLen() (uint64, error) {
	/* This works, but os.Stdin file pointer moves forward
	and doesn't reset, so I can't read the content after */
	var cnt uint64
	p.scanner.Split(bufio.ScanRunes)
	for p.scanner.Scan() {
		cnt++
	}
	return cnt, nil
}

func (p *Piper) ReadLine() (string, int, error) {
	if scanned := p.scanner.Scan(); scanned {
		line := p.scanner.Text()
		byteCnt := utf8.RuneCountInString(line)
		return line, byteCnt, nil
	} else {
		if err := p.scanner.Err(); err == nil {
			// if Scan() is false and Err() is nil, we reached io.EOF
			return "", 0, io.EOF
		} else {
			return "", 0, err
		}
	}
}
