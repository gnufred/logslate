package ngxprs

import (
    "bufio"
    "log"
    "testing"
    "time"
    "os"
)


func TestParseEmpty(t *testing.T) {
    expected := LogEntry{}
    got := Parse("")
    if got != expected {
        log.Fatalf("Parse returned %+v, expected %+v", got, expected)
    }
}

func TestParseUptimeRobot(t *testing.T) {
    line := "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\""

    timeEct, _ := time.Parse(ngxTimeLayout, "05/Jan/2023:07:07:54 +0000")
    expected := LogEntry{
        RemoteAddr: "208.115.199.20",
        RemoteUser: "-",
        TimeLocal: timeEct, 
        RequestMethod: "HEAD",
        RequestUrl: "/?monitoring=uptimerobot",
        RequestProtocol: "HTTP/1.1",
        Status: 200,
        BytesSent: 0,
        HttpReferer: "https://moovadapt.fr/?monitoring=uptimerobot",
        HttpUserAgent: "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
        GzipRatio: 2.8,
    }

    got := Parse(line)

    if got != expected {
        log.Fatalf("Parse returned %+v, expected %+v", got, expected)
    }
}

func TestParseMultiLines(t *testing.T) {
    file, err := os.Open("testdata/access.log")
    if err != nil {
        t.Fatalf("Opening test data failed with %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        Parse(scanner.Text())
    }
}
