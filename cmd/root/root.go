package rootcmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/db"
	"gitlab.com/gnufred/logslate/debug"
	"gitlab.com/gnufred/logslate/parse/slate"
	"gitlab.com/gnufred/logslate/pipe"
	"gitlab.com/gnufred/logslate/testing/prof"

	progbar "github.com/schollz/progressbar/v3"
	nginx "gitlab.com/gnufred/nxparse"
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
		d.ApplyDefaults()
	}

	if a.Db != "" {
		d.Begin()
		defer func(d *db.Dber) {
			d.Commit()
		}(d)
	}

	bar := createBar()

	// Start CPU profiling if argument is set.
	cpuf := prof.StartCpuProfile(a)
	defer prof.StopCpuProfile(a, cpuf)

	// Parser loop.
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
			debug.Log(debug.V, "parseError", line)
			continue
		}

		ls := slate.ParseNginx(nx)

		if !timeMatch(ls.Time, bool(a.Since() != ""), bool(a.Until() != ""), a.StartTime(), a.EndTime()) {
			par.Skip++
			continue
		}

		if a.Db != "" {
			d.InsertLog(ls)
			// Commit every "x" successfully parsed lines.
			// Testing with db/BenchmarkDbInsert shows that "7" is the best number of inserts to batch.
			if par.Ok%7 == 0 {
				d.Commit()
				d.Begin()
			}
		}

		par.Ok++
	}

	// Need to print a return character after the last progress bar terminal update.
	fmt.Print("\n")

	// Stop CPU profiling.
	prof.StopCpuProfile(a, cpuf)

	// Write memory heap if argument is set.
	prof.MemProfile(a)

	return par, nil
}

func createBar() *progbar.ProgressBar {
	bar := progbar.NewOptions(-1,
		progbar.OptionSetDescription("parsing"),
		progbar.OptionSetWriter(os.Stdout),
		progbar.OptionShowBytes(true),
		progbar.OptionSetWidth(10),
		progbar.OptionThrottle(166*time.Millisecond),
		progbar.OptionShowCount(),
		progbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stdout, "\n")
		}),
		progbar.OptionSpinnerType(14),
		progbar.OptionSetRenderBlankState(true),
	)
	return bar
}

func timeMatch(ts time.Time, since bool, until bool, start time.Time, end time.Time) bool {
	// First possibility: no arguments.
	if !since && !until {
		return true
	}

	// Second possibility: since is specified.
	if since && !until {
		if ts.Equal(start) || ts.After(start) {
			return true
		}
	}

	// Third possibility: until is specified.
	if !since && until {
		if ts.Before(end) || ts.Equal(end) {
			return true
		}
	}

	// Fourth possibility: both are specified.
	if since && until {
		if (ts.Equal(start) || ts.After(start)) &&
			(ts.Before(end) || ts.Equal(end)) {
			return true
		}
	}

	return false
}
