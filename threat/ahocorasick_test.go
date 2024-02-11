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
				t.Fatalf("search result doesn't match: got %v, exp %v, for %s", matches, exp, hay)
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

// === benchmarks ===

var bHay = "action&csi=1&cso=0&id=%22%22%3E%3C%2Fscript%3E%3Cscript%3Ealert%28document.domain%29%3C%2Fscript%3E&sfi=4&si=1&so=0&st=message"

func BenchmarkAhoDfaTrue(b *testing.B) {
	testOpts := attackOpts
	testOpts.DFA = true
	aho := newAhoCorasickBuilder(
		attackBuilderString(testPatterns),
		testOpts,
	)
	// benchmark loop
	for i := 0; i < b.N; i++ {
		aho.FindAll(bHay)
	}
}

func BenchmarkAhoDfaFalse(b *testing.B) {
	testOpts := attackOpts
	testOpts.DFA = false
	aho := newAhoCorasickBuilder(
		attackBuilderString(testPatterns),
		testOpts,
	)
	// benchmark loop
	for i := 0; i < b.N; i++ {
		aho.FindAll(bHay)
	}
}

func BenchmarkAhoMatchLeft(b *testing.B) {
	testOpts := attackOpts
	testOpts.MatchKind = ac.LeftMostLongestMatch
	aho := newAhoCorasickBuilder(
		attackBuilderString(testPatterns),
		testOpts,
	)
	// benchmark loop
	for i := 0; i < b.N; i++ {
		aho.FindAll(bHay)
	}
}

func BenchmarkAhoMatchStandard(b *testing.B) {
	testOpts := attackOpts
	testOpts.MatchKind = ac.StandardMatch
	aho := newAhoCorasickBuilder(
		attackBuilderString(testPatterns),
		testOpts,
	)
	// benchmark loop
	for i := 0; i < b.N; i++ {
		aho.FindAll(bHay)
	}
}
