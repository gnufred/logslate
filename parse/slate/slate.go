package slate

import (
	"strings"
	"time"

	nginx "gitlab.com/gnufred/logslate/nxparse"
)

type Log struct {
	Ip      string
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
		Time:    nx.TimeLocal,
		Status:  nx.Status,
		Sent:    nx.BytesSent,
		Referer: nx.Http["referer"],
		Agent:   nx.Http["user_agent"],
	}
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
