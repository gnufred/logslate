package debug

import (
    "fmt"
    "os"
    "strconv"
)

const OFF = int(0)
const V   = int(1)
const VV  = int(2)
const VVV = int(3)

var verbosity = int(0)

func init() {
    wantverb, _ := strconv.Atoi(os.Getenv("LS_DEBUG"))
    setVerbosity(wantverb)
}

func Log(msgverb int, name string, msg interface{}) error {
    verb := getVerbosity()
    
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

func getVerbosity() int {
    return verbosity
}
