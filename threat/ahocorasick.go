package threat

import (
	ac "github.com/petar-dambovaliev/aho-corasick"
)

// Holds the AhoCorasick finders.
var attackAC, botAC *ac.AhoCorasick

// Options for AhoCorasick finders.
var attackOpts = ac.Opts{
	AsciiCaseInsensitive: true,
	MatchOnlyWholeWords:  false,
	MatchKind:            ac.LeftMostLongestMatch,
	DFA:                  true,
}
var botOpts = ac.Opts{
	AsciiCaseInsensitive: false,
	MatchOnlyWholeWords:  false,
	MatchKind:            ac.LeftMostLongestMatch,
	DFA:                  true,
}

// init creates "to be used" AhoCorasick builders.
func init() {
	// DataDumpPatterns() slice is a generated list. We need to
	// add it to the hardcoded attack list before we can use it.
	attackPatterns = append(attackPatterns, DataDumpPatterns()...)
	bStr := attackBuilderString(attackPatterns)
	attackAC = newAhoCorasickBuilder(bStr, attackOpts)
	botAC = newAhoCorasickBuilder(botPatterns, botOpts)
}

// newAhoCorasickBuilder create a new AC matching engine.
// Creating the builder is expesive sor we want to do this only once.
func newAhoCorasickBuilder(bStr []string, opts ac.Opts) *ac.AhoCorasick {
	builder := ac.NewAhoCorasickBuilder(opts)
	ac := builder.Build(bStr)
	return &ac
}

// attackBuilderString create the data set required by the AhoCorasick
// builder.
func attackBuilderString(pts []ptrn) []string {
	var bStr []string
	for _, p := range pts {
		bStr = append(bStr, p.Name)
	}
	return bStr
}
