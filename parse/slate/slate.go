package slate

import (
    "strings"
    "time"

    "gitlab.com/gnufred/logslate/parse/nginx"
)

type Log struct {
    Address string
    User string
    Time time.Time
    Method string
    Path string
    Query string
    Proto string
    Status int64
    Sent int64
    Referer string
    Agent string
}

func ParseNginx(nx nginx.Log) Log {
    ls := Log{
        Address: nx.RemoteAddr,
        User: nx.RemoteUser,
        Time: nx.TimeLocal,
        Status: nx.Status,
        Sent: nx.BytesSent,
        Referer: nx.HttpReferer,
        Agent: nx.HttpUserAgent,
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
