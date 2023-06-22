package debug

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gitlab.com/gnufred/logslate/args"
)

// Verbosity level constants.
const OFF = int(0)
const V = int(1)
const VV = int(2)
const VVV = int(3)

// Verbosity level. Change with setVerbosity().
var verbosity = OFF

// stderr output logger.
var stderr = log.New(os.Stderr, "", 0)

func Init() {
	verb := 0
	a := args.Get()

	argDebug := a.Debug
	if argDebug > 0 {
		verb = argDebug
	}

	envDebug, _ := strconv.Atoi(os.Getenv("LS_DEBUG"))
	if envDebug > 0 {
		verb = envDebug
	}

	if verb == V || verb == VV || verb == VVV {
		setVerbosity(verb)
	}
}

func Log(msgverb int, name string, msg interface{}) error {
	verb := GetVerbosity()

	if verb == OFF {
		return fmt.Errorf("error: debug output disabled")
	}

	if msgverb <= verb {
		stderr.Printf("v%d|%s| %v\n", msgverb, name, msg)
		return nil
	}

	return fmt.Errorf("dedug message level %d does not meet current set verbosity %d", msgverb, verb)
}

func setVerbosity(wantverb int) error {
	if wantverb >= OFF && wantverb <= VVV {
		verbosity = wantverb
		return nil
	}
	return fmt.Errorf("invalid debug verbosity value: %d", wantverb)
}

func GetVerbosity() int {
	return verbosity
}
