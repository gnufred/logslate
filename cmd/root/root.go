package rootcmd

import (
	"fmt"
	"io"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/parse/nginx"
	"gitlab.com/gnufred/logslate/parse/slate"
	"gitlab.com/gnufred/logslate/pipe"

	"github.com/schollz/progressbar/v3"
)

func ParseStdin(a args.Args) (int, int, error) {
	p := pipe.New()

	if !p.IsPiped() {
		return 0, 0, fmt.Errorf("error: no input data")
	}

	d := db.Singleton()
	if a.Db != "" {
		err := d.Create(a.Db)
		if err != nil {
			return 0, 0, fmt.Errorf("error while creating database: %s", err)
		}
	}

	if a.Db != "" {
		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	lineNum, lineErr := 0, 0
	bar := progressbar.DefaultBytes(-1, "parsing")

	for {
		line, byteCnt, err := p.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, 0, fmt.Errorf("error while reading piped io.Stding: %s", err)
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
			d.InsertLog(ls)
			// From my basic tests, commiting 5 inserts at a time is what
			// gives me the best performance gain
			if lineNum%5 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		bar.Add(byteCnt)
		lineNum++
	}
	// Need to print a return character after the last progress bar terminal update
	fmt.Print("\n")

	return lineNum, lineErr, nil
}
