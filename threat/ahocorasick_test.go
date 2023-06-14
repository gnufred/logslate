package threat

import (
	"fmt"
	"testing"

	ac "github.com/petar-dambovaliev/aho-corasick"
)

var testAC *ac.AhoCorasick

func init() {
	bStr := attackBuilderString(testPatterns)
	testAC = newAhoCorasickBuilder(bStr, attackOpts)
}

func TestNewAhoCorasickBuilder(t *testing.T) {
	got := fmt.Sprintf("%T", testAC)
	exp := "*aho_corasick.AhoCorasick"
	if got != exp {
		t.Errorf("type doesn't match, got %s, exp %s", got, exp)
	}
}

func TestACMatches(t *testing.T) {
	type result struct{ P, L, E int }

	testMatch := func(t *testing.T, hay string, exp []result) {
		matches := testAC.FindAll(hay)
		for k, m := range matches {
			l := m.End() - m.Start()
			if exp[k].P != m.Pattern() || exp[k].L != l || exp[k].E != m.End() {
				t.Fatalf("search result doesn't match, got %v, exp %v", matches, exp)
			}
		}
	}

	testMatch(t, "x", []result{{0, 1, 1}})
	testMatch(t, "xy", []result{{1, 2, 2}})
	testMatch(t, "xyz", []result{{2, 3, 3}})
	testMatch(t, "axyz", []result{{2, 3, 4}})
	testMatch(t, "axyzbx", []result{{2, 3, 4}, {0, 1, 6}})
	testMatch(t, "axyzbxy", []result{{2, 3, 4}, {1, 2, 7}})
}
