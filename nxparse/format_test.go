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

func testParseFormat(t *testing.T, frmt string, exp Format) {
	got := parseFormat(frmt)
	for i := 0; i <= len(exp.Vars); i++ {
		if got.Vars[i] != exp.Vars[i] {
			t.Fatalf("miss match: got, exp\n%v\n%v", got, exp)
		}
	}
}
