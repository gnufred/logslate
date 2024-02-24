package root

import (
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

type testData struct {
	file   string
	expect Parsed
}

// Must fail given there's no data in os.Stdin.
func TestNoData(t *testing.T) {
	a := args.Get()
	test.DisableOutput()
	_, err := ParseStdin(a)
	test.EnableOutput()
	test.ErrorExpected(t, err, ErrNoInputData.Error())
}

// Must fail given the database file already exists.
func TestErrDbCreateFail(t *testing.T) {
	a := args.Get()
	a.Db = "root_test.db"
	a.LogFrmt = nginx.Combined

	_, err := os.Create(a.Db)
	if err != nil {
		t.Error(err)
	}

	file := "testdata/20ok.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	_, err = ParseStdin(a)

	// Test: https://gitlab.com/gnufred/logslate/-/issues/106
	// Don't use strings.Contais(). The error must be an exact match.
	test.ErrorExpected(t, err,
		"unable to create database: file root_test.db already exists",
	)
	test.EnableOutput()

	err = os.Remove(a.Db)
	if err != nil {
		t.Error(err)
	}
}

// Must handle buffio.Scan() ErrTooLong properly.
// Cannot test: MockStdin() need to change the os.Pipe buffer at runtime.
// See test.MockStdin comments for more information.
/*
func TestParseTokenTooLong(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined

	file := "testdata/toolong.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	_, err := ParseStdin(a)
	test.EnableOutput()
	test.ErrorExpected(err, bufio.ErrTooLong.Error())
}
*/

// Must fail given it can't parse the input.
func TestParseErrorLimit(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined

	file := "testdata/limitreach.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	_, err = ParseStdin(a)
	test.EnableOutput()
	test.ErrorExpected(t, err, ErrLimitReach.Error())
}

func TestNormalLogs(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		},
	)
}

func TestBrokenLogs(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 0, 3}},
		},
	)
}

func TestAttackDetections(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
		},
	)
}

func TestAdDetections(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
}

func TestDatabase(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	a.Db = ":memory:"
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 0, 3}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		},
	)
}

func TestTimeArgs(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined

	// Empty values should return all results.
	a.SetSinceUntil("", "")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		})

	// Should not return anything.
	a.SetSinceUntil("2000-01-01 00:00", "2000-01-01 00:01")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{0, 50, 0, 0, 0, 0}},
		})

	// Should return all lines, first and last should match.
	a.SetSinceUntil("07/Apr/2023:13:40:26 +0000", "14/Apr/2023:09:40:59 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		})

	// Should return all lines, but the first and last.
	a.SetSinceUntil("07/Apr/2023:13:40:27 +0000", "14/Apr/2023:09:40:58 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{48, 2, 0, 0, 0, 48}},
		})

	a.SetSinceUntil("13/Apr/2023", "")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{9, 41, 0, 0, 0, 9}},
		})

	a.SetSinceUntil("", "13/Apr/2023")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{41, 9, 0, 0, 0, 41}},
		})
}

func TestReport(t *testing.T) {
	a := args.Get()
	a.ReportPath = "testdata/report.txt"
	a.ReportFormat = 2
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
	err := os.Remove(a.ReportPath)
	if err != nil {
		t.Error(err)
	}
}

func TestThreatsScreen(t *testing.T) {
	a := args.Get()
	a.ThreatPath = "screen"
	a.ThreatFormat = 1
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
		},
	)
}

func TestThreatsText(t *testing.T) {
	a := args.Get()
	a.ThreatPath = "testdata/threat.text"
	a.ThreatFormat = 2
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
	err := os.Remove(a.ThreatPath)
	if err != nil {
		t.Error(err)
	}
}

func TestThreatsDebug(t *testing.T) {
	a := args.Get()
	a.Debug = 3
	a.ThreatPath = "screen"
	a.ThreatFormat = 1
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 12, 0, 20}},
		},
	)
}

func TestTranslateError(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/translateerror.log", Parsed{1, 0, 2, 0, 0, 0}},
		},
	)
}

func testInputDataLoop(t *testing.T, a args.Args, d []testData) {
	for _, data := range d {
		testInputData(t, a, data.file, data.expect)
	}
}

func testInputData(t *testing.T, a args.Args, file string, exp Parsed) {
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	got, err := ParseStdin(a)
	test.EnableOutput()
	if err != nil {
		t.Error(err)
	}

	if exp.Ok != got.Ok {
		t.Errorf("ok: %s, exp %d, got %d", file, exp.Ok, got.Ok)
	}

	if exp.Skip != got.Skip {
		t.Errorf("skip: %s, exp %d, got %d", file, exp.Skip, got.Skip)
	}

	if exp.Err != got.Err {
		t.Errorf("err: %s, exp %d, got %d", file, exp.Err, got.Err)
	}

	if exp.Atk != got.Atk {
		t.Errorf("atk: %s, exp %d, got %d", file, exp.Atk, got.Atk)
	}

	if exp.Ad != got.Ad {
		t.Errorf("ad: %s, exp %d, got %d", file, exp.Ad, got.Ad)
	}

	if exp.Bot != got.Bot {
		t.Errorf("bot: %s, exp %d, got %d", file, exp.Bot, got.Bot)
	}
}

// === benchmarks === //

func BenchmarkAttackDetections(b *testing.B) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	test.DisableOutput()

	input, _ := os.ReadFile("testdata/attacks10.log")
	test.MockStdin(input)

	for i := 0; i < b.N; i++ {
		benchInputData(b, a)
	}
	test.EnableOutput()
}

func benchInputData(b *testing.B, a args.Args) {
	_, err := ParseStdin(a)
	if err != nil {
		b.Error(err)
	}
}
