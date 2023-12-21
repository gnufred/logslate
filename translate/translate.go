package translate

import (
	"fmt"
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
	Attack  threat.Attack
	Ad      string
	Bot     string
}

type request struct {
	Method string
	Path   string
	Query  string
	Proto  string
}

// Nginx translate an nxparse.Entry to the format `logslate` wants.
func Nginx(nx nginx.Entry) (Log, error) {
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
	rq, err := splitNginxRequest(nx.Request)
	if err != nil {
		return Log{}, err
	}
	ls.Method = rq.Method
	ls.Path = rq.Path
	ls.Query = rq.Query
	ls.Proto = rq.Proto

	// Get attack detection results.
	thr, err := threat.Detect(ls.Path, ls.Query, ls.Agent)
	ls.Attack.Type = thr.Attack.Type
	ls.Attack.Name = thr.Attack.Name
	ls.Attack.Score = thr.Attack.Score
	ls.Ad = thr.Ad
	ls.Bot = thr.Bot

	// Done.
	return ls, err
}

func splitNginxPathQuery(pq string) (string, string) {
	s := strings.Split(pq, "?")
	if len(s) == 1 {
		return s[0], ""
	}
	return s[0], s[1]
}

// splitNginxRequest parse an NGINX log_format $request string.
// The string is composed of for elements from the said request:
// method, path, query, and protocol.
func splitNginxRequest(req string) (request, error) {
	s := strings.Split(req, " ")
	if len(s) < 3 {
		return request{}, fmt.Errorf("malformed NGINX request string: %s", req)
	}
	path, query := splitNginxPathQuery(s[1])
	return request{s[0], path, query, s[2]}, nil
}
