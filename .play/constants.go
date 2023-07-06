package main

import "fmt"

const (
	// Attack types.
	Unk = byte(iota + 1)
	SQL
	CMD
	Dir
	XSS
)

func main() {
	fmt.Println(Unk, SQL, CMD, Dir, XSS)
}
