package root

import (
	"errors"
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/testing/mock"
	nginx "gitlab.com/gnufred/nxparse"
)

type testData struct {
	file   string
	expect Parsed
}

// Must fail given there's no data in os.Stdin.
func TestNoData(t *testing.T) {
	a := args.Get()
	_, err := ParseStdin(a)
	if !errors.Is(err, ErrNoInputData) {
		t.Errorf("Did not get the expected error")
	}
}

// Must fail given the database file already exists.
func TestDbExists(t *testing.T) {
	input, _ := os.ReadFile("testdata/3ok3broken.log")
	mock.Stdin(input)

	dbPath := "testdata/exist.db"
	os.Create(dbPath)
	defer os.Remove(dbPath)

	a := args.Get()
	a.Db = dbPath
	_, err := ParseStdin(a)
	if errors.Is(err, ErrDbCreateFail) {
		t.Errorf("Did not get the expected error")
	}
}

// Must fail given no --log_format is specified.
func TestNoFormat(t *testing.T) {
	a := args.Get()
	_, err := ParseStdin(a)
	if !errors.Is(err, ErrNoFormat) {
		t.Errorf("Did not get the expected error")
	}
}

// Must fail given it can't parse the input.
func TestParseErrorLimit(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined

	file := "testdata/limitreach.log"
	input, _ := os.ReadFile(file)
	mock.Stdin(input)

	_, err := ParseStdin(a)
	if !errors.Is(err, ErrLimitReach) {
		t.Errorf("Did not get the expected error")
	}
}

func TestNormalLogs(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 20}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 50}},
		},
	)
}

func TestBrokenLogs(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 3}},
		},
	)
}

func TestAttackDetections(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/attacks10.log", Parsed{20, 0, 0, 6, 20}},
		},
	)
}

func TestDatabase(t *testing.T) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	a.Db = ":memory:"
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0, 0, 20}},
			{"testdata/3ok3broken.log", Parsed{3, 0, 3, 0, 3}},
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 50}},
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
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 50}},
		})

	// Should not return anything.
	a.SetSinceUntil("2000-01-01 00:00", "2000-01-01 00:01")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{0, 50, 0, 0, 0}},
		})

	// Should return all lines, first and last should match.
	a.SetSinceUntil("07/Apr/2023:13:40:26 +0000", "14/Apr/2023:09:40:59 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0, 0, 50}},
		})

	// Should return all lines, but the first and last.
	a.SetSinceUntil("07/Apr/2023:13:40:27 +0000", "14/Apr/2023:09:40:58 +0000")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{48, 2, 0, 0, 48}},
		})

	a.SetSinceUntil("13/Apr/2023", "")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{9, 41, 0, 0, 9}},
		})

	a.SetSinceUntil("", "13/Apr/2023")
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{41, 9, 0, 0, 41}},
		})
}

func testInputDataLoop(t *testing.T, a args.Args, d []testData) {
	for _, data := range d {
		testInputData(t, a, data.file, data.expect)
	}
}

func testInputData(t *testing.T, a args.Args, file string, expect Parsed) {
	input, _ := os.ReadFile(file)
	mock.Stdin(input)

	got, err := ParseStdin(a)
	if err != nil {
		t.Error(err)
	}

	if expect.Ok != got.Ok {
		t.Errorf("ok: expect %d, got %d", expect.Ok, got.Ok)
	}

	if expect.Skip != got.Skip {
		t.Errorf("skip: expect %d, got %d", expect.Skip, got.Skip)
	}

	if expect.Err != got.Err {
		t.Errorf("err: expect %d, got %d", expect.Err, got.Err)
	}

	if expect.Atk != got.Atk {
		t.Errorf("atk: expect %d, got %d", expect.Atk, got.Atk)
	}

	if expect.Bot != got.Bot {
		t.Errorf("bot: expect %d, got %d", expect.Bot, got.Bot)
	}
}

func BenchmarkAttackDetections(b *testing.B) {
	a := args.Get()
	a.LogFrmt = nginx.Combined
	for i := 0; i < b.N; i++ {
		benchInputData(b, a, "testdata/attacks10.log")
	}
}

func benchInputData(b *testing.B, a args.Args, file string) {
	input, _ := os.ReadFile(file)
	mock.Stdin(input)

	_, err := ParseStdin(a)
	if err != nil {
		b.Error(err)
	}
}
