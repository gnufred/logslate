package pipe

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Piper struct {
	Bytes  int64
	Chunks int64
}

func New() *Piper {
	return &Piper{
		Bytes:  0,
		Chunks: 0,
	}
}

func (pi *Piper) Read() {
	// Based on https://zetcode.com/golang/pipe/
	pi.Bytes = 0
	pi.Chunks = 0
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	for {

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		pi.Chunks++
		pi.Bytes += int64(len(buf))

		fmt.Println(string(buf))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	fmt.Println("Bytes:", pi.Bytes, "Chunks:", pi.Chunks)
}

/* TODO
func ContainsData() bool {
	// Check if theres data in the pipe
}

func IsText() bool {
	// Check if piped data is text (uncompressed logs)
}

func ReadLine() (string, error) {
	// Read and return a line from the pipe data
}
*/
