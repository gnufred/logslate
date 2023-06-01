package main

import (
	"fmt"
	"time"
)

func main() {
	dur := 10 * time.Second
	lines := 50
	bytes := 500

	lps := uint(int64(lines) / (int64(dur) / 1000000000))
	bps := uint(int64(bytes) / (int64(dur) / 1000000000))

	fmt.Println(dur, lines, bytes, lps, bps)
}
