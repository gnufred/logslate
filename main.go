package main

import (
    "fmt"
	"io"
	"log"
    "os"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/debug"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/parse/nginx"
	"gitlab.com/gnufred/logslate/parse/slate"

	"github.com/schollz/progressbar/v3"
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
	p := pipe.New()

    if p.IsPiped() != true {
        log.Fatalf("Error: no input data")
    }

	db := db.Singleton()
	if a.Db != "" {
		err := db.Create(a.Db)
		if err != nil {
			log.Fatalf("Error while creating database: %s", err)
		}
	}

	lineNum, lineErr := 0, 0
	bar := progressbar.DefaultBytes(-1, "parsing")
	for {
		line, byteCnt, err := p.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading piped io.Stding: %s", err)
			break
		}

		if line == "" {
            lineErr++
			continue
		}

        nx, err := nginx.Parse(line)
        if err != nil {
            lineErr++
            continue
        }

        ls := slate.ParseNginx(nx)

		if a.Db != "" {
			db.InsertLog(ls)
		}

		bar.Add(byteCnt)
		lineNum++
    }
    // Need to print a return character after the last progress bar terminal update
    fmt.Print("\n")

    fmt.Printf("  parsed [%d] error [%d]\n", lineNum, lineErr)
}
