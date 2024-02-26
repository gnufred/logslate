package main

import "fmt"

const (
	// Attack types.
	UNK = byte(iota + 1)
	SQL
	CMD
	DIR
	XSS
)

func main() {
	fmt.Println(UNK, SQL, CMD, DIR, XSS)
}
