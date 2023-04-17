package db

import (
	"testing"
)

func TestCreate(t *testing.T) {
	dns := "file:/tmp/test.db?cache=shared&mode=rw"
	file := "/tmp/test.db"
	dber.Create(dns, file)
}

func TestDestroy(t *testing.T) {
	dber.Destroy()
}
