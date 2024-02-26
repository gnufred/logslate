package log

import (
	"net"
	"time"

	"gitlab.com/gnufred/logslate/threat"
)

type ThreatLog struct {
	threat.Detection
	Time  time.Time `json:"time"`
	Ip    net.IP    `json:"ip,omitempty"`
	Path  string    `json:"path,omitempty"`
	Query string    `json:"query,omitempty"`
}

func SlateToThreat(sl SlateLog) ThreatLog {
	tl := ThreatLog{
		Time: sl.Time,
	}

	if sl.IsAd() && !sl.IsAttack() {
		tl.Ad = sl.Ad
	}

	if sl.IsBot() {
		tl.Bot = sl.Bot
	}

	if sl.IsAttack() {
		tl.Attack = sl.Attack
		tl.Path = sl.Path
		tl.Query = sl.Query
	}

	if !anonymizeIp(sl) {
		tl.Ip = sl.Ip
	}

	return tl
}

func anonymizeIp(sl SlateLog) bool {
	return sl.IsAd()
}
