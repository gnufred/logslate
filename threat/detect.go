package threat

import (
	ac "github.com/petar-dambovaliev/aho-corasick"
	"gitlab.com/gnufred/logslate/debug"
)

// Holds the score that we need to flag a searches as an attack.
var Threshold = uint32(5)

// Attacks search for attack patterns in the URL query string and URL path.
func Attacks(path, query string) (uint32, byte) {

	// Important: It's faster to scan longer strings than searching twice.
	// See https://gitlab.com/gnufred/logslate/-/issues/71 for details.
	to, ty := searchAttackPts(attackAC, attackPts, path+query)
	if to >= Threshold {
		return to, ty
	}

	return 0, 0
}

// scanAttack search for attack pattern in a string
func searchAttackPts(AC *ac.AhoCorasick, pts []ptrn, haystack string) (uint32, byte) {
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

	// Debug outpt
	if total >= Threshold {
		debug.Log(debug.VVV, "high_score", highScore)
	}

	return total, highScore.Type
}
