package log

import "errors"

var StringifyThreatLog = func(data any) (string, error) {
	t, ok := data.(ThreatLog)
	if !ok {
		return "", errors.New("invalid ThreatLog struct data")
	}

	s := ""

	if !t.IsBot() && !t.IsAttack() && t.IsAd() {
		s = s + "redacted -"
	} else {
		s = s + t.Ip.String() + " -"
	}

	s = s + " [" + t.Time.Format("2006-01-02 15:04:05") + "]"

	if t.IsBot() {
		s = s + " \"" + "BOT " + t.Bot + "\""
	}

	if t.IsAd() {
		s = s + " \"" + "ADV " + t.Ad + "\""
	}

	if t.IsAttack() {
		s = s + " \"" + "ATT:" + t.Attack.Name + " " + url(t.Path, t.Query) + "\""
	}

	s = s + "\n"

	return s, nil
}

func url(p, q string) string {
	if p != "" && q != "" {
		return p + "?" + q
	}
	return p + q
}
