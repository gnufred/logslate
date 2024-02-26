package log

import (
	"net"
	"testing"

	"gitlab.com/gnufred/logslate/test"
	"gitlab.com/gnufred/logslate/threat"
)

func TestStringifyErrStruct(t *testing.T) {
	_, err := StringifyThreatLog("")
	test.ErrorExpected(t, err, "struct data")
}

func TestStringify(t *testing.T) {
	tl := ThreatLog{
		Detection: threat.Detection{
			Ad:  "threatscan",
			Bot: "Detectify",
			Attack: &threat.Attack{
				Type: threat.XSS,
			},
		},
		Ip:    net.ParseIP("1.2.3.4"),
		Path:  "/",
		Query: "'PG_SLEEP(100);",
	}

	got, err := StringifyThreatLog(tl)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	exp := `"0001-01-01 00:00:00 +0000 UTC" 1.2.3.4 bot:Detectify adv:threatscan att: /?'PG_SLEEP(100);
`
	if got != exp {
		t.Errorf("stringify miss match, got %s, exp %s", got, exp)
	}
}

func TestURL(t *testing.T) {
	testMatch := func(got, exp string) {
		if got != exp {
			t.Errorf("url miss match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		P string
		Q string
		E string
	}{
		{"", "", ""},
		{"/", "", "/"},
		{"", "a=b", "a=b"},
		{"/", "a=b", "/?a=b"},
	}

	for _, d := range data {
		testMatch(url(d.P, d.Q), d.E)
	}
}
