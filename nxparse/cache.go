package nxparse

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNoCachedFormat = errors.New("cached format not found")
	ErrOverwriteCache = errors.New("cannot overwrite cache")
)

// formatCache manage and store parsed `log_format` objects.
type formatCache struct {
	Formats map[string]Format
}

// Precached formats
var cache = formatCache{
	Formats: map[string]Format{
		Combined: {Vars: combinedVars},
		"":       {Vars: nilVars},
	},
}

// get return a Format from the cache store
func (c *formatCache) get(logFrmt string) (Format, error) {
	parsedFrmt, ok := c.Formats[logFrmt]
	if !ok {
		return Format{}, fmt.Errorf("%v: %w", logFrmt, ErrNoCachedFormat)
	}
	return parsedFrmt, nil
}

// set save a Format into the cache store. If you set an already existing format,
// it will be blindly overwritten.
func (c *formatCache) set(logFrmt string, parsedFrmt Format) error {
	cachedFrmt, ok := c.Formats[logFrmt]
	if ok {
		return fmt.Errorf("%v: %s: %w", cachedFrmt, logFrmt, ErrOverwriteCache)
	}

	var m sync.Mutex
	m.Lock()
	c.Formats[logFrmt] = parsedFrmt
	m.Unlock()

	return nil
}
