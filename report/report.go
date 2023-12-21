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
	Runtime     RuntimeRD              `json:"runtime"`
	Parsed      ParsedRD               `json:"parsed"`
	Size        SizeRD                 `json:"size"`
	TimeArgs    TimeArgsRD             `json:"time_args"`
	Threats     ThreatsRD              `json:"threats"`
	Ads         map[string]AdsRD       `json:"ads"`
	Bots        map[string]BotsRD      `json:"bots"`
	AttackTypes AttackTypesRD          `json:"attack_types"`
	Attackers   map[string]AttackersRD `json:"attackers"`
}

type TextReport struct {
	Report

	AttackersCount []Count
	AdsCount       []Count
	BotsCount      []Count

	AttackersFiltered uint
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
	r.Attackers = make(map[string]AttackersRD)
	r.Ads = make(map[string]AdsRD)
	r.Bots = make(map[string]BotsRD)
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
	t.AttackTypes = r.AttackTypes
	t.Attackers = r.Attackers
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
	sec := int64(rt) / 1000000000
	if sec <= 0 {
		return r.Size.Byte
	}
	bps := int64(r.Size.Byte) / sec
	return uint(bps)
}

func (r *Report) LinesPerSec() uint {
	rt := r.Runtime.Total()
	if rt == 0 {
		return 0
	}
	sec := int64(rt) / 1000000000
	rlt := int64(r.Parsed.Total())
	if sec <= 0 {
		return uint(rlt)
	}
	return uint(rlt / sec)
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
	Ok    uint `json:"ok"`
	Skip  uint `json:"skip"`
	Error uint `json:"error"`
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
	Attacks Cnt `json:"attacks"`
	Ads     Cnt `json:"ads"`
	Bots    Cnt `json:"bots"`
}

type Cnt struct {
	Its  uint    `json:"count"`
	Prct float64 `json:"percent"`
}

// UTM data.
type AdsRD struct {
	Cnt uint `json:"count"`
}

// Bot data.
type BotsRD struct {
	Cnt uint `json:"count"`
}

// Attack data.
type AttackTypesRD struct {
	Unk uint `json:"gen"`
	SQL uint `json:"sql"`
	CMD uint `json:"cmd"`
	Dir uint `json:"dir"`
	XSS uint `json:"xss"`
	Scn uint `json:"scn"`
}

func (a *AttackTypesRD) Total() uint {
	return a.Unk + a.SQL + a.CMD + a.Dir + a.XSS
}

// Attacker data.
type AttackersRD struct {
	Cnt uint `json:"count"`
}
