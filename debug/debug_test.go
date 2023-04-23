package debug

import (
    "testing"
)

/* The reason there's no test that check that Log works as expected
   and does not report an error is that, working debug.Log() output
   messages. This means, the test results are populated with output
   that we don't want to have in there. So, we just test failures.
*/

func TestDebugDisabled(t *testing.T) {
    err := setVerbosity(OFF)
    if err != nil {
        t.Fatal(err)
    }

    verb := getVerbosity()
    if verb != OFF {
        t.Fatal("getVerbosity() unexpected return value")
    }

    err = Log(V, "log_name: test", "log_message: test")
    if err == nil {
        t.Fatal("Log() did not report an error, and should have")
    }
}

func TestVerbFailures(t *testing.T) {
    setVerbosity(V)
    
    err := Log(VV, "log_name: test", "log_message: test")
    if err == nil {
        t.Fatal("Log() did not report an error, and should have")
    }

    err = Log(VVV, "log_name: test", "log_message: test")
    if err == nil {
        t.Fatal("Log() did not report an error, and should have")
    }
}
