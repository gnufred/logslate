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

	if tl.Attack != threat.Nil {
		t.Errorf("threat type miss match, got %d, exp %d", tl.Attack, threat.Nil)
	}
}

func TestSTTattack(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "12/Dec/2023:14:04:39 +0000")
	sl := SlateLog{
		Ip:    net.ParseIP("94.131.101.127"),
		Time:  timeEct,
		Path:  "/b/g-t/a.html",
		Query: "item[price][]=1-1661&p=6&price=1-1661&product=desc-1%20waitfor%20delay%20'0:0:15'%20--%20&mode=list&ajax=1",
		Attack: threat.Attack{
			Type: threat.SQL,
		},
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	if tl.Ip.String() != sl.Ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), sl.Ip.String())
	}
	if tl.Attack != sl.Attack.Type {
		t.Errorf("attack miss match, got %d, exp %d", tl.Attack, sl.Attack.Type)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Path != sl.Path {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, sl.Path)
	}
	if tl.Query != sl.Query {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, sl.Query)
	}
}

func TestSTTbot(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	sl := SlateLog{
		Ip:    net.ParseIP("52.167.144.25"),
		Time:  timeEct,
		Bot:   "Google",
		Path:  "/robots.txt",
		Query: "",
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	if tl.Ip.String() != sl.Ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), sl.Ip.String())
	}
	if tl.Attack != sl.Attack.Type {
		t.Errorf("attack miss match, got %d, exp %d", tl.Attack, sl.Attack.Type)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Path != "" {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, "")
	}
	if tl.Query != "" {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, "")
	}
}

func TestSTTad(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	sl := SlateLog{
		Ip:    net.ParseIP("12.45.192.3"),
		Time:  timeEct,
		Path:  "/product/18230/blue",
		Query: "utm_source=facebook",
		Ad:    "facebook",
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	nilIp := net.IP{}
	if tl.Ip.String() != nilIp.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), sl.Ip.String())
	}
	if tl.Attack != sl.Attack.Type {
		t.Errorf("attack miss match, got %d, exp %d", tl.Attack, sl.Attack.Type)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Path != "" {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, "")
	}
	if tl.Query != "" {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, "")
	}
}
