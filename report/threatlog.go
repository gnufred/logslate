package report

import (
	"gitlab.com/gnufred/logslate/log"
	"gitlab.com/gnufred/logslate/threat"
)

type AdsEnum map[string]uint
type BotsEnum map[string]uint
type AttacksEnum map[threat.AttackType]uint
type AttackersEnum map[string]uint

type ReportTL struct {
	Ads       AdsEnum
	Bots      BotsEnum
	Attacks   AttacksEnum
	Attackers AttackersEnum
}

func NewTL() ReportTL {
	tl := ReportTL{}
	tl.init()
	return tl
}

func (rd *ReportTL) init() {
	rd.Ads = make(AdsEnum)
	rd.Bots = make(BotsEnum)
	rd.Attacks = make(AttacksEnum)
	rd.Attackers = make(AttackersEnum)
}

func (rd *ReportTL) Parse(t log.ThreatLog) {
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
