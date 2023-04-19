package nginx

import (
    "errors"
    "log"
    "strings"
    "time"
    "strconv"

    "gitlab.com/gnufred/logslate/debug"
)

type Log struct {
    RemoteAddr string
    RemoteUser string
    TimeLocal time.Time
    Request string
    Status int64
    BytesSent int64
    HttpReferer string
    HttpUserAgent string
    GzipRatio float64
}

var DNF = errors.New("delimiter field not found")
var TimeLayout = "02/Jan/2006:15:04:05 -0700"

func Parse(log string) Log {
    p := Log{}
    
    if log == "" {
        return Log{}
    }

    p.RemoteAddr, log = parseField("-", log, [3]int{0,-1,+2})
    p.RemoteUser, log = parseField("[", log, [3]int{0,-1,+1})
    timeLocalStr, log := parseField("]", log, [3]int{0,+0,+3})
    p.TimeLocal = parseTime(timeLocalStr)
    p.Request, log = parseField("\"", log, [3]int{0,+0,+2})
    statusStr, log := parseField(" ", log, [3]int{0,0,+1})
    p.Status, _ = strconv.ParseInt(statusStr, 10, 64)
    bytesSentStr, log := parseField(" ", log, [3]int{0,0,+1})
    p.BytesSent, _ = strconv.ParseInt(bytesSentStr, 10, 64)
    p.HttpReferer, log = parseField(" ", log, [3]int{1,-1,+2})
    p.HttpUserAgent, log = parseField("\"", log, [3]int{0,+0,+1})
    p.GzipRatio = parseGzipRatio(log)

    return p
}

func parseGzipRatio(ratio string) float64 {
    if ratio == "" {
        return 0
    }
    trimed := strings.Replace(ratio, " ", "", -1)
    trimed = strings.Replace(trimed, "\"", "", -1)
    conv, _ := strconv.ParseFloat(trimed, 64)
    debug.Log(conv)
    return conv
}

func parseTime(timeLocal string) time.Time {
    parsed, err := time.Parse(TimeLayout, timeLocal)
    if err != nil {
        log.Fatalf("Fatal [unexpected]: time.Parse() returned error\n%s", err)
    }
    return parsed
}

func parseField(s, substr string, mx [3]int) (string, string) {
    cut, substr, err := cutField(s, substr, mx) 
    if err == DNF {
        log.Fatalf("Fatal [DNF]: Can't parse\n%s", substr)
    }
    if err != nil {
        log.Fatalf("Fatal [unexpected]: Can't parse\n%s\n%s", substr, err)
    }
    debug.Log(cut)
    return cut, substr
}

func cutField(s, substr string, mx [3]int) (string, string, error) {
    i := strings.Index(substr, s)
    if i == -1 {
        return "", "", DNF
    }
    return substr[mx[0]:i+mx[1]], substr[i+mx[2]:], nil
}

/* ================
Related documentation
* https://nginx.org/en/docs/http/ngx_http_log_module.html
* https://docs.nginx.com/nginx/admin-guide/monitoring/logging/#setting-up-the-access-log

Similar Golang modules
* https://github.com/fantasticmao/nginx-log-analyzer
* https://github.com/satyrius/gonx
* https://github.com/aasaam/nginx-error-log-parser
* https://sirkon.github.io/ldetool/
*/
