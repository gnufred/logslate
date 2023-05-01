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

func TestMockedData(t *testing.T) {
	a := args.Get()
	dbPath := "testdata/sql.db"

	data := []struct {
		file   string
		expect Parsed
	}{
		{"testdata/3ok3broken.log", Parsed{3, 0, 3}},
		{"testdata/20ok.log", Parsed{20, 0, 0}},
	}

	for _, dt := range data {
		testInputData(t, a, dt.file, dt.expect)
	}

	a.Db = dbPath
	for _, dt := range data {
		testInputData(t, a, dt.file, dt.expect)
		os.Remove(dbPath)
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
