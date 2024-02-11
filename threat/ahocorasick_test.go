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

func benchAho(b *testing.B, opts ac.Opts, pts []ptrn, s string) {
	aho := newAhoCorasickBuilder(attackBuilderString(pts), opts)
	for i := 0; i < b.N; i++ {
		aho.FindAll(s)
	}
}

// Benchmark DFA true/false
// Result: DFA true is faster
func BenchmarkAhoDfaTrue(b *testing.B) {
	opts := attackOpts
	opts.DFA = true
	benchAho(b, opts, attackPatterns, bHay)
}

func BenchmarkAhoDfaFalse(b *testing.B) {
	opts := attackOpts
	opts.DFA = false
	benchAho(b, opts, attackPatterns, bHay)
}

// Benchmark match left/standard
// Result: LeftMostLongestMatch is faster
func BenchmarkAhoMatchLeft(b *testing.B) {
	opts := attackOpts
	opts.MatchKind = ac.LeftMostLongestMatch
	benchAho(b, opts, attackPatterns, bHay)
}

func BenchmarkAhoMatchStandard(b *testing.B) {
	opts := attackOpts
	opts.MatchKind = ac.StandardMatch
	benchAho(b, opts, attackPatterns, bHay)
}

// https://gitlab.com/gnufred/logslate/-/issues/176
//
// See if it's worth changing how we detect data dump `scan` attacks.
// Run with
// go test "./threat/..." --test.bench BenchmarkDumpScan
//
// Result
//
// cpu: 12th Gen Intel(R) Core(TM) i7-12700H
// BenchmarkDumpScanLeft-20        	 2670974	       438.9 ns/op
// BenchmarkDumpScanStandard-20    	 2517352	       471.5 ns/op
//
// Conclusion
//
// LeftMost search is so much faster than standard search, than it's
// still faster to have a lot of patterns, then to try and combine
// matches togheter, in a small list.

func BenchmarkDumpScanLeft(b *testing.B) {
	file := []string{"api", "backup", "bak", "bin", "db", "db_backup", "dbdump", "dump", "html", "html", "old", "public", "release", "root", "site", "test", "www", "admin", "archive", "dump", "another", "file", "type", "because"}
	ext := []string{"sql", "7z", "bz2", "gz", "lz", "rar", "tar", "zip", "targz", "gzzz", "zx", "zx2", "zx3"}

	dump := make([]ptrn, 0, len(file)*len(ext))
	for _, f := range file {
		for _, e := range ext {
			dump = append(dump, ptrn{
				Name:  f + "." + e,
				Type:  Scn,
				Score: 5,
			})
		}
	}

	opts := attackOpts
	benchAho(b, opts, dump, bHay)
}

func BenchmarkDumpScanStandard(b *testing.B) {
	file := []string{"api", "backup", "bak", "bin", "db", "db_backup", "dbdump", "dump", "html", "html", "old", "public", "release", "root", "site", "test", "www", "admin", "archive", "dump", "another", "file", "type", "because"}
	ext := []string{"sql", "7z", "bz2", "gz", "lz", "rar", "tar", "zip", "targz", "gzzz", "zx", "zx2", "zx3"}

	dump := make([]ptrn, 0, len(file)+len(ext))
	for _, f := range file {
		dump = append(dump, ptrn{
			Name:  f,
			Type:  Scn,
			Score: 3,
		})
	}
	for _, f := range ext {
		dump = append(dump, ptrn{
			Name:  f,
			Type:  Scn,
			Score: 3,
		})
	}

	opts := attackOpts
	opts.MatchKind = ac.StandardMatch
	benchAho(b, opts, dump, bHay)
}
