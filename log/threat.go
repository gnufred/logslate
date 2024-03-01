package log

import (
	"encoding/json"
	"net"
	"time"

	"gitlab.com/gnufred/logslate/threat"
)

// WARNING: Changing the struct field order will
// change the threat.json output log order.
type ThreatLog struct {
	Time time.Time `json:"time"`
	Ip   net.IP    `json:"ip,omitempty"`
	threat.Detection
	Path  string `json:"path,omitempty"`
	Query string `json:"query,omitempty"`
}

func SlateToThreat(sl SlateLog) ThreatLog {
	tl := ThreatLog{
		Time: sl.Time,
	}

	if sl.IsBot() || sl.IsAttack() {
		tl.Ip = sl.Ip
	}

	if sl.IsAd() && !sl.IsBot() && !sl.IsAttack() {
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

	return tl
}

func ParseThreat(threat string) (ThreatLog, error) {
	var tl ThreatLog

	err := json.Unmarshal([]byte(threat), &tl)
	if err != nil {
		return ThreatLog{}, err
	}

	return tl, nil
}
