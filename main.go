package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/cmd/root"
	"gitlab.com/gnufred/logslate/debug"
)

var Version = "dev"

func main() {
	Init()
	Banner()
	ParseStdin()
}

func Init() {
	args.Init()
	debug.Init()
	debug.Log(debug.VVV, "os.Args", os.Args)
	debug.Log(debug.VVV, "args.Get", args.Get())
}

func Banner() {
	fmt.Printf("Logslate %s @ https://gitlab.com/gnufred/logslate\n", Version)
}

func ParseStdin() {
	a := args.Get()
	parsed, unparsed, err := cmdroot.ParseStdin(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  parsed [%d] error [%d]\n", parsed, unparsed)
}
