package report

import (
	"time"
)

// Global data set.
type Report struct {
	// Report specifications.
	path   string
	format string
	layout []section

	// Report Data (RD).
	Runtime   RuntimeRD  `json:"runtime"`
	Parsed    ParsedRD   `json:"parsed"`
	Size      SizeRD     `json:"size"`
	TimeArgs  TimeArgsRD `json:"time_args"`
	Threats   ThreatsRD  `json:"threats"`
	ThreatLog ThreatLogRD
	Ads       map[string]ListRD `json:"ads"`
	Bots      map[string]ListRD `json:"bots"`
	Attacks   AttacksRD         `json:"attacks"`
	Attackers map[string]ListRD `json:"attackers"`
	AttackLog []AttackLogRD     `json:"attacks"`
}

type AttackLogRD struct {
	Type  string `json:"type"`
	Time  string `json:"time"`
	Ip    string `json:"ip"`
	Path  string `json:"path"`
	Query string `json:"query"`
}

type TextReport struct {
	Report

	AttackersCount []Count
	AdsCount       []Count
	BotsCount      []Count

	AttackersFiltered uint
	AdsFiltered       uint
	BotsFiltered      uint
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
	r.Ads = make(map[string]ListRD)
	r.Bots = make(map[string]ListRD)
	r.Attackers = make(map[string]ListRD)
	return &r
}

func NewText(r *Report) *TextReport {
	t := TextReport{}
	t.path = r.path
	t.format = r.format
	t.layout = r.layout
	t.Runtime = r.Runtime
	t.Parsed = r.Parsed
	t.Size = r.Size
	t.TimeArgs = r.TimeArgs
	t.Threats = r.Threats
	t.Attacks = r.Attacks
	t.Attackers = r.Attackers
	t.AttackLog = r.AttackLog
	t.Ads = r.Ads
	t.Bots = r.Bots

	return &t
}

// Performance data.
func (r *Report) BytesPerSec() uint {
	rt := r.Runtime.Total()
	if rt == 0 {
		return 0
	}
	msec := int64(rt) / 1000
	if msec <= 0 {
		return 0
	}
	bps := int64(r.Size.Byte) / msec * 1000000
	return uint(bps)
}

func (r *Report) LinesPerSec() uint {
	rt := r.Runtime.Total()
	if rt == 0 {
		return 0
	}
	msec := int64(rt) / 1000000
	rlt := int64(r.Parsed.Total())
	if msec <= 0 {
		return 0
	}
	return uint(rlt / msec * 1000)
}

// Runtime data.
type RuntimeRD struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (r *RuntimeRD) Total() time.Duration {
	if r.Start.Before(r.End) {
		return r.End.Sub(r.Start)
	}
	return 0
}

// Parsed data.
type ParsedRD struct {
	Ok    uint      `json:"ok"`
	Skip  uint      `json:"skip"`
	Error uint      `json:"error"`
	First time.Time `json:"first"`
	Last  time.Time `json:"last"`
}

func (l *ParsedRD) Total() uint {
	return l.Ok + l.Skip + l.Error
}

// Size data.
type SizeRD struct {
	Byte uint `json:"byte"`
}

// Time arguments data.
type TimeArgsRD struct {
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
type ThreatsRD struct {
	Attacks PrctRD `json:"attacks"`
	Ads     PrctRD `json:"ads"`
	Bots    PrctRD `json:"bots"`
}

type PrctRD struct {
	Its  uint    `json:"count"`
	Prct float64 `json:"percent"`
}

// Attack data.
type AttacksRD struct {
	Unk uint `json:"gen"`
	SQL uint `json:"sql"`
	CMD uint `json:"cmd"`
	Dir uint `json:"dir"`
	XSS uint `json:"xss"`
	Scn uint `json:"scn"`
}

func (a *AttacksRD) Total() uint {
	return a.Unk + a.SQL + a.CMD + a.Dir + a.XSS
}

// Generic list with count.
type ListRD struct {
	Cnt uint `json:"count"`
}
