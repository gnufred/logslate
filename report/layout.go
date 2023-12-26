package report

import (
	"strings"
)

type Opts map[string]string

type section struct {
	Name    string
	Options Opts
}

const defaultLayout = "summary,time,threats,ads:order_name,bots:order_name,attacks,attackers:order_name,attacklog"

// parseReportCustomConfig parse the custom report argument string into a usable
// data structure.
func parseLayout(layout string) []section {
	ss := []section{}
	split := strings.Split(layout, ",")
	for _, s := range split {
		ss = append(ss, parseSection(s))
	}
	return ss
}

// parseSection parses the name and options for a section string.
// The format is "name:key1_val1:key2_val2:...".
func parseSection(sec string) section {
	s := section{"", Opts{}}
	// If ":" is absent, then there's just the name.
	split := strings.Split(sec, ":")
	s.Name = split[0]
	if len(split) <= 1 {
		return s
	}

	opts := make(Opts, 1)
	for _, opt := range split[1:] {
		key, val := parseKeyValue(opt)
		opts[key] = val
	}
	s.Options = opts

	return s
}

// parseKeyValue parses the key value pair for a section option string.
// The format is "key_value".
func parseKeyValue(opt string) (string, string) {
	i := strings.Index(opt, "_")
	if i == -1 {
		return opt, ""
	}

	key := opt[:i]
	value := opt[i+1:]
	return key, value
}
