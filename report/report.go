package report

import (
	"time"
)

// Global data set.
type Report struct {
	path   string
	format string
	layout []section

	// Report aggregated data.
	Run   Runtimes         `json:"runtimes"`
	Line  Lines            `json:"parsed"`
	Size  Sizes            `json:"size"`
	Time  Times            `json:"time_args"`
	Thrt  Thrts            `json:"threats"`
	AtkTy AtkTypes         `json:"attack_types"`
	Atkr  map[string]Atkrs `json:"attackers"`
	Bot   map[string]Bots  `json:"bots"`
}

type TextReport struct {
	Report
	// Order by count data structures.
	AtkrCnt []Count
	BotCnt  []Count
}

type Count struct {
	Num  uint
	Item string
}

func New(path, format, layout string) *Report {
	r := Report{}
	r.path = path
	r.format = format
	if layout == "" {
		layout = defaultLayout
	}
	r.layout = parseLayout(layout)
	r.Atkr = make(map[string]Atkrs)
	r.Bot = make(map[string]Bots)
	return &r
}

func NewText(r *Report) *TextReport {
	t := TextReport{}
	t.path = r.path
	t.format = r.format
	t.layout = r.layout
	t.Run = r.Run
	t.Line = r.Line
	t.Size = r.Size
	t.Time = r.Time
	t.Thrt = r.Thrt
	t.AtkTy = r.AtkTy
	t.Atkr = r.Atkr
	t.Bot = r.Bot

	return &t
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

// Threat counts.
type Thrts struct {
	Atks Cnt `json:"attacks"`
	Bots Cnt `json:"bots"`
}

type Cnt struct {
	Its  uint    `json:"count"`
	Prct float64 `json:"percent"`
}

// Attack data.
type AtkTypes struct {
	Unk uint `json:"gen"`
	SQL uint `json:"sql"`
	CMD uint `json:"cmd"`
	Dir uint `json:"dir"`
	XSS uint `json:"xss"`
}

func (a *AtkTypes) Total() uint {
	return a.Unk + a.SQL + a.CMD + a.Dir + a.XSS
}

// Attackers data.
type Atkrs struct {
	Cnt uint `json:"count"`
}

// Bots data.
type Bots struct {
	Cnt uint `json:"count"`
}
