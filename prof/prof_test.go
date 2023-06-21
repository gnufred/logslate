package prof

import (
	"os"
	"testing"

	"gitlab.com/gnufred/logslate/args"
)

func TestStartCpuProfile(t *testing.T) {
	startCpuNoArgs(t)
	startCpuProf(t)
}

func startCpuNoArgs(t *testing.T) {
	a := args.Get()
	exp := os.File{}
	got, err := StartCpuProfile(a)
	if err != nil {
		t.Error(err)
	}
	if got != nil {
		t.Errorf("exp %v, got %v", exp, got)
	}
}

func startCpuProf(t *testing.T) {
	a := args.Get()
	file := "test_cpu.prof"
	a.CpuProf = file

	prof, err := StartCpuProfile(a)
	if err != nil {
		t.Error(err)
	}

	StopCpuProfile(a, prof)

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
	a := args.Get()
	err := MemProfile(a)
	if err != nil {
		t.Error(err)
	}
}

func startMemProf(t *testing.T) {
	a := args.Get()
	file := "test_mem.prof"
	a.MemProf = file

	err := MemProfile(a)
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
