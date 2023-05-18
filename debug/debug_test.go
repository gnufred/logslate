package debug

import (
	"strings"
	"testing"
)

/* The reason there's no test that check that Log works as expected
   and does not report an error is that, working debug.Log() output
   messages. This means, the test results are populated with output
   that we don't want to have in there. So, we just test failures.
*/

func TestValidVerbosityLevels(t *testing.T) {
	testValidVerbosity(t, OFF)
	testValidVerbosity(t, V)
	testValidVerbosity(t, VV)
	testValidVerbosity(t, VVV)
}

func testValidVerbosity(t *testing.T, wantverb int) {
	err := setVerbosity(wantverb)
	if err != nil {
		t.Fatal(err)
	}
	verb := GetVerbosity()
	if verb != wantverb {
		t.Fatalf("GetVerbosity() got %d, expected %d", verb, wantverb)
	}

}

func TestInvalidVerbosityLevels(t *testing.T) {
	testInvalidVerbosity(t, -1)
	testInvalidVerbosity(t, 4)
}

func testInvalidVerbosity(t *testing.T, wantverb int) {
	err := setVerbosity(wantverb)
	exp := "invalid debug verb"
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("did not get expected error: %s", err)
	}

}

func TestDebugDisabled(t *testing.T) {
	setVerbosity(OFF)
	err := Log(V, "log_name: test", "log_message: test")
	exp := "debug output disabled"
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("did not get expected error: %s", err)
	}
}

func TestVerbFailures(t *testing.T) {
	setVerbosity(V)
	exp := "does not meet"

	err := Log(VV, "log_name: test", "log_message: test")
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("did not get expected error: %s", err)
	}

	err = Log(VVV, "log_name: test", "log_message: test")
	if !strings.Contains(err.Error(), exp) {
		t.Fatalf("did not get expected error: %s", err)
	}
}
