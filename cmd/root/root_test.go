package rootcmd

import (
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/testing/mock"
)

// Must fail given there's no data in os.Stdin
func TestNoData(t *testing.T) {
	a := args.Get()
	_, _, err := ParseStdin(a)
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
	_, _, err := ParseStdin(a)
	if err == nil {
		t.Fatalf("Did not get the expected error")
	}
}

func TestMockedData(t *testing.T) {
	a := args.Get()
	dbPath := "testdata/sql.db"

	data := []struct {
		file     string
		parsed   int
		unparsed int
	}{
		{"testdata/3ok3broken.log", 3, 3},
		{"testdata/20ok.log", 20, 0},
	}

	for _, dt := range data {
		testInputData(t, a, dt.file, dt.parsed, dt.unparsed)
	}

	a.Db = dbPath
	for _, dt := range data {
		testInputData(t, a, dt.file, dt.parsed, dt.unparsed)
		os.Remove(dbPath)
	}
}

func testInputData(t *testing.T, a args.Args, file string, parsedExpect int, unparsedExpect int) {
	input, _ := os.ReadFile(file)
	mock.Stdin(input)

	parsedGot, unparsedGot, err := ParseStdin(a)

	if parsedExpect != parsedGot {
		t.Fatalf("Expected parsed %d, got %d", parsedExpect, parsedGot)
	}

	if unparsedExpect != unparsedGot {
		t.Fatalf("Expected unparsed %d, got %d", unparsedExpect, unparsedGot)
	}

	if err != nil {
		t.Fatal(err)
	}
}
