package nxparse

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	TimeLayoutIso8601 = time.RFC3339
	TimeLayoutLocal   = "02/Jan/2006:15:04:05 -0700"
)

// parse errors
var (
	ErrNilVarName   = errors.New("format variable nil after first position")
	ErrNilSeparator = errors.New("format seperator nil before last position")
	ErrParseFailed  = errors.New("parsing log line failed")
	ErrSepNotFound  = errors.New("separator not found")
	ErrVarNotExist  = errors.New("log variable does not exist")
	ErrConverFail   = errors.New("value conversion failed")
	ErrIpConvFail   = errors.New("IP address conversion failed")
)

// An Entry represent a line from the NGINX access log.
// Typically, you'll use Parse() to populate this.
type Entry struct {
	BodyBytesSent      int64
	BytesSent          int64
	Connection         string
	ConnectionRequests int64
	Http               map[string]string
	Msec               int64
	Pipe               bool
	RemoteAddr         net.IP
	RemoteUser         string
	Request            string
	RequestLength      int64
	RequestTime        int64
	Status             int64
	TimeIso8601        time.Time
	TimeLocal          time.Time
}

// parseLine does the actual parsing of the log string into structured data.
func parseLine(frmt Format, line string) (Entry, error) {
	// prepare scan
	frmtLen := len(frmt.Vars)
	ent := Entry{}

	// scan each formatVars
	for i := 0; i < frmtLen; i++ {
		// prepare variables
		name := frmt.Vars[i].Name
		sep := frmt.Vars[i].Seperator
		idx := 0

		// if there's a sperator, try to find it
		if sep != "" {
			idx = strings.Index(line, sep)
			if idx < 0 {
				return Entry{}, fmt.Errorf("%s: %w", sep, ErrSepNotFound)
			}
		}

		// if we found the seperator, and there's a variable name
		if idx >= 0 && name != "" {
			value := line[:idx]
			err := assignToEntry(&ent, name, value)
			if err != nil {
				return Entry{}, err
			}

			//truncate line
			start := idx + len(sep)
			line = line[start:]
		}
	}

	return ent, nil
}

// assignToEntry convert a NGINX log variable value to it's appropriate GoLang type
// and save it into the entry struct. It blindly overwrite content without punity.
func assignToEntry(ent *Entry, name, val string) error {
	// first check if the value is an HTTP header
	if strings.Index(name, "http_") == 0 {
		// if the map is empty, initialize it
		if len(ent.Http) == 0 {
			ent.Http = make(map[string]string)
		}
		// the [5:] removes the "http_" from the index
		ent.Http[name[5:]] = val
		return nil
	}

	switch name {
	default:
		return fmt.Errorf("%s: %w", name, ErrVarNotExist)
	case "body_bytes_sent":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.BodyBytesSent = i //int64
	case "bytes_sent":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.BytesSent = i //int64
	case "connection":
		ent.Connection = val //string
	case "connection_requests":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.ConnectionRequests = i //int64
	case "msec":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.Msec = i //int64
		// pipe possible values are "p" for piped or "." for not piped
	case "pipe":
		var b bool
		if val == "p" {
			b = true
		} else if val == "." {
			b = false
		} else {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.Pipe = b //bool
	case "remote_addr":
		addr := net.ParseIP(val)
		if addr == nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrIpConvFail)
		}
		ent.RemoteAddr = addr //net.IP
	case "remote_user":
		ent.RemoteUser = val //string
	case "request":
		ent.Request = val //string
	case "request_length":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.RequestLength = i //int64
	case "request_time":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.RequestTime = i //int64
	case "status":
		i, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.Status = i //int64
	case "time_iso8601":
		t, e := time.Parse(TimeLayoutIso8601, val)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.TimeIso8601 = t //time.Time
	case "time_local":
		t, e := time.Parse(TimeLayoutLocal, val)
		if e != nil {
			return fmt.Errorf("%s: %s: %w", name, val, ErrConverFail)
		}
		ent.TimeLocal = t //time.Time
	}
	return nil
}
