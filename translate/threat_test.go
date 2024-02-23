package translate

import (
	"net"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/threat"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestSTTnil(t *testing.T) {
	sl := SlateLog{}
	tl := SlateToThreat(sl)

	if tl.Type != "" {
		t.Errorf("threat type miss match, got %s, exp %s", tl.Type, "")
	}
}

func TestSTTattack(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "12/Dec/2023:14:04:39 +0000")
	ip := net.ParseIP("94.131.101.127")
	path := "/b/g-t/a.html"
	query := "item[price][]=1-1661&p=6&price=1-1661&product=desc-1%20waitfor%20delay%20'0:0:15'%20--%20&mode=list&ajax=1"
	attack := "sql_inject"
	sl := SlateLog{
		Ip:    ip,
		Time:  timeEct,
		Path:  path,
		Query: query,
		Attack: threat.Attack{
			Name: attack,
		},
	}

	tl := SlateToThreat(sl)

	if tl.Time != timeEct {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, timeEct)
	}
	if tl.Ip.String() != ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), ip.String())
	}
	if tl.Type != attack {
		t.Errorf("threat type miss match, got %s, exp %s", tl.Type, attack)
	}
	exp := path + "?" + query
	if tl.Info != exp {
		t.Errorf("info miss match, got %s, exp %s", tl.Type, exp)
	}
}

func TestAttackInfo(t *testing.T) {
	var got, exp, q, p string
	q = "index.html"
	p = "hello=world"

	got = attackInfo(q, p)
	exp = q + "?" + p
	if got != exp {
		t.Errorf("attack info miss match, got %s, exp %s", got, exp)
	}

	got = attackInfo("", p)
	exp = p
	if got != exp {
		t.Errorf("attack info miss match, got %s, exp %s", got, exp)
	}

	got = attackInfo(q, "")
	exp = q
	if got != exp {
		t.Errorf("attack info miss match, got %s, exp %s", got, exp)
	}
}

func TestSTTbot(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	ip := net.ParseIP("52.167.144.25")
	bot := "Google"
	sl := SlateLog{
		Ip:   ip,
		Time: timeEct,
		Bot:  bot,
	}

	tl := SlateToThreat(sl)

	if tl.Time != timeEct {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, timeEct)
	}
	if tl.Ip.String() != ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), ip.String())
	}
	if tl.Type != "bot" {
		t.Errorf("threat type miss match, got %s, exp %s", tl.Type, "bot")
	}
	if tl.Info != bot {
		t.Errorf("info miss match, got %s, exp %s", tl.Type, bot)
	}
}

func TestSTTad(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	ip := net.IP{}
	path := "/index.html"
	query := "utm_source=facebook"
	ad := "facebook"
	sl := SlateLog{
		Ip:    ip,
		Time:  timeEct,
		Path:  path,
		Query: query,
		Ad:    ad,
	}

	tl := SlateToThreat(sl)

	if tl.Time != timeEct {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, timeEct)
	}
	if tl.Ip.String() != ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), ip.String())
	}
	if tl.Type != "ad" {
		t.Errorf("threat type miss match, got %s, exp %s", tl.Type, "ad")
	}
	if tl.Info != ad {
		t.Errorf("info miss match, got %s, exp %s", tl.Type, ad)
	}
}
