package nginx

import (
    "bufio"
    "log"
    "testing"
    "time"
    "os"
)


func TestParseEmpty(t *testing.T) {
    expected := Log{}
    got, _ := Parse("")
    if got != expected {
        log.Fatalf("Parse returned %+v, expected %+v", got, expected)
    }
}

/* This simulate cases when the access.log have been truncated in the middle of a line.
   On such case, the first line will be missing a random amount of characters. */
func TestTruncatedLogs(t *testing.T) {
    lines := []string{
        // panic: " - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        " [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "ing=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        " 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "\"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "t\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "zilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\"",
        "; http://www.uptimerobot.com/)\" \"2.80\"",
        "timerobot.com/)\" \"2.80\"",
        "\"2.80\"",
        "80\"",
        "\"",
        // should error but doesnt: "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2",
        "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
    }
    for _, line := range lines {
        checkBrokenLine(t, line)
    }
}

func checkBrokenLine(t *testing.T, line string) {
    expected := Log{}

    got, err := Parse(line)
    if got != expected {
        t.Fatal("Broken line does not generate Log{}")
    }
    if err == nil {
        t.Fatal("Broken line does not generate an error")
    }
}

func TestParseUptimeRobot(t *testing.T) {
    line := "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\" \"2.80\""

    timeEct, _ := time.Parse(TimeLayout, "05/Jan/2023:07:07:54 +0000")
    expected := Log{
        RemoteAddr: "208.115.199.20",
        RemoteUser: "-",
        TimeLocal: timeEct, 
        Request: "HEAD /?monitoring=uptimerobot HTTP/1.1",
        Status: 200,
        BytesSent: 0,
        HttpReferer: "https://moovadapt.fr/?monitoring=uptimerobot",
        HttpUserAgent: "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
        GzipRatio: 2.8,
    }

    got, err := Parse(line)
    if err != nil {
        t.Fatalf("Parse() error %s", err)
    }

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
