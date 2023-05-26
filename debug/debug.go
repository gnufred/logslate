package debug

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/gnufred/logslate/args"
)

const OFF = int(0)
const V = int(1)
const VV = int(2)
const VVV = int(3)

var verbosity = OFF

func Init() {
	// Set verbosity from the environment variable.
	env, _ := strconv.Atoi(os.Getenv("LS_DEBUG"))
	if env == OFF || env == V || env == VV || env == VVV {
		setVerbosity(env)
	}

	// Override value from command line flag if set.
	arg := args.Get()
	if arg.V {
		setVerbosity(V)
	}
	if arg.Vv {
		setVerbosity(VV)
	}
	if arg.Vvv {
		setVerbosity(VVV)
	}
}

func Log(msgverb int, name string, msg interface{}) error {
	verb := GetVerbosity()

	if verb == OFF {
		return fmt.Errorf("error: debug output disabled")
	}

	if msgverb <= verb {
		fmt.Printf("(debug%d)|%s| %v\n", msgverb, name, msg)
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
