package slate

import (
    "testing"
    "time"
    "log"

    "gitlab.com/gnufred/logslate/parse/nginx"
)

func TestParseNginx(t *testing.T) {
    line := "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\""

    timeEct, _ := time.Parse(nginx.TimeLayout, "05/Jan/2023:07:07:54 +0000")
    expected := Log{
        Ip: "208.115.199.20",
        User: "-",
        Time: timeEct, 
        Method: "HEAD",
        Path: "/",
        Query: "monitoring=uptimerobot",
        Proto: "HTTP/1.1",
        Status: 200,
        Sent: 0,
        Referer: "https://moovadapt.fr/?monitoring=uptimerobot",
        Agent: "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
    }

    ngxLog, err := nginx.Parse(line)
    if err != nil {
        t.Fatalf("nginx.Parse() error %s", err)
    }
    got := ParseNginx(ngxLog)

    if got != expected {
        log.Fatalf("Parse returned %+v, expected %+v", got, expected)
    }
}
