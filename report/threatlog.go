package report

import (
	"time"

	"gitlab.com/gnufred/logslate/threat"
	"gitlab.com/gnufred/logslate/translate"
)

type AdsEnum map[string]uint
type BotsEnum map[string]uint
type AttacksEnum map[threat.Attacks]uint
type AttackersEnum map[string]uint
type AttackLogEnum []AttackLogRD

type ThreatLogRD struct {
	Ads       AdsEnum
	Bots      BotsEnum
	Attacks   AttacksEnum
	Attackers AttackersEnum
	AttackLog AttackLogEnum
	First     time.Time
	Last      time.Time
}

func (rd *ThreatLogRD) parseInit() {
	rd.Ads = make(AdsEnum)
	rd.Bots = make(BotsEnum)
	rd.Attacks = make(AttacksEnum)
	rd.Attackers = make(AttackersEnum)
	rd.AttackLog = make(AttackLogEnum, 0)

	// "Zero" values for first/last entry must be, infinitely in the future/past.
	rd.First, _ = time.Parse("2006-01-02", "9999-01-02")
}

func (rd *ThreatLogRD) Parse(tl []translate.ThreatLog) {
	rd.parseInit()
	for _, t := range tl {

		if t.Time.Before(rd.First) {
			rd.First = t.Time
		}
		if t.Time.After(rd.Last) {
			rd.Last = t.Time
		}

		if t.Ad != "" {
			rd.Ads[t.Ad]++
		}
		if t.Bot != "" {
			rd.Bots[t.Bot]++
		}

		if t.Attack != threat.Nil {
			rd.Attacks[t.Attack]++
			ip := t.Ip.String()
			rd.Attackers[ip]++
			rd.parseAttackLog(t, ip)
		}
	}
}

func (rd *ThreatLogRD) parseAttackLog(t translate.ThreatLog, ip string) {
	rd.AttackLog = append(rd.AttackLog,
		AttackLogRD{
			Type:  threat.String(t.Attack),
			Time:  t.Time.Format("2006-01-02 15:04:05"),
			Ip:    ip,
			Path:  t.Path,
			Query: t.Query,
		},
	)
}
