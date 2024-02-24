package translate

import (
	"net"
	"time"

	"gitlab.com/gnufred/logslate/threat"
)

type ThreatLog struct {
	Time   time.Time
	Ip     net.IP
	Attack threat.Attacks
	Bot    string
	Ad     string
	Path   string
	Query  string
}

func SlateToThreat(sl SlateLog) ThreatLog {

	tl := ThreatLog{
		Time:   sl.Time,
		Attack: sl.Attack.Type,
		Bot:    sl.Bot,
		Ad:     sl.Ad,
	}

	if tl.Ad != "" {
		// Anonymise the IP
		tl.Ip = net.IP{}
	} else {
		tl.Ip = sl.Ip
	}

	if tl.Attack != threat.Nil {
		tl.Path = sl.Path
		tl.Query = sl.Query
	}

	return tl
}
