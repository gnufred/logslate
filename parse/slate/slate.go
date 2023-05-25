package slate

import (
	"net"
	"strings"
	"time"

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
	Status  int64
	Sent    int64
	Referer string
	Agent   string
}

func ParseNginx(nx nginx.Entry) Log {
	ls := Log{
		Ip:      nx.RemoteAddr,
		User:    nx.RemoteUser,
		Status:  nx.Status,
		Referer: nx.Http["referer"],
		Agent:   nx.Http["user_agent"],
	}

	// assigned the parsed time
	// given we'll have TimeLocal most of the time, we test it first
	if !nx.TimeLocal.IsZero() {
		ls.Time = nx.TimeLocal
	} else if !nx.TimeIso8601.IsZero() {
		ls.Time = nx.TimeIso8601
	}

	// assign bytes set
	// here, we want BytesSent, if available, otherwise BodyBytesSent
	if nx.BytesSent > 0 {
		ls.Sent = nx.BytesSent
	} else {
		ls.Sent = nx.BodyBytesSent
	}

	//splint NGINX request into 4 different parts
	ls.Method, ls.Path, ls.Query, ls.Proto = splitNginxRequest(nx.Request)
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
