package printer

import (
	"errors"
	"os"
	"strconv"
	"testing"

	"gitlab.com/gnufred/logslate/args"
	"gitlab.com/gnufred/logslate/test"
)

type AB struct {
	A string `json:"a"`
	B int    `json:"b"`
}

var StringifyAB = func(data any) (string, error) {
	d, ok := data.(AB)
	if !ok {
		return "", errors.New("invalid AB struct data")
	}
	s := "a: " + d.A + ", b: " + strconv.Itoa(d.B)
	return s, nil
}

func TestOpenCloseFile(t *testing.T) {
	fp := "testdata/test.json"
	p := Printer{Path: fp}

	err := p.OpenFile()

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	p.CloseFile()
	os.Remove(fp)
}

func TestOpenFileFail(t *testing.T) {
	fp := "/path/fail/because/not/exist/test.json"
	p := Printer{Path: fp}

	err := p.OpenFile()
	exp := "no such file or directory"
	test.ErrorExpected(t, err, exp)
}

func TestError(t *testing.T) {
	Error(errors.New("test error, please ignore"))
}

func matchReadFile(t *testing.T, path, exp string) {
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error while reading %s: %s", path, err)
	}
	if string(got) != exp {
		t.Errorf("file content miss match, got %s, exp %s", got, exp)
	}
}

func TestWriteString(t *testing.T) {
	fp := "testdata/test.log"
	defer os.Remove(fp)

	p := Printer{Path: fp}
	p.OpenFile()
	defer p.CloseFile()

	exp := "hello world"
	err := p.WriteString(exp)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	matchReadFile(t, fp, exp)
}

func TestPrintScreen(t *testing.T) {
	p := Printer{}
	p.PrintScreen("test")
}

func TestPrintText(t *testing.T) {
	fp := "testdata/test.text"
	defer os.Remove(fp)

	p := Printer{Path: fp}
	p.OpenFile()
	defer p.CloseFile()

	exp := "hello world"
	err := p.PrintText(exp)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	matchReadFile(t, fp, exp)
	os.Remove(fp)
}

func TestPrintJson(t *testing.T) {
	fp := "testdata/test.json"
	defer os.Remove(fp)

	p := Printer{Path: fp}
	p.OpenFile()
	defer p.CloseFile()

	ab := AB{"hello", 999}
	exp := `{"a":"hello","b":999}`

	err := p.PrintJson(ab)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	matchReadFile(t, fp, exp)
	os.Remove(fp)
}

func TestPrintNoFormat(t *testing.T) {
	p := Printer{}
	exp := "no format"
	err := p.Print("")
	test.ErrorExpected(t, err, exp)
}

func TestPrintNoStringifier(t *testing.T) {
	p := Printer{Format: args.SCREEN}

	ab := AB{"hello", 999}
	exp := "no stringifier set"

	err := p.Print(ab)
	test.ErrorExpected(t, err, exp)
}

func TestPrintStringifyFail(t *testing.T) {
	p := Printer{
		Format:      args.SCREEN,
		Stringifier: StringifyAB,
	}

	ab := struct{ Fail bool }{true}
	exp := "struct data"

	err := p.Print(ab)
	test.ErrorExpected(t, err, exp)
}

func TestPrintFormatScreen(t *testing.T) {
	p := Printer{
		Format:      args.SCREEN,
		Stringifier: StringifyAB,
	}

	ab := AB{"hello", 999}

	err := p.Print(ab)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestPrintFormatText(t *testing.T) {
	fp := "testdata/test.text"
	defer os.Remove(fp)

	p := Printer{
		Path:        fp,
		Format:      args.TEXT,
		Stringifier: StringifyAB,
	}
	p.OpenFile()
	defer p.CloseFile()

	ab := AB{"hello", 999}
	exp, _ := StringifyAB(ab)

	err := p.Print(ab)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	matchReadFile(t, fp, exp)
}

func TestPrintFormatJson(t *testing.T) {
	fp := "testdata/test.json"
	defer os.Remove(fp)

	p := Printer{
		Path:        fp,
		Format:      args.JSON,
		Stringifier: StringifyAB,
	}
	p.OpenFile()
	defer p.CloseFile()

	ab := AB{"hello", 999}
	exp := `{"a":"hello","b":999}`

	err := p.Print(ab)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	matchReadFile(t, fp, exp)
}
