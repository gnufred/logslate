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

type Parsed struct {
	Ok   int
	Skip int
	Err  int
}

func ParseStdin(a args.Args) (Parsed, error) {
	pi := pipe.New()
	par := Parsed{0, 0, 0}

	if !pi.IsPiped() {
		return par, fmt.Errorf("error: no input data")
	}

	d := db.Singleton()
	if a.Db != "" {
		err := d.Create(a.Db)
		if err != nil {
			return par, fmt.Errorf("error while creating database: %s", err)
		}
	}

	if a.Db != "" {
		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	bar := progressbar.DefaultBytes(-1, "parsing")

	for {
		line, byteCnt, err := pi.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			return par, fmt.Errorf("error while reading piped io.Stding: %s", err)
		}

		if line == "" {
			par.Err++
			continue
		}

		nx, err := nginx.Parse(line)
		if err != nil {
			par.Err++
			continue
		}

		ls := slate.ParseNginx(nx)

		if a.Db != "" {
			d.InsertLog(ls)
			// From my basic tests, commiting 5 inserts at a time is what
			// gives me the best performance gain
			if par.Ok%5 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		bar.Add(byteCnt)
		par.Ok++
	}
	// Need to print a return character after the last progress bar terminal update
	fmt.Print("\n")

	return par, nil
}
