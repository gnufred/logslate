package rootcmd

import (
	"errors"
	"fmt"
	"io"
	"time"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	nginx "gitlab.com/gnufred/logslate/nxparse"
	"gitlab.com/gnufred/logslate/parse/slate"
	"gitlab.com/gnufred/logslate/pipe"

	"github.com/schollz/progressbar/v3"
)

type Parsed struct {
	Ok   int
	Skip int
	Err  int
}

var (
	ErrNoFormat     = errors.New("no log_format specified")
	ErrNoInputData  = errors.New("no input data")
	ErrDbCreateFail = errors.New("database creation error")
	ErrReadPipe     = errors.New("error while reading piped io.Stdin")
)

func ParseStdin(a args.Args) (Parsed, error) {
	pi := pipe.New()
	par := Parsed{0, 0, 0}

	if !pi.IsPiped() {
		return par, ErrNoInputData
	}

	if a.LogFrmt == "" {
		return par, ErrNoFormat
	}

	d := db.Singleton()
	if a.Db != "" {
		err := d.Create(a.Db)
		if err != nil {
			return par, fmt.Errorf("%s: %w", err, ErrDbCreateFail)
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
		bar.Add(byteCnt)

		if err == io.EOF {
			break
		}

		if err != nil {
			return par, fmt.Errorf("%s: %w", err, ErrReadPipe)
		}

		if line == "" {
			par.Err++
			continue
		}

		nx, err := nginx.Parse(a.LogFrmt, line)
		if err != nil {
			par.Err++
			continue
		}

		ls := slate.ParseNginx(nx)

		if !timeMatch(ls.Time, bool(a.Since() != ""), bool(a.Until() != ""), a.StartTime(), a.EndTime()) {
			par.Skip++
			continue
		}

		if a.Db != "" {
			d.InsertLog(ls)
			// Commit every "x" succesfully parsed lines
			// Testing with db/BenchmarkDbInsert shows that "7" is the best number of inserts to batch
			if par.Ok%7 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		par.Ok++
	}
	// Need to print a return character after the last progress bar terminal update
	fmt.Print("\n")

	return par, nil
}

func timeMatch(ts time.Time, since bool, until bool, start time.Time, end time.Time) bool {
	// first possibility: no arguments
	if !since && !until {
		return true
	}

	// second possibility: since is specified
	if since && !until {
		if ts.Equal(start) || ts.After(start) {
			return true
		}
	}

	// third possibility: until is specified
	if !since && until {
		if ts.Before(end) || ts.Equal(end) {
			return true
		}
	}

	// fourth possibility: both are specified
	if since && until {
		if (ts.Equal(start) || ts.After(start)) &&
			(ts.Before(end) || ts.Equal(end)) {
			return true
		}
	}

	return false
}
