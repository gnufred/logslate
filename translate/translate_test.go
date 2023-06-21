package translate

import (
	"net"
	"strings"
	"testing"
	"time"

	nginx "gitlab.com/gnufred/nxparse"
)

func TestNginxCombined(t *testing.T) {
	lineCombined := "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\""

	timeEct, _ := time.Parse(nginx.TimeLayoutLocal, "05/Jan/2023:07:07:54 +0000")
	exp := Log{
		Ip:      net.ParseIP("208.115.199.20"),
		User:    "-",
		Time:    timeEct,
		Method:  "HEAD",
		Path:    "/",
		Query:   "monitoring=uptimerobot",
		Proto:   "HTTP/1.1",
		Status:  200,
		Sent:    0,
		Referer: "https://moovadapt.fr/?monitoring=uptimerobot",
		Agent:   "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
	}

	ngxLog, err := nginx.Parse(nginx.Combined, lineCombined)
	if err != nil {
		t.Errorf("Parse() error %s", err)
	}
	got, err := Nginx(ngxLog)
	if err != nil {
		t.Error(err)
	}
	compareStruct(t, got, exp)
}

func TestNginxIso8601(t *testing.T) {

	frmtCombIso8601 := "$remote_addr - $remote_user [$time_iso8601] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\""
	lineCombIso8601 := "82.100.10.100 - - [2023-05-15T09:25:22+00:00] \"POST /locator/index/ HTTP/1.1\" 302 31 \"https://www.example.com/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36\""

	timeEct, _ := time.Parse(nginx.TimeLayoutIso8601, "2023-05-15T09:25:22+00:00")
	exp := Log{
		Ip:      net.ParseIP("82.100.10.100"),
		User:    "-",
		Time:    timeEct,
		Method:  "POST",
		Path:    "/locator/index/",
		Query:   "",
		Proto:   "HTTP/1.1",
		Status:  302,
		Sent:    31,
		Referer: "https://www.example.com/",
		Agent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
	}

	ngxLog, err := nginx.Parse(frmtCombIso8601, lineCombIso8601)
	if err != nil {
		t.Errorf("Parse() error %s", err)
	}
	got, err := Nginx(ngxLog)
	if err != nil {
		t.Error(err)
	}
	compareStruct(t, got, exp)
}

// https://gitlab.com/gnufred/logslate/-/issues/92
func TestIssue92Fix(t *testing.T) {
	crashString := "GET /rest/V1/directory/countries?_=16869P/1.1"
	_, err := splitNginxRequest(crashString)
	if !strings.Contains(err.Error(), "malformed NGINX request") {
		t.Errorf("did not get expected error, got %s", err)
	}

}

func compareStruct(t *testing.T, got, exp Log) {
	checkMatch(t, got.Ip.String(), exp.Ip.String())
	checkMatch(t, got.User, exp.User)
	checkMatch(t, got.Time, exp.Time)
	checkMatch(t, got.Method, exp.Method)
	checkMatch(t, got.Path, exp.Path)
	checkMatch(t, got.Query, exp.Query)
	checkMatch(t, got.Proto, exp.Proto)
	checkMatch(t, got.Status, exp.Status)
	checkMatch(t, got.Sent, exp.Sent)
	checkMatch(t, got.Referer, exp.Referer)
	checkMatch(t, got.Agent, exp.Agent)
}

func checkMatch(t *testing.T, got, exp interface{}) {
	if got != exp {
		t.Errorf("wrong value assigned to entry: got %v, exp %v", got, exp)
	}
}
