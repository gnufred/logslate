package nxparse

import (
	"testing"
)

func TestParseVerySimpleFormat(t *testing.T) {
	testParseFormat(t, "$one $two ", Format{
		Vars: map[int]FormatVar{
			0: {"one", " "},
			1: {"two", " "},
		}},
	)
}

func TestParseNewLine(t *testing.T) {
	testParseFormat(t, `$one
$two `, Format{
		Vars: map[int]FormatVar{
			0: {"one", "\n"},
			1: {"two", " "},
		}},
	)
}

func TestSpecialCaseStartingSeperator(t *testing.T) {
	testParseFormat(t, "[$one] $two ", Format{
		Vars: map[int]FormatVar{
			0: {"", "["},
			1: {"one", "] "},
			2: {"two", " "},
		}},
	)
}

func TestSpecialCaseNoEndSeperator(t *testing.T) {
	testParseFormat(t, "$one $two", Format{
		Vars: map[int]FormatVar{
			0: {"one", " "},
			1: {"two", ""},
		}},
	)
}

func TestBothSpecialCase(t *testing.T) {
	testParseFormat(t, "[$one] $two", Format{
		Vars: map[int]FormatVar{
			0: {"", "["},
			1: {"one", "] "},
			2: {"two", ""},
		}},
	)
}

func TestParseCombined(t *testing.T) {
	frmt, _ := cache.get(Combined)
	testParseFormat(t, Combined, frmt)
}

func TestParseCombined8601(t *testing.T) {
	testParseFormat(t,
		"$remote_addr - $remote_user [$time_iso8601] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\"",
		Format{Vars: map[int]FormatVar{
			0: {"remote_addr", " - "},
			1: {"remote_user", " ["},
			2: {"time_iso8601", "] \""},
			3: {"request", "\" "},
			4: {"status", " "},
			5: {"body_bytes_sent", " \""},
			6: {"http_referer", "\" \""},
			7: {"http_user_agent", "\""},
		}})
}

func testParseFormat(t *testing.T, frmt string, exp Format) {
	got := parseFormat(frmt)
	for i := 0; i <= len(exp.Vars); i++ {
		if got.Vars[i] != exp.Vars[i] {
			t.Fatalf("miss match: got, exp\n%v\n%v", got, exp)
		}
	}
}
