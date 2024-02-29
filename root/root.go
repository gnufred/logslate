package root

import (
	"errors"
	"os"
	"time"

	progbar "github.com/schollz/progressbar/v3"
	"gitlab.com/gnufred/logslate/args"
)

type Parsed struct {
	Ok   uint
	Skip uint
	Err  uint
	Atk  uint
	Ad   uint
	Bot  uint
}

type OutputTo struct {
	Debug     bool
	Progress  bool
	DB        bool
	ThreatLog bool
	Report    bool
}

var (
	ErrNoInputData = errors.New("no input data")
	ErrReadPipe    = errors.New("error while reading piped io.Stdin")
	ErrLimitReach  = errors.New("parsing error limit reached")
)

func initOutput() *OutputTo {
	outputToProgress := func(file args.LogFormat) bool {
		return file != args.SCREEN
	}
	outputToFile := func(file string) bool {
		return file != ""
	}

	a := args.Singleton()
	return &OutputTo{
		Progress:  outputToProgress(a.Threat.Format),
		DB:        outputToFile(a.Parse.Db),
		ThreatLog: outputToFile(a.Threat.Path),
		Report:    outputToFile(a.Report.Path),
	}
}

func createBar() *progbar.ProgressBar {
	bar := progbar.NewOptions(-1,
		progbar.OptionSetDescription("parsing"),
		progbar.OptionSetWriter(os.Stdout),
		progbar.OptionShowBytes(true),
		progbar.OptionSetWidth(10),
		progbar.OptionThrottle(166*time.Millisecond),
		progbar.OptionShowCount(),
		progbar.OptionSpinnerType(14),
		progbar.OptionSetRenderBlankState(true),
	)
	return bar
}
