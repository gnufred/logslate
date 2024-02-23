package translate

import (
	"net"
	"time"
)

type ThreatLog struct {
	Time time.Time
	Ip   net.IP
	Type string
	Info string
}

func SlateToThreat(sl SlateLog) ThreatLog {
	if sl.Attack.Name != "" {
		return attackToThreat(sl)
	} else if sl.Bot != "" {
		return botToThreat(sl)
	} else if sl.Ad != "" {
		return adToThreat(sl)
	}
	return ThreatLog{}
}

func attackToThreat(sl SlateLog) ThreatLog {
	return ThreatLog{
		Time: sl.Time,
		Ip:   sl.Ip,
		Type: sl.Attack.Name,
		Info: attackInfo(sl.Path, sl.Query),
	}
}

func attackInfo(p, q string) string {
	if p != "" && q != "" {
		return p + "?" + q
	} else {
		return p + q
	}
}

func botToThreat(sl SlateLog) ThreatLog {
	return ThreatLog{
		Time: sl.Time,
		Ip:   sl.Ip,
		Type: "bot",
		Info: sl.Bot,
	}
}

func adToThreat(sl SlateLog) ThreatLog {
	return ThreatLog{
		Time: sl.Time,
		Ip:   net.IP{},
		Type: "ad",
		Info: sl.Ad,
	}
}
