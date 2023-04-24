package nginx

import (
    "errors"
    "strings"
    "time"
    "strconv"
    "fmt"

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

func Parse(log string) (Log, error) {
    p := Log{}
    var err error = nil
   
    // special case: empty string
    if log == "" {
        return Log{}, err
    }

    cut := log // this copy is truncated as fields are found

    p.RemoteAddr, cut, err = parseField("-", cut, [3]int{0,-1,+2})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.RemoteUser, cut, err = parseField("[", cut, [3]int{0,-1,+1})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    timeLocalStr, cut, err := parseField("]", cut, [3]int{0,+0,+3})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.TimeLocal, err = parseTime(timeLocalStr)
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.Request, cut, err = parseField("\"", cut, [3]int{0,+0,+2})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    statusStr, cut, err := parseField(" ", cut, [3]int{0,0,+1})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.Status, err = strconv.ParseInt(statusStr, 10, 64)
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, fmt.Errorf("parseStatus [unexpected]: strconv.ParseInt error %s", err)
    }

    bytesSentStr, cut, err := parseField(" ", cut, [3]int{0,0,+1})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.BytesSent, err = strconv.ParseInt(bytesSentStr, 10, 64)
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.HttpReferer, cut, err = parseField(" ", cut, [3]int{1,-1,+2})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.HttpUserAgent, cut, err = parseField("\"", cut, [3]int{0,+0,+1})
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    p.GzipRatio, err = parseGzipRatio(cut)
    if err != nil {
        logParseFailure(log, cut, err)
        return Log{}, err
    }

    return p, err
}

func logParseFailure(log string, cut string, err error) {
    if debug.GetVerbosity() != debug.OFF {
        fmt.Println()
    }
    debug.Log(debug.V, "NxParseFail", err)
    debug.Log(debug.V, "FailingLog", log)
}

func parseGzipRatio(ratio string) (float64, error) {
    if ratio == "" {
        return 0, nil
    }
    trimed := strings.Replace(ratio, " ", "", -1)
    trimed = strings.Replace(trimed, "\"", "", -1)
    conv, err := strconv.ParseFloat(trimed, 64)
    if err != nil {
        return 0, fmt.Errorf("parseGzipRatio [unexpected]: strconv.ParseFloat error %s", err)
    }
    return conv, err
}

func parseTime(timeLocal string) (time.Time, error) {
    parsed, err := time.Parse(TimeLayout, timeLocal)
    if err != nil {
        return time.Now(), fmt.Errorf("parseTime [unexpected]: time.Parse() error %s", err)
    }
    return parsed, nil
}

func parseField(s, substr string, mx [3]int) (string, string, error) {
    cut, sub, err := cutField(s, substr, mx) 
    if err == DNF {
        return "", "", fmt.Errorf("parseField [DNF]\n(log) %s", substr)
    }
    if err != nil {
        return "", "", fmt.Errorf("parseField [unexpected]: %s %s", substr, err)
    }
    return cut, sub, nil
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
