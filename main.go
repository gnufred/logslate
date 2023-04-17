package main

import (
	"fmt"

	"gitlab.com/gnufred/logslate/pipe"
)

func main() {
	pipe := pipe.New()

	text, _ := pipe.ReadLine()
	fmt.Println(text)

	text, _ = pipe.ReadLine()
	fmt.Println(text)
}
