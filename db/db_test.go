package db

import (
	"bufio"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	dns := "file:/tmp/test.db?cache=shared&mode=rw"
	file := "/tmp/test.db"
	err := dber.Create(dns, file)
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
	for scanner.Scan() {
		err = dber.InsertLog(scanner.Text())
		if err != nil {
			t.Fatalf("Fail: InsertLog() failed with error %s", err)
		}
	}
}

func TestClose(t *testing.T) {
	err := dber.Close()
	if err != nil {
		t.Fatalf("Fail: Close() failed with error %s", err)
	}
}

func TestDestroy(t *testing.T) {
	err := dber.Destroy()
	if err != nil {
		t.Fatalf("Fail: Destroy() failed with error %s", err)
	}
}
