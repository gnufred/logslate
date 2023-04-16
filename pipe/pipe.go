package pipe

import (
	"bufio"
	"fmt"
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

func (pi *Piper) Read() {
	for pi.scanner.Scan() {
		fmt.Println(pi.scanner.Text())
	}
}
