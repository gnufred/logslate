package rootcmd

import (
	"os"
	"testing"
	"time"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/testing/mock"
)

// Must fail given there's no data in os.Stdin
func TestNoData(t *testing.T) {
	a := args.Get()
	_, err := ParseStdin(a)
	if err == nil {
		t.Fatalf("Did not get the expected error")
	}
}

// Must fail given the database file already exists
func TestDbExists(t *testing.T) {
	input, _ := os.ReadFile("testdata/3ok3broken.log")
	mock.Stdin(input)

	dbPath := "testdata/exist.db"
	os.Create(dbPath)
	defer os.Remove(dbPath)

	a := args.Get()
	a.Db = dbPath
	_, err := ParseStdin(a)
	if err == nil {
		t.Fatalf("Did not get the expected error")
	}
}

type testData struct {
	file   string
	expect Parsed
}

func TestNormalLogs(t *testing.T) {
	a := args.Get()
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0}},
			{"testdata/googlebot.log", Parsed{50, 0, 0}},
		},
	)
}

func TestBrokenLogs(t *testing.T) {
	a := args.Get()
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/3ok3broken.log", Parsed{3, 0, 3}},
		},
	)
}

func TestDatabase(t *testing.T) {
	a := args.Get()
	a.Db = "testdata/test.db"
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/20ok.log", Parsed{20, 0, 0}},
			{"testdata/3ok3broken.log", Parsed{3, 0, 3}},
			{"testdata/googlebot.log", Parsed{50, 0, 0}},
		},
	)
}

func TestTimeArgs(t *testing.T) {
	a := args.Get()

	// empty values should return all results
	a.Since = ""
	a.Until = ""
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0}},
		})

	// should not return anything
	a.Since = "2000-01-01 00:00"
	a.Until = "2000-01-01 00:01"
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{0, 50, 0}},
		})

	// should return all lines, first and last should match
	a.Since = "07/Apr/2023:13:40:26 +0000"
	a.Until = "14/Apr/2023:09:40:59 +0000"
	a.StartTime = args.ParseTimestamp(a.Since, args.UnsetTime)
	a.EndTime = args.ParseTimestamp(a.Until, time.Now())
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{50, 0, 0}},
		})

	// should return all lines, but the first and last
	a.Since = "07/Apr/2023:13:40:27 +0000"
	a.Until = "14/Apr/2023:09:40:58 +0000"
	a.StartTime = args.ParseTimestamp(a.Since, args.UnsetTime)
	a.EndTime = args.ParseTimestamp(a.Until, time.Now())
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{48, 2, 0}},
		})

	a.Since = "13/Apr/2023"
	a.Until = ""
	a.StartTime = args.ParseTimestamp(a.Since, args.UnsetTime)
	a.EndTime = args.ParseTimestamp(a.Until, time.Now())
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{12, 38, 0}},
		})

	a.Since = ""
	a.Until = "13/Apr/2023"
	a.StartTime = args.ParseTimestamp(a.Since, args.UnsetTime)
	a.EndTime = args.ParseTimestamp(a.Until, time.Now())
	testInputDataLoop(t, a,
		[]testData{
			{"testdata/googlebot.log", Parsed{38, 12, 0}},
		})
}

func testInputDataLoop(t *testing.T, a args.Args, d []testData) {
	for _, data := range d {
		testInputData(t, a, data.file, data.expect)
		if a.Db != "" {
			os.Remove(a.Db)
		}
	}
}

func testInputData(t *testing.T, a args.Args, file string, expect Parsed) {
	input, _ := os.ReadFile(file)
	mock.Stdin(input)

	got, err := ParseStdin(a)

	if expect.Ok != got.Ok {
		t.Fatalf("ok: expect %d, got %d", expect.Ok, got.Ok)
	}

	if expect.Skip != got.Skip {
		t.Fatalf("skip: expect %d, got %d", expect.Skip, got.Skip)
	}

	if expect.Err != got.Err {
		t.Fatalf("err: expect %d, got %d", expect.Err, got.Err)
	}

	if err != nil {
		t.Fatal(err)
	}
}
