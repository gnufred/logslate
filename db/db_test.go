package db

import (
	"bufio"
	"os"
	"testing"

    "gitlab.com/gnufred/logslate/parse/nginx"
    "gitlab.com/gnufred/logslate/parse/slate"
)

func TestCreate(t *testing.T) {
	file := "testdata/test.db"
	db := Singleton()
	err := db.Create(file)
	if err != nil {
        t.Fatalf("Create() failed: %s", err)
	}
}

func TestInsertLog(t *testing.T) {
	file, err := os.Open("testdata/ahrefs.log")
	if err != nil {
        t.Fatalf("Opening test data file failed: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	db := Singleton()
    db.Begin()
    defer db.Commit()

	for scanner.Scan() {
        nx, _ := nginx.Parse(scanner.Text())
        if err != nil {
            t.Fatalf("nginx.Parsed() failed: %s", err)
        }
        sl := slate.ParseNginx(nx)
		err = db.InsertLog(sl)
		if err != nil {
            t.Fatalf("InsertLog() failed: %s", err)
		}
	}
}

func TestClose(t *testing.T) {
	db := Singleton()
	err := db.Close()
	if err != nil {
        t.Fatalf("Close() failed: %s", err)
	}
}

func TestDestroy(t *testing.T) {
	db := Singleton()
	err := db.Destroy()
	if err != nil {
        t.Fatalf("Destroy() failed: %s", err)
	}
}
