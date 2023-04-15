package pipe

import "fmt"

func ContainsData() bool {
	// Check if theres data in the pipe
}

func IsText() bool {
	// Check if piped data is text (uncompressed logs)
}

func IsGzip() bool {
	// Check if piped data is Gzip (.gz compressed logs)
}

func Unzip() (bytes, err) {
	// Unzip Gziped compressed data
}

func ReadLine() (string, err) {
	// Read and return a line from the pipe data
}

func Read() {
	// Place holder test function, will be delete
	fmt.Println("This is a pipe reader")
}
