package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/debug"
	"gitlab.com/gnufred/logslate/root"
)

var Version = "dev"

func main() {
	Init()
	Banner()
	Run()
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

func Run() {
	a := args.Get()
	debug.Log(debug.VV, "start time", a.StartTime())
	debug.Log(debug.VV, "end time", a.EndTime())

	_, err := root.ParseStdin(a)
	if err != nil {
		log.Fatal(err)
	}
}
