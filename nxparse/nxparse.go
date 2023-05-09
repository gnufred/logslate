// Package nxparse (NGINX parse) provides an API to parse NGINX access logs.
//
// The API respects [log_format](https://nginx.org/en/docs/http/ngx_http_log_module.html#log_format)
package nxparse

import (
	"errors"
)

// Parse parses a single line of NGINX access log and returns a structured
// representation of the entry.
//
// The format is the same (almost) string as you have in the nginx.conf `log_format`
// directive. You need to escape " characters with \ to avoid breaking GoLang strings.
// The format doesn't support new lines either. So, don't use those.
// By default, NGINX uses the `combined` log format, so you'll typically just use the
// nxparse.Combined constant which define this format.
// If you use anything else, you'll have to use your own format.
//
// The line must be the entry from the access log including the first and
// last characters of the line.
func Parse(format, line string) (Entry, error) {
	// first, check if have the parsed format in cache
	frmtParsed, err := cache.get(format)
	if err != nil && !errors.Is(err, ErrNoCachedFormat) {
		return Entry{}, err
	}

	// there's no cached object, we need to parse it
	if err != nil && errors.Is(err, ErrNoCachedFormat) {
		frmtParsed = parseFormat(format)
		err = cache.set(format, frmtParsed)
		if err != nil {
			return Entry{}, err
		}
	}

	// then finally, parse the actual line of text
	parsedLine, err := parseLine(frmtParsed, line)
	if err != nil {
		return Entry{}, err
	}
	return parsedLine, nil
}
