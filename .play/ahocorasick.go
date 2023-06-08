package main

import (
	"fmt"

	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

// Attack types.
// Useful resource https://owasp.org/www-community/attacks/
// Here, we use bytes because we want to use the smallest possible value in our data set.
const (
	AtkSqlInj  = byte(1) // SQL injection
	AtkCmdInj  = byte(2) // Command injection
	AtkDirTrav = byte(3) // Directory traversal
	AtkXSS     = byte(4) // Cross site scripting
)

// Attack pattern structure.
type Pattern struct {
	Name  string
	Type  byte
	Score byte
	Hit   bool
}

// Attack patterns.
// This is what we're going to save in the ahocorasick search values.
//
// Given we're going to have a very large data set and loop through this once
// to populate the search values, we're using a slice given it performs better to add
// loop data.
var patterns = []Pattern{
	{Name: "%20AND", Type: AtkSqlInj, Score: 10, Hit: true},
	{Name: "%20OR", Type: AtkSqlInj, Score: 10, Hit: true},
	{Name: "%20XOR", Type: AtkSqlInj, Score: 10, Hit: true},
	{Name: ".%252F", Type: AtkDirTrav, Score: 10, Hit: false},
	{Name: "..%252F", Type: AtkDirTrav, Score: 10, Hit: false},
	{Name: "..%252f", Type: AtkDirTrav, Score: 10, Hit: false},
	{Name: "..%c0%af", Type: AtkDirTrav, Score: 5, Hit: false},
}

func main() {
	builder := ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
		AsciiCaseInsensitive: false,
		MatchOnlyWholeWords:  false,
		MatchKind:            ahocorasick.LeftMostLongestMatch,
		DFA:                  true,
	})

	var builderStr []string
	for _, p := range patterns {
		builderStr = append(builderStr, p.Name)
	}
	ac := builder.Build(builderStr)
	fmt.Println(builderStr)

	line := "/..%252F.%252F...%252F....%252F..%c0%af..%c0%af..%252F..%252f..%252f..%252F..%252F..%252Fetc%252Fpasswd"
	matches := ac.FindAll(line)

	for _, m := range matches {
		fmt.Println(
			m.Pattern(),
			m.Start(),
			m.End(),
			line[m.Start():m.End()],
			patterns[m.Pattern()],
		)
	}
}
