package main

import "fmt"

func main() {
	reportPath := "/user/run/1000/example.com/report.txt"
	rlen := len(reportPath)
	fmt.Println(rlen)

	fmt.Println(reportPath[rlen-4:])
	if reportPath[rlen-4:] == "json" {
		fmt.Println("Format is json")
	} else if reportPath[rlen-3:] == "txt" {
		fmt.Println("Format is txt")
	} else {
		fmt.Println("Format is unknown")
	}
}
