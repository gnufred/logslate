package prof

import (
	"os"
	"runtime"
	"runtime/pprof"

	"gitlab.com/gnufred/logslate/args"
)

func StartCpuProfile(a args.Args) (*os.File, error) {
	if a.CpuProf == "" {
		return nil, nil
	}

	cpuf, err := os.Create(a.CpuProf)
	if err != nil {
		return nil, err
	}
	err = pprof.StartCPUProfile(cpuf)
	if err != nil {
		return nil, err
	}

	return cpuf, nil
}

func StopCpuProfile(a args.Args, cpuf *os.File) {
	if a.CpuProf == "" {
		return
	}
	pprof.StopCPUProfile()
	cpuf.Close()
}

func MemProfile(a args.Args) error {
	if a.MemProf == "" {
		return nil
	}

	memf, err := os.Create(a.MemProf)
	if err != nil {
		return err
	}
	defer memf.Close()
	runtime.GC()
	err = pprof.WriteHeapProfile(memf)
	if err != nil {
		return err
	}
	return nil
}
