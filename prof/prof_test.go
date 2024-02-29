package prof

import (
	"os"
	"testing"
)

func TestStartCpuProfile(t *testing.T) {
	startCpuNoArgs(t)
	startCpuProf(t)
}

func startCpuNoArgs(t *testing.T) {
	file := ""
	exp := os.File{}
	got, err := StartCpuProfile(file)
	if err != nil {
		t.Error(err)
	}
	if got != nil {
		t.Errorf("exp %v, got %v", exp, got)
	}
}

func startCpuProf(t *testing.T) {
	file := "test_cpu.prof"

	prof, err := StartCpuProfile(file)
	if err != nil {
		t.Error(err)
	}

	StopCpuProfile(file, prof)

	_, err = os.Stat(file)
	if err != nil {
		t.Error("no cpu profiling file created")
	}

	err = os.Remove(file)
	if err != nil {
		t.Error(err)
	}
}

func TestMemProfile(t *testing.T) {
	startMemNoArgs(t)
	startMemProf(t)
}

func startMemNoArgs(t *testing.T) {
	file := ""
	err := MemProfile(file)
	if err != nil {
		t.Error(err)
	}
}

func startMemProf(t *testing.T) {
	file := "test_mem.prof"

	err := MemProfile(file)
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat(file)
	if err != nil {
		t.Error("no mem profiling file created")
	}

	err = os.Remove(file)
	if err != nil {
		t.Error(err)
	}
}
