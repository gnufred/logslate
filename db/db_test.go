package db

import (
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
