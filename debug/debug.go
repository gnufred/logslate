package debug

import (
    "fmt"
    "os"
    "strconv"

    "gitlab.com/gnufred/logslate/args"
)

const OFF = int(0)
const V   = int(1)
const VV  = int(2)
const VVV = int(3)

var verbosity = OFF

func Init() {
    // set verbosity from environment variable
    v, _ := strconv.Atoi(os.Getenv("LS_DEBUG"))
    if v == OFF || v == V || v == VV || v == VVV {
        setVerbosity(v)
    }

    // override value from command line flag if set
    a := args.Get()
    if a.V { setVerbosity(V) }
    if a.Vv { setVerbosity(VV) }
    if a.Vvv { setVerbosity(VVV) }
}

func Log(msgverb int, name string, msg interface{}) error {
    verb := GetVerbosity()
    
    if verb == OFF {
        return fmt.Errorf("error: debug output disabled")
    }
    
    if msgverb <= verb {
        fmt.Printf("(debug%d){%s} %v\n", msgverb, name, msg)
        return nil
    }

    return fmt.Errorf("Dedug message level %d does not meet current set verbosity %d", msgverb, verb)
}

func setVerbosity(wantverb int) error {
    if wantverb >= OFF && wantverb <= VVV { 
        verbosity = wantverb
        return nil
    }
    return fmt.Errorf("Invalid debug verbosity value: %d", wantverb)
}

func GetVerbosity() int {
    return verbosity
}
