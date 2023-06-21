package prof

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"gitlab.com/gnufred/logslate/args"
)

func StartCpuProfile(a args.Args) *os.File {
	if a.CpuProf == "" {
		return nil
	}

	cpuf, err := os.Create(a.CpuProf)
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(cpuf); err != nil {
		log.Fatal(err)
	}

	return cpuf
}

func StopCpuProfile(a args.Args, cpuf *os.File) {
	if a.CpuProf == "" {
		return
	}
	pprof.StopCPUProfile()
	cpuf.Close()
}

func MemProfile(a args.Args) {
	if a.MemProf == "" {
		return
	}

	memf, err := os.Create(a.MemProf)
	if err != nil {
		log.Fatal(err)
	}
	defer memf.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal(err)
	}
}
