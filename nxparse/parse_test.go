package nxparse

import (
	"bufio"
	"errors"
	"os"
	"testing"
	"time"
)

// We need to test the Combined parsing both in nxparse and in parse given it's cached.
// Other formats don't need duplicated test. We can test them in nxparse given the format
// will be parsed, cached, and the line will also be parsed.
func TestParseLine(t *testing.T) {
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

func TestParseCombinedNormal(t *testing.T) {
	file, err := os.Open("testdata/combined_normal.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	frmt, _ := cache.get(Combined)

	for scanner.Scan() {
		_, err := parseLine(frmt, scanner.Text())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParseCombinedTruncated(t *testing.T) {
	file, err := os.Open("testdata/combined_truncated.log")
	if err != nil {
		t.Fatal("opening test data failed: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	frmt, _ := cache.get(Combined)

	for scanner.Scan() {
		txt := scanner.Text()
		_, err := parseLine(frmt, txt)
		if !errors.Is(err, ErrIpConvFail) && !errors.Is(err, ErrSepNotFound) {
			t.Fatalf("truncated line did not generate an error\n%s\n%s\n", txt, err)
		}
	}
}

func TestAssignToEntry(t *testing.T) {
	e := Entry{}

	checkMatch := func(t *testing.T, got, exp interface{}) {
		if got != exp {
			t.Fatalf("wrong value assigned to entry: got %v, exp %v", got, exp)
		}
	}

	err := assignToEntry(&e, "non_existent_field", "")
	if !errors.Is(err, ErrVarNotExist) {
		t.Fatalf("assignment to non existent variable didn't fail as expected")
	}

	assignToEntry(&e, "body_bytes_sent", "5234")
	checkMatch(t, e.BodyBytesSent, int64(5234))

	assignToEntry(&e, "bytes_sent", "5234068")
	checkMatch(t, e.BytesSent, int64(5234068))

	assignToEntry(&e, "connection", "aj28zkj10123yihADSFlkadf")
	checkMatch(t, e.Connection, "aj28zkj10123yihADSFlkadf")

	assignToEntry(&e, "connection_requests", "321")
	checkMatch(t, e.ConnectionRequests, int64(321))

	assignToEntry(&e, "http_referer", "https//www.google.com")
	checkMatch(t, e.Http["referer"], "https//www.google.com")

	ua := "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	assignToEntry(&e, "http_user_agent", ua)
	checkMatch(t, e.Http["user_agent"], ua)

	assignToEntry(&e, "msec", "4027")
	checkMatch(t, e.Msec, int64(4027))

	assignToEntry(&e, "pipe", "p")
	checkMatch(t, e.Pipe, bool(true))
	assignToEntry(&e, "pipe", ".")
	checkMatch(t, e.Pipe, bool(false))
	// bad value
	err = assignToEntry(&e, "pipe", "e")
	if !errors.Is(err, ErrConverFail) {
		t.Fatalf("pipe bad assigned didn't fail as expected")
	}

	assignToEntry(&e, "remote_addr", "234.28.148.10")
	checkMatch(t, e.RemoteAddr.String(), "234.28.148.10")

	assignToEntry(&e, "remote_user", "bad_actor_net")
	checkMatch(t, e.RemoteUser, "bad_actor_net")

	assignToEntry(&e, "request", "/wp-login.php")
	checkMatch(t, e.Request, "/wp-login.php")

	assignToEntry(&e, "request_length", "30238742")
	checkMatch(t, e.RequestLength, int64(30238742))

	assignToEntry(&e, "request_time", "742")
	checkMatch(t, e.RequestTime, int64(742))

	assignToEntry(&e, "status", "200")
	checkMatch(t, e.Status, int64(200))

	assignToEntry(&e, "time_iso8601", "2022-09-27 18:00:00.000")
	tm, _ := time.Parse(TimeLayoutIso8601, "2022-09-27 18:00:00.000")
	checkMatch(t, e.TimeIso8601, tm)

	assignToEntry(&e, "time_local", "27/Apr/2023:09:49:32 +0000")
	tm, _ = time.Parse(TimeLayoutLocal, "27/Apr/2023:09:49:32 +0000")
	checkMatch(t, e.TimeLocal, tm)
}

// === benchmarks ===

func BenchmarkParseLine(b *testing.B) {
	lineCombined := "216.245.221.84 - - [27/Apr/2023:09:49:32 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://example.com/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""
	frmt, _ := cache.get(Combined)

	for i := 0; i < b.N; i++ {
		_, err := parseLine(frmt, lineCombined)
		if err != nil {
			b.Fatal(err)
		}
	}
}
