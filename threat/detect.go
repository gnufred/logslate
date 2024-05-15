package threat

import (
	"regexp"

	ac "github.com/petar-dambovaliev/aho-corasick"
	"gitlab.com/gnufred/logslate/debug"
)

// Advertising compiled regexp
var regexpAdsGeneric *regexp.Regexp
var regexpAdsGoogle *regexp.Regexp

// Holds the score that we need to flag a searches as an attack.
var Threshold = uint32(5)

// WARNING: Changing the struct field order will
// change the threat.json output log order.
type Detection struct {
	Bot    string  `json:"bot,omitempty"`
	Ad     string  `json:"ad,omitempty"`
	Attack *Attack `json:"attack,omitempty"`
}

func (d *Detection) IsAd() bool     { return d.Ad != "" }
func (d *Detection) IsBot() bool    { return d.Bot != "" }
func (d *Detection) IsAttack() bool { return d.Attack != nil && d.Attack.Type != NIL }
func (d *Detection) IsThreat() bool { return d.IsBot() || d.IsAttack() || d.IsAd() }

func (d *Detection) AttackName() string {
	if d.IsAttack() {
		return d.Attack.Name
	}
	return ""
}

type Attack struct {
	Score uint32     `json:"-"`
	Type  AttackType `json:"type"`
	Name  string     `json:"name"`
}

func init() {
	// MustCompile() cost about 90% of the execution time.
	// We need to run this outside of the detection loop.
	regexpAdsGeneric = regexp.MustCompile(`utm_(?:source|campaign)=(\w*)`)
	regexpAdsGoogle = regexp.MustCompile(`(?i)(gclid|dclid)=`)
}

// Detect threats and return information about them.
func Detect(path, query, agent string) (Detection, error) {
	var attackPtr *Attack
	attack, score := attackScan(path, query, agent)
	if attack != NIL {
		attackPtr = &Attack{
			Type:  attack,
			Score: score,
			Name:  String(attack),
		}
	}

	// If it's an attack, obviously, it not an advertising campaign.
	ad := ""
	if attack == NIL {
		ad = adScan(query)
	}

	// We always want to know if it's a bot.
	// Some attacks comes from legit bots, like security scanners.
	// Advertisement, are also sometimes from bots.
	bot, err := botScan(agent)

	return Detection{
		Ad:     ad,
		Bot:    bot,
		Attack: attackPtr,
	}, err
}

// attackScan searches for attack patterns in the URL query string and URL path.
func attackScan(path, query, agent string) (AttackType, uint32) {
	// Important: It's faster to scan longer strings than searching twice.
	// See https://gitlab.com/gnufred/logslate/-/issues/71 for details.
	attack, total := attackAhoSearch(attackAC, attackPatterns, path+" "+query+" "+agent)
	if total >= Threshold {
		debug.Log(debug.VVV, "high_score", attack)
		return attack.Type, total
	}

	return 0, 0
}

// attackAhoSearch searches for attack patterns.
func attackAhoSearch(AC *ac.AhoCorasick, pts []ptrn, haystack string) (ptrn, uint32) {
	var total uint32
	highScore := ptrn{}

	for _, m := range AC.FindAll(haystack) {
		// Matched attack pattern.
		found := pts[m.Pattern()]

		// Save the highest scoring match.
		if found.Score > highScore.Score {
			highScore = found
		}

		// Count total score.
		total += uint32(found.Score)
	}

	return highScore, total
}

// botScan for bot patterns in the user agent string.
func botScan(agent string) (string, error) {
	bot, err := botAhoSearch(botAC, agent)
	return bot, err
}

// botAhoSearch searches for bot patterns.
func botAhoSearch(AC *ac.AhoCorasick, haystack string) (string, error) {
	matches := AC.FindAll(haystack)

	if len(matches) == 0 {
		return "", nil
	}
	// Bot names are usually duplicated in the user agent string.
	// Ex.: (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)
	// It's perfectly normal to find more than one match.
	// Only the first match is used.
	m := matches[0]

	// As we can see in the benchmarks, it's two time faster to get the bot name
	// from the substring than from the pattern []string.
	bot := haystack[m.Start():m.End()]

	return bot, nil
}

// adScan search a query string for "utm_source" query param and returns the value.
func adScan(query string) string {
	var matches []string

	// Search for generic advertisement.
	matches = regexpAdsGeneric.FindStringSubmatch(query)
	if len(matches) > 0 {
		return matches[1]
	}

	// Search for Google advertisement.
	// https://support.google.com/google-ads/answer/9744275?hl=en
	// URLs for such campaigns will include `gclid=`.
	// Legacy param is `dclid=` which is still in use.
	matches = regexpAdsGoogle.FindStringSubmatch(query)
	if len(matches) > 0 {
		return matches[1]
	}

	return ""
}
