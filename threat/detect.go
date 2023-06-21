package threat

import (
	"fmt"
	"strings"

	ac "github.com/petar-dambovaliev/aho-corasick"
	"gitlab.com/gnufred/logslate/debug"
)

// Holds the score that we need to flag a searches as an attack.
var Threshold = uint32(5)

type Threat struct {
	Attack Attack
	Bot    string
}

type Attack struct {
	Id    byte
	Score uint32
	Name  string
}

// Detect threats and return information about them.
func Detect(path, query, agent string) (Threat, error) {
	attack, score := attackScan(path, query)
	bot, err := botScan(agent)

	return Threat{
		Attack: Attack{
			Id:    attack,
			Score: score,
			Name:  String(attack),
		},
		Bot: bot,
	}, err
}

// attackScan searches for attack patterns in the URL query string and URL path.
func attackScan(path, query string) (byte, uint32) {
	// Important: It's faster to scan longer strings than searching twice.
	// See https://gitlab.com/gnufred/logslate/-/issues/71 for details.
	attack, total := attackAhoSearch(attackAC, attackPatterns, path+query)
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

	// The default bot format is "botname/v9.9.9" and some patterns need to match
	// this character. '/' is byte(47).
	if strings.HasSuffix(bot, "/") {
		return bot[:len(bot)-1], err
	}

	return bot, err
}

// botAhoSearch searches for bot patterns.
func botAhoSearch(AC *ac.AhoCorasick, haystack string) (string, error) {
	matches := AC.FindAll(haystack)

	if len(matches) == 0 {
		return "", nil
	}

	m := matches[0]
	// As we can see in the benchmarks, it's two time faster to get the bot name
	// from the substring than from the pattern []string.
	bot := haystack[m.Start():m.End()]

	// If the user agent string is "real", the shouldn't be more than one match.
	// That being said, the AhoSearch can return more than one matches, so we
	// kinda need to handle that case.
	// The other matches should be safe to ignore in 99.99% of the cases.
	// I can only think of problem with "fake" user agents.
	if len(matches) > 1 {
		var bots string
		for _, m := range matches {
			bots += haystack[m.Start():m.End()] + " "
		}
		return bot, fmt.Errorf("found more than one bot: %s", bots)
	}

	return bot, nil
}
