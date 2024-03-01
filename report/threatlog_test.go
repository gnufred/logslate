package report

import (
	"net"
	"testing"

	"gitlab.com/gnufred/logslate/log"
	"gitlab.com/gnufred/logslate/threat"
)

func TestNewThreatLog(t *testing.T) {
	tl := NewTL()

	if len(tl.Ads) != 0 {
		t.Error("tl.Ads not equal 0")
	}
	if len(tl.Bots) != 0 {
		t.Error("tl.Bots not equal 0")
	}
	if len(tl.Attacks) != 0 {
		t.Error("tl.Attacks not equal 0")
	}
	if len(tl.Attackers) != 0 {
		t.Error("tl.Attackers not equal 0")
	}
}

func TestParseAds(t *testing.T) {
	in := "Pinterest"
	tl := log.ThreatLog{Detection: threat.Detection{Ad: in}}

	tlr := NewTL()
	tlr.Parse(tl)
	if tlr.Ads[in] != 1 {
		t.Error("ad count miss match")
	}
}

func TestParseBots(t *testing.T) {
	in := "Google"
	tl := log.ThreatLog{Detection: threat.Detection{Bot: in}}

	tlr := NewTL()
	tlr.Parse(tl)
	if tlr.Bots[in] != 1 {
		t.Error("bot count miss match")
	}
}

func TestParseAttacks(t *testing.T) {
	ip := net.IP("1.2.3.4")
	in := threat.XSS
	tl := log.ThreatLog{
		Detection: threat.Detection{
			Attack: &threat.Attack{0, in, ""},
		},
		Ip: ip,
	}

	tlr := NewTL()
	tlr.Parse(tl)
	if tlr.Attacks[in] != 1 {
		t.Error("attack count miss match")
	}
}
