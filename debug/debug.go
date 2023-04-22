package debug

import (
    "fmt"
    "os"
)

func Log(msg interface{}) {
    if os.Getenv("LS_DEBUG") != "1" {
        return
    }
    fmt.Fprint(os.Stdout, "(debug) ")
    fmt.Fprint(os.Stdout, msg)
    fmt.Fprint(os.Stdout, "\n")
}
