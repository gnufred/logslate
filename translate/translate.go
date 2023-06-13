package translate

import (
	"net"
	"strings"
	"time"

	"gitlab.com/gnufred/logslate/threat"

	nginx "gitlab.com/gnufred/nxparse"
)

type Log struct {
	Ip      net.IP
	User    string
	Time    time.Time
	Method  string
	Path    string
	Query   string
	Proto   string
	Status  int
	Sent    int
	Referer string
	Agent   string
	Score   uint32 // uint/uint64 is too big for SQLite3
	AtkId   byte
	AtkName string
}

// Nginx translate an nxparse.Entry to the format `logslate` wants.
func Nginx(nx nginx.Entry) Log {
	ls := Log{
		Ip:      nx.RemoteAddr,
		User:    nx.RemoteUser,
		Status:  nx.Status,
		Referer: nx.Http["referer"],
		Agent:   nx.Http["user_agent"],
	}

	// Assigned the parsed time.
	// Given we'll have TimeLocal most of the time, we test it first.
	if !nx.TimeLocal.IsZero() {
		ls.Time = nx.TimeLocal
	} else if !nx.TimeIso8601.IsZero() {
		ls.Time = nx.TimeIso8601
	}

	// Assign bytes set.
	// Here, we want BytesSent, if available, otherwise BodyBytesSent.
	if nx.BytesSent > 0 {
		ls.Sent = nx.BytesSent
	} else {
		ls.Sent = nx.BodyBytesSent
	}

	// Splint NGINX request into 4 different parts.
	ls.Method, ls.Path, ls.Query, ls.Proto = splitNginxRequest(nx.Request)

	// Get attack detection results.
	score, attack := threat.Attacks(ls.Path, ls.Query)
	ls.Score = score
	ls.AtkId = attack
	ls.AtkName = threat.AttackTypeString(attack)

	// Done.
	return ls
}

func splitNginxPathQuery(pq string) (string, string) {
	s := strings.Split(pq, "?")
	if len(s) == 1 {
		return s[0], ""
	}
	return s[0], s[1]
}

func splitNginxRequest(req string) (string, string, string, string) {
	s := strings.Split(req, " ")
	p, q := splitNginxPathQuery(s[1])
	return s[0], p, q, s[2]
}
