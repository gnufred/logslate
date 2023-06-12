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
// * 6-9: not used.
// * 10: absolutely an attack.
// * 11+: not used.
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
	{Name: "\x22", Type: Gen, Score: 5},
	{Name: "\x22()", Type: Gen, Score: 10},
	// SQL injection attacks.
	{Name: "'", Type: Gen, Score: 4},
	{Name: "1=1", Type: Gen, Score: 4},
	{Name: "%20and", Type: SQL, Score: 5},
	{Name: "%20and", Type: SQL, Score: 5},
	{Name: "%20or", Type: SQL, Score: 5},
	{Name: "%22xor", Type: SQL, Score: 5},
	{Name: "%'Xxor", Type: SQL, Score: 5},
	{Name: "%29,sleep", Type: SQL, Score: 5},
	{Name: "%28sleep", Type: SQL, Score: 5},
	{Name: "sleep(", Type: SQL, Score: 5},
	{Name: "pg_sleep", Type: SQL, Score: 10},
	{Name: "%20waitfor", Type: SQL, Score: 1},
	{Name: "%20delay", Type: SQL, Score: 1},
	{Name: "%2Csleep", Type: SQL, Score: 1},
	{Name: "%27%7c%7c%27", Type: SQL, Score: 5},
	{Name: "pipe.receive_message", Type: SQL, Score: 10},
	{Name: "select(", Type: SQL, Score: 5},
	{Name: "select", Type: SQL, Score: 2},
	// Directory traversal attacks.
	{Name: "./", Type: Dir, Score: 5},
	{Name: ".%2f", Type: Dir, Score: 5},
	{Name: ".0x5c", Type: Dir, Score: 5},
	{Name: ".%5c", Type: Dir, Score: 5},
	{Name: ".%252f", Type: Dir, Score: 5},
	{Name: ".%c0%af", Type: Dir, Score: 5},
	{Name: ".%c1%9c", Type: Dir, Score: 5},
	{Name: ".%25c0", Type: Dir, Score: 5},
	{Name: ".%10c", Type: Dir, Score: 5},
	{Name: ".0x5c", Type: Dir, Score: 5},
	{Name: "%2e%2e%5c", Type: Dir, Score: 5},
	{Name: "etc/", Type: Dir, Score: 2},
	{Name: "etc%2f", Type: Dir, Score: 5},
	{Name: "%2fetc", Type: Dir, Score: 5},
	{Name: "%2fstatic", Type: Dir, Score: 5},
	{Name: "%2fweb-inf", Type: Dir, Score: 5},
	{Name: "c:%", Type: Dir, Score: 5},
	{Name: "windows", Type: Dir, Score: 2},
	{Name: ".ini", Type: Dir, Score: 2},
	{Name: "passwd", Type: Dir, Score: 5},
	{Name: "etc/passwd", Type: Dir, Score: 10},
	// Command injection attacks.
	{Name: "concat", Type: CMD, Score: 5},
	{Name: "socket.", Type: CMD, Score: 5},
	{Name: "socket\\", Type: CMD, Score: 5},
	{Name: "socket'", Type: CMD, Score: 5},
	{Name: "socket%", Type: CMD, Score: 5},
	{Name: "nslookup", Type: CMD, Score: 5},
	{Name: "gethostbyname", Type: CMD, Score: 5},
	{Name: "echo", Type: CMD, Score: 5},
	{Name: "=var_dump", Type: CMD, Score: 5},
	{Name: "=phpinfo", Type: CMD, Score: 5},
	// Cross site scripting attacks.
	{Name: "<!--", Type: XSS, Score: 10},
	{Name: "esi:include", Type: XSS, Score: 10},
	{Name: "response.write", Type: XSS, Score: 10},
	{Name: "body_onload", Type: XSS, Score: 10},
	{Name: "iframe_src", Type: XSS, Score: 10},
	{Name: "print%", Type: XSS, Score: 5},
	{Name: "print'", Type: XSS, Score: 5},
	{Name: "print(", Type: XSS, Score: 5},
	{Name: "md5", Type: XSS, Score: 5},
	{Name: "onmouseover", Type: XSS, Score: 5},
	{Name: "injected_by", Type: XSS, Score: 5},
	{Name: "script", Type: XSS, Score: 2},
	{Name: "script%", Type: XSS, Score: 5},
	{Name: "script=", Type: XSS, Score: 5},
	{Name: "script>", Type: XSS, Score: 5},
	{Name: "replace", Type: XSS, Score: 2},
	{Name: "replace;", Type: XSS, Score: 10},
	{Name: "replace(", Type: XSS, Score: 10},
	{Name: "replace\\", Type: XSS, Score: 10},
	{Name: "iframe", Type: XSS, Score: 10},
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

// TODO: Some misteries to resolve latter
/*
'\x22\x5C'\x5C\x22);|]*%00{%0d%0a<%00>%bf%27'\xF0\x9F\x92\xA1
%27%22
%2Ccaptcha%2C
%2Cdirectory-data%2C
%0abcc:074625.365-19302.365.5e6a4.19475.2@bxss.me
%24{%40print(md5(31337))}
[%24eq]=true
[%24testing]=1
[!%2B!]=
!(()%26%26!%7C*%7C*%7C
%5E(%23%24!%40%23%24)(()))******
%24{10000016%2B10000120}

%64%66%62%5B%5B%24%7B%39%38%39%39%31%2A%39%37%39%39%36%7D%5D%5D%78%63%61
%22%73%54%59%4C%65%3D%27%61%63%75%3A%45%78%70%72%65%2F%2A%2A%2F%53%53%69%6F%6E%28%31%58%4C%35%28%39%37%31%34%29%29%27%62%61%64%3D%22
/1%7D%7D%22%7D%7D%27%7D%7D1%25%3E%22%25%3E%27%25%3E%3C%25=%7B%7B=%7B@%7B%23%7B$%7Bdfb%7D%7D%25%3E
%2522%256F%256E%256D%256F%2575%2573%2565%256F%2576%2565%2572%253D%2527%2575%2548%2543%254A%2528%2539%2534%2530%2531%2534%2529%2527%2562%2561%2564%253D%2522=
%2522900914@=
%2522;9432%2522%2528%2529;%257D%255D9101=
2563=&aHR0cHM6Ly93d3cudGFzc2ltby5jb20vY2hkZQ,,
*/
