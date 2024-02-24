package report

import (
	"gitlab.com/gnufred/logslate/log"
	"gitlab.com/gnufred/logslate/threat"
)

type AdsEnum map[string]uint
type BotsEnum map[string]uint
type AttacksEnum map[threat.AttackType]uint
type AttackersEnum map[string]uint

type ThreatLogRD struct {
	Ads       AdsEnum
	Bots      BotsEnum
	Attacks   AttacksEnum
	Attackers AttackersEnum
}

func NewThreatLog() ThreatLogRD {
	tl := ThreatLogRD{}
	tl.init()
	return tl
}

func (rd *ThreatLogRD) init() {
	rd.Ads = make(AdsEnum)
	rd.Bots = make(BotsEnum)
	rd.Attacks = make(AttacksEnum)
	rd.Attackers = make(AttackersEnum)
}

func (rd *ThreatLogRD) Parse(t log.ThreatLog) {
	if t.IsAd() {
		rd.Ads[t.Ad]++
	}
	if t.IsBot() {
		rd.Bots[t.Bot]++
	}

	if t.IsAttack() {
		rd.Attacks[t.Attack.Type]++
		ip := t.Ip.String()
		rd.Attackers[ip]++
	}
}
