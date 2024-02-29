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

func init() {
	args.Init()
}

// Must fail given there's no data in os.Stdin.
func TestNoData(t *testing.T) {
	a := args.Singleton()
	test.DisableOutput()
	_, err := ParseStdin(a, "test")
	test.EnableOutput()
	test.ErrorExpected(t, err, ErrNoInputData.Error())
}

// Must fail given the database file already exists.
func TestErrDbCreateFail(t *testing.T) {
	a := args.Singleton()
	a.Parse.Db = "root_test.db"
	a.Parse.LogFrmt = nginx.Combined

	_, err := os.Create(a.Parse.Db)
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
	_, err = ParseStdin(a, "test")

	// Test: https://gitlab.com/gnufred/logslate/-/issues/106
	// Don't use strings.Contais(). The error must be an exact match.
	test.ErrorExpected(t, err,
		"unable to create database: file root_test.db already exists",
	)
	test.EnableOutput()

	err = os.Remove(a.Parse.Db)
	if err != nil {
		t.Error(err)
	}

	a.Parse.Db = ""
}

// Must handle buffio.Scan() ErrTooLong properly.
// Cannot test: MockStdin() need to change the os.Pipe buffer at runtime.
// See test.MockStdin comments for more information.
/*
func TestParseTokenTooLong(t *testing.T) {
	a := args.Singleton()
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
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined

	file := "testdata/limitreach.log"
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	_, err = ParseStdin(a, "test")
	test.EnableOutput()
	test.ErrorExpected(t, err, ErrLimitReach.Error())
}

func TestNormalLogs(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		},
	)
}

func TestBrokenLogs(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 0, 3}},
		},
	)
}

func TestAttackDetections(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
		},
	)
}

func TestAdDetections(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
}

func TestDatabase(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	a.Parse.Db = ":memory:"
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 0, 3}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		},
	)
}

func TestTimeArgs(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined

	// Empty values should return all results.
	a.Time.SetSince("")
	a.Time.SetUntil("")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		})

	// Should not return anything.
	a.Time.SetSince("2000-01-01 00:00")
	a.Time.SetUntil("2000-01-01 00:01")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{0, 50, 0, 0, 0, 0}},
		})

	// Should return all lines, first and last should match.
	a.Time.SetSince("07/Apr/2023:13:40:26 +0000")
	a.Time.SetUntil("14/Apr/2023:09:40:59 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 0, 50}},
		})

	// Should return all lines, but the first and last.
	a.Time.SetSince("07/Apr/2023:13:40:27 +0000")
	a.Time.SetUntil("14/Apr/2023:09:40:58 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{48, 2, 0, 0, 0, 48}},
		})

	a.Time.SetSince("13/Apr/2023")
	a.Time.SetUntil("")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{9, 41, 0, 0, 0, 9}},
		})

	a.Time.SetSince("")
	a.Time.SetUntil("13/Apr/2023")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{41, 9, 0, 0, 0, 41}},
		})
}

func TestReport(t *testing.T) {
	a := args.Singleton()
	a.Report.Path = "testdata/report.txt"
	a.Report.Format = args.TEXT
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
	err := os.Remove(a.Report.Path)
	if err != nil {
		t.Error(err)
	}
	a.Report.Path = ""
}

func TestThreatsScreen(t *testing.T) {
	a := args.Singleton()
	a.Threat.Path = "screen"
	a.Threat.Format = args.SCREEN
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
		},
	)
}

func TestThreatsText(t *testing.T) {
	a := args.Singleton()
	a.Threat.Path = "testdata/threat.text"
	a.Threat.Format = args.TEXT
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
			{"testdata/utm.log", Parsed{6, 0, 0, 0, 6, 0}},
		},
	)
	err := os.Remove(a.Threat.Path)
	if err != nil {
		t.Error(err)
	}
}

func TestThreatsDebug(t *testing.T) {
	a := args.Singleton()
	a.Debug.Verbosity = 3
	a.Threat.Path = "screen"
	a.Threat.Format = args.SCREEN
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 0, 20}},
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 0, 20}},
		},
	)
}

func TestTranslateError(t *testing.T) {
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/translateerror.log", Parsed{1, 0, 2, 0, 0, 0}},
		},
	)
}

func TestReportCommand(t *testing.T) {
	a := args.Singleton()
	a.Command = args.REPARSE
	a.Parse.LogFrmt = nginx.Combined
	a.Threat.Path = ""
	a.Threat.Format = 0
	a.Time.SetSince("2000")
	a.Time.SetUntil("2030")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/threat.json", Parsed{1, 0, 0, 0, 0, 1}},
			{"testdata/invalid.json", Parsed{0, 0, 1, 0, 0, 0}},
		},
	)
}

func testInputDataLoop(t *testing.T, a *args.ArgsV2, d []testData) {
	for _, data := range d {
		testInputData(t, a, data.file, data.expect)
	}
}

func testInputData(t *testing.T, a *args.ArgsV2, file string, exp Parsed) {
	input, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	test.MockStdin(input)

	test.DisableOutput()
	got, err := ParseStdin(a, "test")
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
	a := args.Singleton()
	a.Parse.LogFrmt = nginx.Combined
	test.DisableOutput()

	input, _ := os.ReadFile("testdata/attacks10.log")
	test.MockStdin(input)

	for i := 0; i < b.N; i++ {
		benchInputData(b, a)
	}
	test.EnableOutput()
}

func benchInputData(b *testing.B, a *args.ArgsV2) {
	_, err := ParseStdin(a, "test")
	if err != nil {
		b.Error(err)
	}
}
