package db

import (
	"bufio"
	"os"
	"testing"

    "gitlab.com/gnufred/logslate/parse/nginx"
    "gitlab.com/gnufred/logslate/parse/slate"
)

func TestCreate(t *testing.T) {
	file := "/tmp/test.db"
	db := Singleton()
	err := db.Create(file)
	if err != nil {
		t.Fatalf("Fail: Create() failed with error %s", err)
	}
}

func TestInsertLog(t *testing.T) {
	file, err := os.Open("testdata/ahrefs.log")
	if err != nil {
		t.Fatalf("Fail: Opening test data file failed with error %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	db := Singleton()
	for scanner.Scan() {
        nx := nginx.Parse(scanner.Text())
        sl := slate.ParseNginx(nx)
		err = db.InsertLog(sl)
		if err != nil {
			t.Fatalf("Fail: InsertLog() failed with error %s", err)
		}
	}
}

func TestClose(t *testing.T) {
	db := Singleton()
	err := db.Close()
	if err != nil {
		t.Fatalf("Fail: Close() failed with error %s", err)
	}
}

func TestDestroy(t *testing.T) {
	db := Singleton()
	err := db.Destroy()
	if err != nil {
		t.Fatalf("Fail: Destroy() failed with error %s", err)
	}
}
