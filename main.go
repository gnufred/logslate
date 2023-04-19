package main

import (
	"io"
	"log"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/parse/nginx"
	"gitlab.com/gnufred/logslate/parse/slate"

	"github.com/schollz/progressbar/v3"
)

func main() {
	a := args.Parse()
	p := pipe.New()
	db := db.Singleton()
	bar := progressbar.DefaultBytes(-1, "parsing")

	if a.Db != "" {
		err := db.Create(a.Db)
		if err != nil {
			log.Fatalf("Error while creating database: %s", err)
		}
	}

	lineNum := 0
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
