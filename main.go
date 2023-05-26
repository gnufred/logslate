package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
	rootcmd "gitlab.com/gnufred/logslate/cmd/root"
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
	par, err := rootcmd.ParseStdin(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed: ok [%d] skip [%d] error [%d]\n", par.Ok, par.Skip, par.Err)
}
