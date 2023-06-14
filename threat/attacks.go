package threat

// I use bytes over strings because I want to use the smallest possible values.
const (
	// Attack types.
	Unk = byte(1) // Unknown or uncategorized
	SQL = byte(2) // SQL injection
	CMD = byte(3) // Command injection
	Dir = byte(4) // Directory traversal
	XSS = byte(5) // Cross site scripting
)

// Attack pattern structure.
// The Score scales from 1 to 10 meaning:
// * 1: it might be an attack but probably not.
// * 2-3: scale in between those two.
// * 5: we're pretty sure it's an attack.
// * 6+: not used.
type ptrn struct {
	Name  string
	Type  byte
	Score byte
}

// Attack patterns.
// TODO: In theory, it will be more performant to use a slice over a map, but
// that remains to be benchmarked properly.
var attackPatterns = []ptrn{

	// !!! Unknown attacks patterns !!! //
	// Patterns I've seen mixed in with other attacks.
	{Name: "%40%40", Type: Unk, Score: 2},
	{Name: ",,", Type: Unk, Score: 4},
	{Name: "(((", Type: Unk, Score: 5},
	{Name: ")))", Type: Unk, Score: 5},
	// "\x22" is all over the place but it's never alone.
	{Name: "\x22", Type: Unk, Score: 3},

	// !!! SQL injection attacks !!! //
	// https://portswigger.net/web-security/sql-injection/cheat-sheet

	// Generic select.
	// "select+" can be in legit URLs.
	{Name: "select+", Type: SQL, Score: 2}, // All
	{Name: "select ", Type: SQL, Score: 5}, // All
	{Name: "select(", Type: SQL, Score: 5}, // All
	{Name: "select%", Type: SQL, Score: 5}, // All

	// String concatenation.
	{Name: "'|", Type: SQL, Score: 5}, // Oracle, PSQL
	{Name: "'+", Type: SQL, Score: 5}, // Microsoft
	{Name: "' ", Type: SQL, Score: 5}, // MySQL

	// Substring.
	{Name: "substr(", Type: SQL, Score: 5},    // Oracle
	{Name: "substr%", Type: SQL, Score: 5},    // Oracle
	{Name: "substring(", Type: SQL, Score: 5}, // Microsoft, PSQL, MySQL
	{Name: "substring%", Type: SQL, Score: 5}, // Microsoft, PSQL, MySQL

	// Time delays.
	{Name: "receive_message", Type: CMD, Score: 5}, // Oracle
	// Variants: waitfor( waitfor% waitfor+
	{Name: "waitfor", Type: SQL, Score: 5}, // Microsoft
	// Variants: delay( delay% delay+
	{Name: "delay", Type: SQL, Score: 5},    // Microsoft
	{Name: "pg_sleep", Type: SQL, Score: 5}, // PSQL
	// Can't use "sleep|benchmark" alone. In many legit URLs.
	{Name: "sleep(", Type: SQL, Score: 5},     // PSQL, MySQL
	{Name: "sleep%", Type: SQL, Score: 5},     // PSQL, MySQL
	{Name: "sleep+", Type: SQL, Score: 5},     // PSQL, MySQL
	{Name: "benchmark(", Type: SQL, Score: 5}, // MySQL
	{Name: "benchmark%", Type: SQL, Score: 5}, // MySQL
	{Name: "benchmark+", Type: SQL, Score: 5}, // MySQL

	// DNS lookup.
	{Name: "extractvalue", Type: SQL, Score: 5},     // Oracle
	{Name: "get_host_address", Type: SQL, Score: 5}, // Oracle
	{Name: "exec master", Type: SQL, Score: 5},      // Microsoft
	{Name: "load_file", Type: SQL, Score: 5},        // PSQL

	// Uncategorized.
	{Name: "'", Type: SQL, Score: 4},
	{Name: "'=", Type: SQL, Score: 5},
	{Name: "!='", Type: SQL, Score: 5},
	{Name: "!%2b!", Type: SQL, Score: 5},
	{Name: "%21+%21", Type: SQL, Score: 5},
	{Name: "'%", Type: SQL, Score: 5},
	{Name: "!((", Type: SQL, Score: 5},
	{Name: "1=0", Type: SQL, Score: 5},
	{Name: "1=1", Type: SQL, Score: 5},
	{Name: "'and", Type: SQL, Score: 5},
	{Name: "%20and%", Type: SQL, Score: 5},
	{Name: "'or", Type: SQL, Score: 5},
	{Name: "%20or%", Type: SQL, Score: 5},
	{Name: "'xor", Type: SQL, Score: 5},
	{Name: "xor\\", Type: SQL, Score: 5},

	// !!! Directory traversal attacks !!! //
	// https://portswigger.net/web-security/file-path-traversal
	{Name: "%252e", Type: Dir, Score: 4},
	{Name: "%252f", Type: Dir, Score: 4},
	{Name: "%2e%2f", Type: Dir, Score: 4},
	{Name: "%2e%5c", Type: Dir, Score: 4},
	{Name: ".%10c", Type: Dir, Score: 4},
	{Name: ".%2550", Type: Dir, Score: 4},
	{Name: ".%25c0", Type: Dir, Score: 4},
	{Name: ".%2f", Type: Dir, Score: 4},
	{Name: ".%5c", Type: Dir, Score: 4},
	{Name: ".%c0%af", Type: Dir, Score: 4},
	{Name: ".%c1%9c", Type: Dir, Score: 4},
	{Name: ".%ef%bc%8f", Type: Dir, Score: 4},
	{Name: "./", Type: Dir, Score: 4},
	{Name: ".0x5c", Type: Dir, Score: 4},
	{Name: ".0x5c", Type: Dir, Score: 4},
	{Name: ".ini", Type: Dir, Score: 2},
	{Name: "c:%", Type: Dir, Score: 5},
	{Name: "etc%2f", Type: Dir, Score: 5},
	{Name: "etc/", Type: Dir, Score: 2},
	{Name: "etc/passwd", Type: Dir, Score: 5},
	{Name: "passwd", Type: Dir, Score: 3},
	{Name: "web-inf", Type: Dir, Score: 5},
	{Name: "windows", Type: Dir, Score: 2},

	// !!! OS command injection attacks !!! //
	// https://portswigger.net/web-security/os-command-injection
	// "contact" can also be SQL injection but I've seen it mostly.
	// with other OS command injection attack.
	{Name: "concat(", Type: CMD, Score: 5},
	{Name: "concat%", Type: CMD, Score: 5},
	{Name: "chr(", Type: CMD, Score: 5},
	{Name: "chr%", Type: CMD, Score: 5},
	{Name: "hex(", Type: CMD, Score: 5},
	{Name: "hex%", Type: CMD, Score: 5},
	{Name: "echo", Type: CMD, Score: 5},
	{Name: "gethostbyname", Type: CMD, Score: 5},
	{Name: "ipconfig", Type: CMD, Score: 5},
	{Name: "netstat", Type: CMD, Score: 5},
	{Name: "nslookup", Type: CMD, Score: 5},
	{Name: "ps -", Type: CMD, Score: 5},
	{Name: "ps%20", Type: CMD, Score: 5},
	{Name: "tasklist", Type: CMD, Score: 2},
	{Name: "uname", Type: CMD, Score: 5},
	{Name: "whoami", Type: CMD, Score: 5},
	// "socket" is a legit word by it's never alone.
	{Name: "socket", Type: CMD, Score: 2},

	// !!! Cross site scripting attacks !!! //
	// https://portswigger.net/web-security/cross-site-scripting/cheat-sheet
	// Wasn't nearly as complete as what I've seen. Added my observations.

	// JS and browser function injections.
	{Name: "alert%", Type: XSS, Score: 5},
	{Name: "alert(", Type: XSS, Score: 5},
	{Name: "assert%", Type: XSS, Score: 5},
	{Name: "assert(", Type: XSS, Score: 5},
	{Name: "log%", Type: XSS, Score: 5},
	{Name: "log(", Type: XSS, Score: 5},
	{Name: "md5%", Type: XSS, Score: 5},
	{Name: "md5(", Type: XSS, Score: 5},
	{Name: "print%", Type: XSS, Score: 5},
	{Name: "print'", Type: XSS, Score: 5},
	{Name: "print(", Type: XSS, Score: 5},
	{Name: "replace", Type: XSS, Score: 2},
	{Name: "replace(", Type: XSS, Score: 5},
	{Name: "replace;", Type: XSS, Score: 5},
	{Name: "replace\\", Type: XSS, Score: 5},
	{Name: "sendBeacon%", Type: XSS, Score: 5},
	{Name: "sendBeacon(", Type: XSS, Score: 5},
	{Name: "write%", Type: XSS, Score: 5},
	{Name: "write%", Type: XSS, Score: 5},
	{Name: "write(", Type: XSS, Score: 5},
	{Name: "write(", Type: XSS, Score: 5},

	// HTML, JS tag injection.
	{Name: "<!--", Type: XSS, Score: 5},
	{Name: "esi:include", Type: XSS, Score: 5},
	{Name: "iframe%", Type: XSS, Score: 5},
	{Name: "iframe_", Type: XSS, Score: 5},
	{Name: "iframe_src", Type: XSS, Score: 5},
	{Name: "injected_by", Type: XSS, Score: 5},
	{Name: "script%", Type: XSS, Score: 5},
	{Name: "script&gt", Type: XSS, Score: 5},
	{Name: "script=", Type: XSS, Score: 5},
	{Name: "script>", Type: XSS, Score: 5},
}

// String translate an attack type byte to a human readable string
func String(threatType byte) string {
	switch threatType {
	case Unk:
		return "generic"
	case SQL:
		return "sql_injection"
	case CMD:
		return "command_injection"
	case Dir:
		return "directory_traversal"
	case XSS:
		return "xss"
	default:
		return ""
	}
}
