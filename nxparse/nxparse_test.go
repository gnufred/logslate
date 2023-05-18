package nxparse

import (
	"bufio"
	"errors"
	"os"
	"testing"
	"time"
)

// Test an empty line, which should return zero value for all fields
func TestMainEmptyStuff(t *testing.T) {
	_, err := Parse("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = Parse(Combined, "")
	if err != nil && !errors.Is(err, ErrSepNotFound) {
		t.Fatalf("got error: %s\nexpected: %s", err, ErrSepNotFound)
	}
}

// Test a single line of `combined` log format for the expected values.
// We need to test the Combined parsing both in nxparse and in parse given it's cached.
// Other formats don't need duplicated test. We can test them in here given the format
// will be parsed, cached, and the line will also be parsed.
func TestMainCombinedOneLine(t *testing.T) {
	checkMatch := func(t *testing.T, got, exp interface{}) {
		if got != exp {
			t.Fatalf("bad log variable value: got %v, exp %v", got, exp)
		}
	}

	lineCombined := "216.245.221.84 - - [27/Apr/2023:09:49:32 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://example.com/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""

	e, err := Parse(Combined, lineCombined)
	if err != nil {
		t.Fatal(err)
	}

	checkMatch(t, e.RemoteAddr.String(), "216.245.221.84")
	checkMatch(t, e.RemoteUser, "-")
	tm, _ := time.Parse(TimeLayoutLocal, "27/Apr/2023:09:49:32 +0000")
	checkMatch(t, e.TimeLocal, tm)
	checkMatch(t, e.Request, "HEAD /?monitoring=uptimerobot HTTP/1.1")
	checkMatch(t, e.Status, int64(200))
	checkMatch(t, e.BodyBytesSent, int64(0))
	checkMatch(t, e.Http["referer"], "https://example.com/?monitoring=uptimerobot")
	checkMatch(t, e.Http["user_agent"], "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)")
}

func TestMainCombIso8601OneLine(t *testing.T) {
	checkMatch := func(t *testing.T, got, exp interface{}) {
		if got != exp {
			t.Fatalf("bad log variable value: got %v, exp %v", got, exp)
		}
	}

	frmtCombIso8601 := "$remote_addr - $remote_user [$time_iso8601] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\""
	lineCombIso8601 := "82.100.10.100 - - [2023-05-15T09:25:22+00:00] \"POST /locator/index/ HTTP/1.1\" 302 31 \"https://www.example.com/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36\""

	e, err := Parse(frmtCombIso8601, lineCombIso8601)
	if err != nil {
		t.Fatal(err)
	}

	checkMatch(t, e.RemoteAddr.String(), "82.100.10.100")
	checkMatch(t, e.RemoteUser, "-")
	tm, _ := time.Parse(TimeLayoutIso8601, "2023-05-15T09:25:22+00:00")
	checkMatch(t, e.TimeIso8601, tm)
	checkMatch(t, e.Request, "POST /locator/index/ HTTP/1.1")
	checkMatch(t, e.Status, int64(302))
	checkMatch(t, e.BodyBytesSent, int64(31))
	checkMatch(t, e.Http["referer"], "https://www.example.com/")
	checkMatch(t, e.Http["user_agent"], "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
}

func TestMainCombinedNormal(t *testing.T) {
	file, err := os.Open("testdata/combined_normal.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		_, err := Parse(Combined, scanner.Text())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestMainCombinedTruncated(t *testing.T) {
	file, err := os.Open("testdata/combined_truncated.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txt := scanner.Text()
		_, err := Parse(Combined, txt)
		if !errors.Is(err, ErrIpConvFail) && !errors.Is(err, ErrSepNotFound) {
			t.Fatalf("truncated line did not generate an error\n%s\n", txt)
		}
	}
}

// === benchmarks ===

func BenchmarkParse(b *testing.B) {
	lineCombined := "216.245.221.84 - - [27/Apr/2023:09:49:32 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://example.com/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""

	for i := 0; i < b.N; i++ {
		_, err := Parse(Combined, lineCombined)
		if err != nil {
			b.Fatal(err)
		}
	}
}
