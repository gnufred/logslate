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
	outputToProgress := func(hide bool, file args.LogFormat) bool {
		return !hide && file != args.SCREEN
	}
	outputToFile := func(file string) bool {
		return file != ""
	}

	a := args.Singleton()
	return &OutputTo{
		Progress:  outputToProgress(a.Parse.HideProg, a.Threat.Format),
		DB:        outputToFile(a.Parse.Db),
		ThreatLog: outputToFile(a.Threat.Path),
		Report:    outputToFile(a.Report.Path),
	}
}

func createBar(command args.SubCmd) *progbar.ProgressBar {
	bar := progbar.NewOptions(-1,
		progbar.OptionSetDescription(barActionWord(command)),
		progbar.OptionSetWriter(os.Stdout),
		progbar.OptionShowBytes(true),
		progbar.OptionSetWidth(10),
		progbar.OptionThrottle(50*time.Millisecond),
		progbar.OptionShowCount(),
		progbar.OptionSpinnerType(14),
		progbar.OptionSetRenderBlankState(true),
	)
	return bar
}

func barActionWord(command args.SubCmd) string {
	switch command {
	case args.PARSE:
		return "parsing"
	case args.REPARSE:
		return "reparsing"
	default:
		return "processing"
	}

}

func withinRange(log, since, until time.Time) bool {
	return !log.Before(since) && !log.After(until)
}
