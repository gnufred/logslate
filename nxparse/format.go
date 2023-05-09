package nxparse

import "strings"

// Default `combined` `log_format`.
const (
	Combined = "$remote_addr - $remote_user [$time_local] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\""
)

// Pre defined FormatVar. They should only be used with formatCache and tests.
var (
	// Default `combined` `log_format`.
	combinedVars = map[int]FormatVar{
		0: {"remote_addr", " - "},
		1: {"remote_user", " ["},
		2: {"time_local", "] \""},
		3: {"request", "\" "},
		4: {"status", " "},
		5: {"body_bytes_sent", " \""},
		6: {"http_referer", "\" \""},
		7: {"http_user_agent", "\""},
	}
	// Nil value.
	nilVars = map[int]FormatVar{
		0: {"", ""},
	}
)

// A Format represent the NGINX `log_format` directive.
type Format struct {
	Vars map[int]FormatVar
}

// A FormatVar represent single variable from the NGINX `log_format` directive.
type FormatVar struct {
	Name      string
	Seperator string
}

// A varPosiMatrix represent a position matrix for a FormatVar. It has a matrix for
// both the variable and the seperator.
type varPosMatrix struct {
	Var posMatrix
	Sep posMatrix
}
type posMatrix struct {
	Beg int
	End int
}

// parseFormat parse the NGNINX `log_format` string into structured data that we can
// use to parse the log line.
//
// To parse to format we need to loop through the string one character at a time and
// understand what's going on.
//
// NGINGX `log_format` identify variables as a string that start with a '$' followed
// by [a-z_] characters. After that comes the seperator, which can include any
// characters.
//
// Two special cases can occurs.
// First case: We can have a seperator, or no seperator, before the first variable.
// This means that the FormatVar[0] could have two zero values.
// Second case: The last variable, has no seperator.
func parseFormat(format string) Format {
	frmt := Format{Vars: make(map[int]FormatVar)}
	pos := varPosMatrix{}

	inVar := false
	lastChar := len(format) - 1

	// First special case: There's a sperator before the first variable.
	if dolIdx := strings.Index(format, "$"); dolIdx > 0 {
		frmt.Vars[0] = FormatVar{"", format[0:dolIdx]}
	}

	for i, c := range format {

		// Scan for var name start.
		if c == '$' {
			inVar = true
			pos.Sep.End = i - 1

			// It's not the first variable we encounter.
			if pos.Var.Beg > 0 {
				frmt.appendVar(format, pos)
			}

			pos.Var.Beg = i + 1
		} else

		// Scan for var name ending.
		if inVar && !isVarNameChar(c) {
			inVar = false
			pos.Var.End = i - 1
			pos.Sep.Beg = i
		}

		// Handle last character properly.
		if i == lastChar {
			if inVar {
				pos.Var.End = i
				pos.Sep.Beg = i
			}
			pos.Sep.End = i
			frmt.appendVar(format, pos)
		}
	}

	return frmt
}

// isVarNameChar check if a character is a valid character for an NGINX `log_format`
// variable. Valid characters [a-z_].
func isVarNameChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || c == '_'
}

// Format.appendVar appends a new FormatVar at the end of a FormatVar map
func (f *Format) appendVar(format string, pos varPosMatrix) {
	s := ""
	v := format[pos.Var.Beg : pos.Var.End+1]

	// Handle second special case: last variable might not have a seperator.
	if varHasSep(pos) {
		s = format[pos.Sep.Beg : pos.Sep.End+1]
	}

	vLen := len(f.Vars)
	f.Vars[vLen] = FormatVar{v, s}
}

// varHasSep check that a variable has seperators based on the position matrix. If
// the the var end position equals both the seperator start and end position, it
// doesn't have a seperator.
func varHasSep(pos varPosMatrix) bool {
	return pos.Var.End != pos.Sep.Beg && pos.Var.End != pos.Sep.End
}
