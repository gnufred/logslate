package main

import (
    "fmt"
	"io"
	"log"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/parse/nginx"
	"gitlab.com/gnufred/logslate/parse/slate"

	"github.com/schollz/progressbar/v3"
)

var Version = "dev"

func main() {
    Banner()
    ParseStdin()
}

func Banner() {
    fmt.Printf("Logslate %s @ https://gitlab.com/gnufred/logslate\n", Version)
}

func ParseStdin() {
	a := args.Parse()
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

	lineNum := 0
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
			continue
		}

        log := slate.ParseNginx(nginx.Parse(line))

		if a.Db != "" {
			db.InsertLog(log)
		}

		bar.Add(byteCnt)
		lineNum++
    }
}
