package detect

import ac "github.com/petar-dambovaliev/aho-corasick"

var attackAC *ac.AhoCorasick

// init creates "to be used" AhoCorasick builders
func init() {
	attackAC = newAhoCorasickBuilder(attackPts)
}

// newAhoCorasickBuilder create a new AC matching engine.
// Creating the builder is expesive sor we want to do this only once.
func newAhoCorasickBuilder(ptrns []ptrn) *ac.AhoCorasick {
	builder := ac.NewAhoCorasickBuilder(ac.Opts{
		AsciiCaseInsensitive: true,
		MatchOnlyWholeWords:  false,
		MatchKind:            ac.LeftMostLongestMatch,
		DFA:                  true,
	})

	var builderStr []string
	for _, p := range ptrns {
		builderStr = append(builderStr, p.Name)
	}
	ac := builder.Build(builderStr)

	return &ac
}
