package debug

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gitlab.com/gnufred/logslate/args"
)

type verbosity int

// Verbosity level constants.
const (
	OFF verbosity = iota
	V
	VV
	VVV
)

// Verbosity level. Change with setVerbosity().
var verbositySetting = OFF

// stderr output logger.
var stderr = log.New(os.Stderr, "", 0)

func Init() {
	var verb, argDebug, envDebug verbosity
	verb = OFF
	a := args.Get()

	argDebug = verbosity(a.Debug)
	if argDebug > 0 {
		verb = argDebug
	}

	r, _ := strconv.Atoi(os.Getenv("LS_DEBUG"))
	envDebug = verbosity(r)
	if envDebug > 0 {
		verb = envDebug
	}

	if verb == V || verb == VV || verb == VVV {
		setVerbosity(verb)
	}
}

func Log(msgverb verbosity, name string, msg interface{}) error {
	verb := GetVerbosity()

	if verb == OFF {
		return fmt.Errorf("error: debug output disabled")
	}

	if msgverb <= verb {
		stderr.Printf("v%d|%s| %v\n", msgverb, name, msg)
		return nil
	}

	return fmt.Errorf(
		"dedug message level %d does not meet current set verbosity %d",
		msgverb, verb,
	)
}

func setVerbosity(want verbosity) error {
	if want >= OFF && want <= VVV {
		verbositySetting = want
		return nil
	}
	return fmt.Errorf("invalid debug verbosity value: %d", want)
}

func GetVerbosity() verbosity {
	return verbositySetting
}
