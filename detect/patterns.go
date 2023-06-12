package detect

// Attack types.
// Useful resource https://owasp.org/www-community/attacks/
// I use bytes over strings because I want to use the smallest possible values.
const (
	Gen = byte(1) // SQL injection
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
var attackPts = []ptrn{
	// Common attacks patterns.
	{Name: ",,", Type: Gen, Score: 4},
	{Name: "%40%40", Type: Gen, Score: 4},
	{Name: "\x22", Type: Gen, Score: 5},
	{Name: "\x22()", Type: Gen, Score: 5},
	{Name: "(((", Type: Gen, Score: 5},
	{Name: ")))", Type: Gen, Score: 5},
	// SQL injection attacks.
	{Name: "'", Type: Gen, Score: 4},
	{Name: "'=", Type: Gen, Score: 5},
	{Name: "' ", Type: Gen, Score: 5},
	{Name: "!='", Type: Gen, Score: 5},
	{Name: "!%2b!", Type: Gen, Score: 5},
	{Name: "%21+%21", Type: Gen, Score: 5},
	{Name: "'%", Type: Gen, Score: 5},
	{Name: "!((", Type: Gen, Score: 5},
	{Name: "1=0", Type: Gen, Score: 5},
	{Name: "1=1", Type: Gen, Score: 5},
	{Name: "'and", Type: SQL, Score: 5},
	{Name: "%20and%", Type: SQL, Score: 5},
	{Name: "'or", Type: SQL, Score: 5},
	{Name: "or%", Type: SQL, Score: 5},
	{Name: "'xor", Type: SQL, Score: 5},
	{Name: "xor\\", Type: SQL, Score: 5},
	{Name: "sleep(", Type: SQL, Score: 5},
	{Name: "sleep%", Type: SQL, Score: 5},
	{Name: "sleep+", Type: SQL, Score: 5},
	{Name: "waitfor(", Type: SQL, Score: 5},
	{Name: "waitfor%", Type: SQL, Score: 5},
	{Name: "waitfor+", Type: SQL, Score: 5},
	{Name: "delay(", Type: SQL, Score: 5},
	{Name: "delay%", Type: SQL, Score: 5},
	{Name: "delay+", Type: SQL, Score: 5},
	{Name: "select(", Type: SQL, Score: 5},
	{Name: "select%", Type: SQL, Score: 5},
	{Name: "select+", Type: SQL, Score: 5},
	{Name: "benchmark(", Type: SQL, Score: 5},
	{Name: "benchmark%", Type: SQL, Score: 5},
	{Name: "benchmark+", Type: SQL, Score: 5},
	{Name: "pg_sleep", Type: SQL, Score: 5},
	// Directory traversal attacks.
	{Name: "./", Type: Dir, Score: 4},
	{Name: ".%2f", Type: Dir, Score: 4},
	{Name: ".0x5c", Type: Dir, Score: 4},
	{Name: ".%5c", Type: Dir, Score: 4},
	{Name: ".%252f", Type: Dir, Score: 4},
	{Name: ".%c0%af", Type: Dir, Score: 4},
	{Name: ".%c1%9c", Type: Dir, Score: 4},
	{Name: ".%25c0", Type: Dir, Score: 4},
	{Name: ".%2550", Type: Dir, Score: 4},
	{Name: ".%10c", Type: Dir, Score: 4},
	{Name: ".0x5c", Type: Dir, Score: 4},
	{Name: "%2e%5c", Type: Dir, Score: 4},
	{Name: "etc/", Type: Dir, Score: 2},
	{Name: "etc%2f", Type: Dir, Score: 5},
	{Name: "%2fstatic", Type: Dir, Score: 5},
	{Name: "web-inf", Type: Dir, Score: 5},
	{Name: "c:%", Type: Dir, Score: 5},
	{Name: "windows", Type: Dir, Score: 2},
	{Name: ".ini", Type: Dir, Score: 2},
	{Name: "passwd", Type: Dir, Score: 3},
	{Name: "etc/passwd", Type: Dir, Score: 5},
	// Command injection attacks.
	{Name: "concat(", Type: CMD, Score: 5},
	{Name: "concat%", Type: CMD, Score: 5},
	{Name: "echo%", Type: CMD, Score: 5},
	{Name: "socket.", Type: CMD, Score: 5},
	{Name: "socket\\", Type: CMD, Score: 5},
	{Name: "socket'", Type: CMD, Score: 5},
	{Name: "socket%", Type: CMD, Score: 5},
	{Name: "nslookup", Type: CMD, Score: 5},
	{Name: "gethostbyname", Type: CMD, Score: 5},
	{Name: "echo", Type: CMD, Score: 5},
	{Name: "=var_dump", Type: CMD, Score: 5},
	{Name: "=phpinfo", Type: CMD, Score: 5},
	{Name: "pipe.receive", Type: CMD, Score: 5},
	// Cross site scripting attacks.
	{Name: "<!--", Type: XSS, Score: 5},
	{Name: "assert(", Type: XSS, Score: 5},
	{Name: "esi:include", Type: XSS, Score: 5},
	{Name: "response.write", Type: XSS, Score: 5},
	{Name: "onload=", Type: XSS, Score: 5},
	{Name: "iframe_src", Type: XSS, Score: 5},
	{Name: "print%", Type: XSS, Score: 5},
	{Name: "print'", Type: XSS, Score: 5},
	{Name: "print(", Type: XSS, Score: 5},
	{Name: "md5%", Type: XSS, Score: 5},
	{Name: "md5(", Type: XSS, Score: 5},
	{Name: "onerror=", Type: XSS, Score: 5},
	{Name: "onerror%", Type: XSS, Score: 5},
	{Name: "onmouseover=", Type: XSS, Score: 5},
	{Name: "onmouseover%", Type: XSS, Score: 5},
	{Name: "focus=", Type: XSS, Score: 5},
	{Name: "focus%", Type: XSS, Score: 5},
	{Name: "injected_by", Type: XSS, Score: 5},
	{Name: "script", Type: XSS, Score: 2},
	{Name: "script%", Type: XSS, Score: 5},
	{Name: "script=", Type: XSS, Score: 5},
	{Name: "script&gt", Type: XSS, Score: 5},
	{Name: "script>", Type: XSS, Score: 5},
	{Name: "replace", Type: XSS, Score: 2},
	{Name: "replace;", Type: XSS, Score: 5},
	{Name: "replace(", Type: XSS, Score: 5},
	{Name: "replace\\", Type: XSS, Score: 5},
	{Name: "iframe%", Type: XSS, Score: 5},
	{Name: "iframe_", Type: XSS, Score: 5},
	{Name: "write(", Type: XSS, Score: 5},
	{Name: "write%", Type: XSS, Score: 5},
}

// AttackTypeString translate the attack type byte to a string to make
// the result human readable.
func AttackTypeString(attack byte) string {
	switch attack {
	case Gen:
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
