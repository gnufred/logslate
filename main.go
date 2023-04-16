package main

import (
	"fmt"

	"gitlab.com/gnufred/logslate/pipe"
)

func main() {
	pipe := pipe.New()
	fmt.Println(pipe.ReadLine())
	fmt.Println(pipe.ReadLine())
}
