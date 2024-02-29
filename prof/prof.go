package prof

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func StartCpuProfile(path string) (*os.File, error) {
	if path == "" {
		return nil, nil
	}

	cpuf, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	err = pprof.StartCPUProfile(cpuf)
	if err != nil {
		return nil, err
	}

	return cpuf, nil
}

func StopCpuProfile(path string, cpuf *os.File) {
	if path == "" {
		return
	}
	pprof.StopCPUProfile()
	cpuf.Close()
}

func MemProfile(path string) error {
	if path == "" {
		return nil
	}

	memf, err := os.Create(path)
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
