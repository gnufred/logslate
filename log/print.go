package log

import "errors"

var StringifyThreatLog = func(data any) (string, error) {
	t, ok := data.(ThreatLog)
	if !ok {
		return "", errors.New("Fatal: error with ThreatLog struct data")
	}

	var s string
	sep := " "
	q := "\""

	s = s + q + t.Time.String() + q

	if t.Ip != nil {
		s = s + sep + t.Ip.String()
	}

	if t.IsBot() {
		s = s + sep + "bot:" + t.Bot
	}

	if t.IsAd() {
		s = s + sep + "adv:" + t.Ad
	}

	if t.IsAttack() {
		s = s + sep + "att:" + t.Attack.Name + sep + url(t.Path, t.Query)
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
