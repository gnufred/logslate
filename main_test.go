package main

import (
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

func TestInit(t *testing.T) {
	Init()
}

func TestBanner(t *testing.T) {
	Banner()
}

func TestRun(t *testing.T) {
	testRunOk(t)
	testRunTooLong(t)
}

func testRunOk(t *testing.T) {
	file := "root/testdata/20ok.log"
	input, _ := os.ReadFile(file)
	test.MockStdin(input)
	Run()
}

func testRunTooLong(t *testing.T) {
	file := "root/testdata/20toolong.log"
	input, _ := os.ReadFile(file)
	test.MockStdin(input)
	Run()
}
