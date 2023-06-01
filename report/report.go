package report

import (
	"time"
)

// Global data set.
type Report struct {
	Path string `json:"path"`
	Frmt string `json:"format"`

	// Report aggregated data.
	Run  Runtimes `json:"runtimes"`
	Line Lines    `json:"parsed"`
	Size Sizes    `json:"size"`
	Time Times    `json:"time_args"`
	Db   Dbs      `json:"database"`
}

func New() *Report {
	return &Report{}
}

// Performance data.
func (r *Report) BytesPerSec() uint {
	rt := r.Run.Total()
	if rt == 0 {
		return 0
	}
	sec := int64(rt) / 1000000000
	if sec <= 0 {
		return r.Size.Byte
	}
	bps := int64(r.Size.Byte) / sec
	return uint(bps)
}

func (r *Report) LinesPerSec() uint {
	rt := r.Run.Total()
	if rt == 0 {
		return 0
	}
	sec := int64(rt) / 1000000000
	rlt := int64(r.Line.Total())
	if sec <= 0 {
		return uint(rlt)
	}
	return uint(rlt / sec)
}

// Runtime data.
type Runtimes struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (r *Runtimes) Total() time.Duration {
	if r.Start.Before(r.End) {
		return r.End.Sub(r.Start)
	}
	return 0
}

// Line data.
type Lines struct {
	Ok    uint `json:"ok"`
	Skip  uint `json:"skip"`
	Error uint `json:"error"`
}

func (l *Lines) Total() uint {
	return l.Ok + l.Skip + l.Error
}

// Size data.
type Sizes struct {
	Byte uint `json:"byte"`
}

// Time options data.
type Times struct {
	Since string    `json:"since"`
	Until string    `json:"until"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// Database data.
type Dbs struct {
	Path string `json:"path"`
}
