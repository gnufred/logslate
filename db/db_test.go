package db

import (
	"testing"
)

func TestNew(t *testing.T) {
	dns := "file:test.db?cache=shared&mode=rw"
	file := "test.db"
	dber.New(dns, file)
}
