package detect

import "testing"

var testPts = []ptrn{
	{Name: "x", Type: Gen, Score: 1},
	{Name: "xy", Type: Gen, Score: 2},
	{Name: "xyz", Type: Gen, Score: 3},

	{Name: "\x99()", Type: Gen, Score: 255},
	{Name: "'SLEEP(0)", Type: SQL, Score: 255},
	{Name: "socket.hitmycommandcenter", Type: CMD, Score: 255},
	{Name: "../etc/passwd", Type: Dir, Score: 255},
	{Name: "<script>", Type: XSS, Score: 255},
}

func TestAttackTypeString(t *testing.T) {
	testMatch := func(t *testing.T, in byte, exp string) {
		got := AttackTypeString(in)
		if got != exp {
			t.Errorf("attack type string doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		In  byte
		Exp string
	}{
		{0, ""},
		{Gen, "generic"},
		{SQL, "sql_injection"},
		{CMD, "command_injection"},
		{Dir, "directory_traversal"},
		{XSS, "xss"},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
