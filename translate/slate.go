package translate

import (
	"fmt"
	"net"
	"strings"
	"time"

	"gitlab.com/gnufred/logslate/threat"

	nginx "gitlab.com/gnufred/nxparse"
)

type SlateLog struct {
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

// Translate an nxparse.Entry to the format `logslate` wants.
func NginxToSlate(nx nginx.Entry) (SlateLog, error) {
	sl := SlateLog{
		Ip:      nx.RemoteAddr,
		User:    nx.RemoteUser,
		Status:  nx.Status,
		Referer: nx.Http["referer"],
		Agent:   nx.Http["user_agent"],
	}

	// Assigned the parsed time.
	// Given we'll have TimeLocal most of the time, we test it first.
	if !nx.TimeLocal.IsZero() {
		sl.Time = nx.TimeLocal
	} else if !nx.TimeIso8601.IsZero() {
		sl.Time = nx.TimeIso8601
	}

	// Assign bytes set.
	// Here, we want BytesSent, if available, otherwise BodyBytesSent.
	if nx.BytesSent > 0 {
		sl.Sent = nx.BytesSent
	} else {
		sl.Sent = nx.BodyBytesSent
	}

	// Splint NGINX request into 4 different parts.
	rq, err := splitNginxRequest(nx.Request)
	if err != nil {
		return SlateLog{}, err
	}
	sl.Method = rq.Method
	sl.Path = rq.Path
	sl.Query = rq.Query
	sl.Proto = rq.Proto

	// Get attack detection results.
	thr, err := threat.Detect(sl.Path, sl.Query, sl.Agent)
	sl.Attack.Type = thr.Attack.Type
	sl.Attack.Name = thr.Attack.Name
	sl.Attack.Score = thr.Attack.Score
	sl.Ad = thr.Ad
	sl.Bot = thr.Bot

	// Done.
	return sl, err
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
