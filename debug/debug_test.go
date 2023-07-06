package debug

import (
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

func TestInit(t *testing.T) {
	Init()
}

// It's currently not possible to test that Log() output works as expected.
func TestLog(t *testing.T) {
	setVerbosity(VVV)
	Log(VVV, "testing", "test debug message")
}

func TestValidVerbosityLevels(t *testing.T) {
	testValidVerbosity(t, OFF)
	testValidVerbosity(t, V)
	testValidVerbosity(t, VV)
	testValidVerbosity(t, VVV)
}

func testValidVerbosity(t *testing.T, want verbosity) {
	err := setVerbosity(want)
	if err != nil {
		t.Error(err)
	}
	verb := GetVerbosity()
	if verb != want {
		t.Errorf("GetVerbosity() got %d, expected %d", verb, want)
	}

}

func TestInvalidVerbosityLevels(t *testing.T) {
	testInvalidVerbosity(t, -1)
	testInvalidVerbosity(t, 4)
}

func testInvalidVerbosity(t *testing.T, want verbosity) {
	err := setVerbosity(want)
	test.ErrorExpected(t, err, "invalid debug verb")

}

func TestDebugDisabled(t *testing.T) {
	setVerbosity(OFF)
	err := Log(V, "log_name: test", "log_message: test")
	test.ErrorExpected(t, err, "debug output disabled")
}

func TestVerbFailures(t *testing.T) {
	setVerbosity(V)

	err := Log(VV, "log_name: test", "log_message: test")
	test.ErrorExpected(t, err, "does not meet")

	err = Log(VVV, "log_name: test", "log_message: test")
	test.ErrorExpected(t, err, "does not meet")
}
