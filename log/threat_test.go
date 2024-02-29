package log

import (
	"net"
	"os"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/test"
	"gitlab.com/gnufred/logslate/threat"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestSTTnil(t *testing.T) {
	sl := SlateLog{}
	tl := SlateToThreat(sl)

	if tl.Attack != sl.Attack {
		t.Errorf("attack miss match, got %v, exp nil", tl.Attack)
	}
}

func TestSTTattack(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "12/Dec/2023:14:04:39 +0000")
	sl := SlateLog{
		Detection: threat.Detection{
			Attack: &threat.Attack{
				Type: threat.SQL,
			},
		},
		Ip:    net.ParseIP("94.131.101.127"),
		Time:  timeEct,
		Path:  "/b/g-t/a.html",
		Query: "item[price][]=1-1661&p=6&price=1-1661&product=desc-1%20waitfor%20delay%20'0:0:15'%20--%20&mode=list&ajax=1",
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	if tl.Ip.String() != sl.Ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), sl.Ip.String())
	}
	if tl.Attack.Type != sl.Attack.Type {
		t.Errorf("attack miss match, got %d, exp %d", tl.Attack.Type, sl.Attack.Type)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Query != sl.Query {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, sl.Query)
	}
	if tl.Path != sl.Path {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, sl.Path)
	}
}

func TestSTTbot(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	sl := SlateLog{
		Detection: threat.Detection{
			Bot: "Google",
		},
		Ip:    net.ParseIP("52.167.144.25"),
		Time:  timeEct,
		Path:  "/robots.txt",
		Query: "debug=true",
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	if tl.Ip.String() != sl.Ip.String() {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), sl.Ip.String())
	}
	if tl.Attack != sl.Attack {
		t.Errorf("attack miss match, got %v, exp nil", tl.Attack)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Query != "" {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, "")
	}
	if tl.Path != "" {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, "")
	}
}

func TestSTTad(t *testing.T) {
	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "26/Dec/2023:14:38:02 +0000")
	sl := SlateLog{
		Detection: threat.Detection{
			Ad: "facebook",
		},
		Ip:    net.ParseIP("12.45.192.3"),
		Time:  timeEct,
		Path:  "/product/18230/blue",
		Query: "utm_source=facebook",
	}

	tl := SlateToThreat(sl)

	if tl.Time != sl.Time {
		t.Errorf("time miss match, got %s, exp %s", tl.Time, sl.Time)
	}
	if tl.Ip != nil {
		t.Errorf("ip miss match, got %s, exp %s", tl.Ip.String(), "<nil>")
	}
	if tl.Attack != sl.Attack {
		t.Errorf("attack miss match, got %v, exp nil", tl.Attack)
	}
	if tl.Bot != sl.Bot {
		t.Errorf("bot miss match, got %s, exp %s", tl.Bot, sl.Bot)
	}
	if tl.Ad != sl.Ad {
		t.Errorf("ad miss match, got %s, exp %s", tl.Ad, sl.Ad)
	}
	if tl.Query != "" {
		t.Errorf("query miss match, got %s, exp %s", tl.Query, "")
	}
	if tl.Path != "" {
		t.Errorf("path miss match, got %s, exp %s", tl.Path, "")
	}
}

func TestParseThreatFail(t *testing.T) {
	_, err := ParseThreat(string(";"))
	test.ErrorExpected(t, err, "invalid character")
}

func TestParseThreat(t *testing.T) {
	file := "testdata/threat.json"
	line, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	tl, err := ParseThreat(string(line))
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	testMatch := func(t *testing.T, in, exp string) {
		if in != exp {
			t.Errorf("threat log data miss match, got %s, exp %s", in, exp)
		}
	}

	testMatch(t, tl.Bot, "NodePing")
	testMatch(t, tl.Time.String(), "2024-02-28 21:28:15 +0000 UTC")
	testMatch(t, tl.Ip.String(), "200.20.240.40")
}
