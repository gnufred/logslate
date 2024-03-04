package args

import (
	"time"
)

type ArgsV2 struct {
	Help    bool
	Command SubCmd

	Debug  debugArgs
	Parse  parseArgs
	Time   timeArgs // Use Get/Set functions.
	Threat threatArgs
	Report reportArgs
}

type debugArgs struct {
	Verbosity int
	CpuProf   string
	MemProf   string
}

type parseArgs struct {
	Db       string
	LogFrmt  string
	HideProg bool
}

type timeArgs struct {
	// Set with .SetSinceUntil()
	since     string    // Geter .Since()
	until     string    // Geter .Until()
	startTime time.Time // Geter .StartTime()
	endTime   time.Time // Geter .EndTime()
}

type threatArgs struct {
	Path   string
	Format LogFormat
}

type reportArgs struct {
	Path   string
	Format LogFormat
	Layout string
}

var argInst *ArgsV2

func Singleton() *ArgsV2 {
	if argInst == nil {
		panic("FATAL: argument not instanciated")
	}
	return argInst
}
