package threat

// I use bytes over strings because I want to use the smallest possible values.
type AttackType byte

const (
	Nil AttackType = iota
	Unk            // Unknown or uncategorized
	SQL            // SQL injection
	CMD            // Command injection
	Dir            // Directory traversal
	XSS            // Cross site scripting
	Scn            // Leak & portal scanning
)

// Attack pattern structure.
// The Score scales from 1 to 10 meaning:
// * 1: it might be an attack but probably not.
// * 2-3: scale in between those two.
// * 5: we're pretty sure it's an attack.
// * 6+: not used.
type ptrn struct {
	Name  string
	Type  AttackType
	Score byte
}

// Attack patterns.
// Patterns that are commented with Score: 0, have to many false positive to be
// useful. I'm keeping them commented to avoid putting them back in the future.
var attackPatterns = []ptrn{

	// !!! Unknown attacks patterns !!! //
	// Patterns I've seen mixed in with other attacks.
	{Name: "%40%40", Type: Unk, Score: 2},
	{Name: ",,", Type: Unk, Score: 2},
	// {Name: "***", Type: Unk, Score: 1}, // False positive: w2.02, can be used in "hidden" credit card numbers
	{Name: "(((", Type: Unk, Score: 5},
	{Name: ")))", Type: Unk, Score: 5},
	// {Name: "}}", Type: Unk, Score: 0}, // False positive: w1.01
	// {Name: "{{", Type: Unk, Score: 0}, // False positive: w1.01
	{Name: "'+", Type: SQL, Score: 2},   // In XSS, SQL, and legit URLs.
	{Name: "\x22", Type: Unk, Score: 3}, // All over the place and never alone.

	// !!! SQL injection attacks !!! //
	// https://portswigger.net/web-security/sql-injection/cheat-sheet

	// Generic select.
	// "select%" is all over the place in GraphQL queries.
	// {Name: "select%", Type: SQL, Score: 0},
	{Name: "select+", Type: SQL, Score: 2}, // Some legit hits.
	{Name: "select ", Type: SQL, Score: 5},
	{Name: "select(", Type: SQL, Score: 5},

	// Substring.
	{Name: "substr(", Type: SQL, Score: 5},
	{Name: "substr%", Type: SQL, Score: 5},
	{Name: "substring(", Type: SQL, Score: 5},
	{Name: "substring%", Type: SQL, Score: 5},

	// Time delays.
	{Name: "receive_message", Type: CMD, Score: 5},
	{Name: "waitfor%", Type: SQL, Score: 5},
	{Name: "waitfor(", Type: SQL, Score: 5},
	{Name: "waitfor+", Type: SQL, Score: 5},
	// {Name: "delay", Type: SQL, Score: 0},
	{Name: "delay+", Type: SQL, Score: 2},
	{Name: "delay%", Type: SQL, Score: 4},
	{Name: "delay(", Type: SQL, Score: 5},
	{Name: "pg_sleep", Type: SQL, Score: 5},
	// {Name: "sleep", Type: SQL, Score: 0},
	{Name: "sleep%", Type: SQL, Score: 2},
	{Name: "sleep+", Type: SQL, Score: 2},
	{Name: "sleep(", Type: SQL, Score: 5},
	// {Name: "benchmark", Type: SQL, Score: 0},
	{Name: "benchmark(", Type: SQL, Score: 5},
	{Name: "benchmark%", Type: SQL, Score: 5},
	{Name: "benchmark+", Type: SQL, Score: 5},

	// DNS lookup.
	{Name: "extractvalue", Type: SQL, Score: 5},
	{Name: "get_host_address", Type: SQL, Score: 5},
	{Name: "exec master", Type: SQL, Score: 5},
	{Name: "load_file(", Type: SQL, Score: 5},
	{Name: "load_file%28", Type: SQL, Score: 5},

	// Uncategorized.
	{Name: "--+", Type: SQL, Score: 5},
	{Name: "' ", Type: SQL, Score: 5},
	{Name: "')", Type: SQL, Score: 5},
	{Name: "'|", Type: SQL, Score: 5},
	{Name: "''", Type: SQL, Score: 4},
	{Name: "'=", Type: SQL, Score: 5},
	{Name: "!='", Type: SQL, Score: 5},
	{Name: "!+!", Type: SQL, Score: 5},
	{Name: "!%2b!", Type: SQL, Score: 5},
	{Name: "%21+%21", Type: SQL, Score: 5},
	{Name: "'%", Type: SQL, Score: 2},
	{Name: "'%22", Type: SQL, Score: 5},
	{Name: "'%bf", Type: SQL, Score: 5},
	{Name: "'%f0", Type: SQL, Score: 5},
	{Name: "'%7d", Type: SQL, Score: 5},
	{Name: "!((", Type: SQL, Score: 4},
	{Name: "'1=0", Type: SQL, Score: 5},
	{Name: "'1=1", Type: SQL, Score: 5},
	{Name: "'and", Type: SQL, Score: 5},
	{Name: "and%20'", Type: SQL, Score: 5},
	{Name: "'%20and", Type: SQL, Score: 5},
	{Name: "'or", Type: SQL, Score: 5},
	{Name: "or%20'", Type: SQL, Score: 5},
	{Name: "'%20or", Type: SQL, Score: 5},
	{Name: "'xor", Type: SQL, Score: 5},
	{Name: "xor\\", Type: SQL, Score: 5},
	{Name: ",null", Type: SQL, Score: 5},
	{Name: "null:null", Type: SQL, Score: 5},
	{Name: "qsqli_", Type: SQL, Score: 5},

	// !!! Directory traversal attacks !!! //
	// https://portswigger.net/web-security/file-path-traversal
	{Name: "%252e", Type: Dir, Score: 2}, // Some legit hits.
	// {Name: "%252f", Type: Dir, Score: 0},
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
	{Name: "./", Type: Dir, Score: 2},
	{Name: ".0x5c", Type: Dir, Score: 4},
	// {Name: "\x5c", Type: Dir, Score: 0},
	{Name: ".ini", Type: Dir, Score: 2},
	{Name: "win.ini", Type: Dir, Score: 5},
	{Name: "webadmin.ini", Type: Dir, Score: 5},
	{Name: "c:%", Type: Dir, Score: 5},
	{Name: "etc%2f", Type: Dir, Score: 5},
	{Name: "etc%5c", Type: Dir, Score: 5},
	{Name: "etc/", Type: Dir, Score: 2},
	{Name: "etc/passwd", Type: Dir, Score: 5},
	{Name: "passwd", Type: Dir, Score: 3},
	{Name: "web-inf", Type: Dir, Score: 5},
	{Name: "file://", Type: Dir, Score: 4},
	{Name: "file:%2f%2f", Type: Dir, Score: 5},

	// !!! OS command injection attacks !!! //
	// https://portswigger.net/web-security/os-command-injection
	{Name: "concat(", Type: CMD, Score: 5}, // In SQL injec but mostly in OS inject.
	{Name: "concat%", Type: CMD, Score: 5},
	{Name: "chr(", Type: CMD, Score: 5},
	{Name: "chr%", Type: CMD, Score: 1}, // Can match German "spiegelschr".
	{Name: "hex(", Type: CMD, Score: 5},
	{Name: "echo", Type: CMD, Score: 1}, // In legit URLs and attacks.
	{Name: "gethostbyname", Type: CMD, Score: 5},
	{Name: "ipconfig", Type: CMD, Score: 5},
	{Name: "netstat", Type: CMD, Score: 5},
	{Name: "nslookup", Type: CMD, Score: 5},
	{Name: "ps -", Type: CMD, Score: 5},
	{Name: "ps%20%2D", Type: CMD, Score: 5},
	{Name: "tasklist", Type: CMD, Score: 2},
	// "uname" matches lots of words.
	{Name: "uname", Type: CMD, Score: 1},
	{Name: "whoami", Type: CMD, Score: 5},
	{Name: "socket", Type: CMD, Score: 1}, // Legit word. In attacks, with other triggers.
	{Name: "%00{.", Type: CMD, Score: 5},  // CVE-2014-6287 https://packetstormsecurity.com/files/128243/HttpFileServer-2.3.x-Remote-Command-Execution.html

	// !!! Cross site scripting attacks !!! //
	// https://portswigger.net/web-security/cross-site-scripting/cheat-sheet
	// Wasn't nearly as complete as what I've seen. Added my observations.

	// JS and browser function injections.
	{Name: "alert%", Type: XSS, Score: 2}, // In legit URLs.
	{Name: "alert(", Type: XSS, Score: 5},
	{Name: "assert%", Type: XSS, Score: 5},
	{Name: "assert(", Type: XSS, Score: 5},
	// {Name: "log", Type: XSS, Score: 0}, // Hits "blog", "catalog", and such.
	{Name: "log%28", Type: XSS, Score: 1},
	{Name: "log(", Type: XSS, Score: 5},
	{Name: "md5%", Type: XSS, Score: 2},
	{Name: "md5(", Type: XSS, Score: 5},
	{Name: "print%", Type: XSS, Score: 2},
	{Name: "print'", Type: XSS, Score: 5},
	{Name: "print(", Type: XSS, Score: 4}, // Has some legit uses.
	{Name: "replace(", Type: XSS, Score: 5},
	{Name: "replace;", Type: XSS, Score: 5},
	{Name: "replace\\", Type: XSS, Score: 5},
	{Name: "sendBeacon%", Type: XSS, Score: 5},
	{Name: "sendBeacon(", Type: XSS, Score: 5},
	{Name: "write%", Type: XSS, Score: 2}, // Has legit uses.
	{Name: "write(", Type: XSS, Score: 4},
	{Name: "random(", Type: XSS, Score: 4},
	{Name: "function(", Type: XSS, Score: 4},

	// HTML, JS tag injection.
	{Name: "<!--", Type: XSS, Score: 5},
	{Name: "=)(", Type: XSS, Score: 5},
	{Name: "%3D%29%28", Type: XSS, Score: 5},
	{Name: "esi:include", Type: XSS, Score: 5},
	{Name: "iframe%", Type: XSS, Score: 1},
	{Name: "iframe_", Type: XSS, Score: 2},
	{Name: "iframe_src", Type: XSS, Score: 5},
	{Name: "injected_by", Type: XSS, Score: 5},
	{Name: "script=", Type: XSS, Score: 4}, // False positive: w2.03
	{Name: "script>", Type: XSS, Score: 5},
	{Name: "script&gt", Type: XSS, Score: 5},
	{Name: "script%3e", Type: XSS, Score: 5},
	{Name: "script%2f%26", Type: XSS, Score: 5},
	{Name: "%253a%252f%252flocalhost%252f", Type: XSS, Score: 5}, // Double URL encoded "://localhost/".
	{Name: "document.domain", Type: XSS, Score: 5},
	{Name: "void%280%29", Type: XSS, Score: 5}, // From URL encoded "javascript:void(0)".

	// Config scans.
	{Name: "Cargo.lock", Type: Scn, Score: 5},
	{Name: "Cargo.toml", Type: Scn, Score: 5},
	{Name: "DLPCenter/loginform.sms", Type: Scn, Score: 5},
	{Name: "Dockerrun.aws.json", Type: Scn, Score: 5},
	{Name: "Gemfile", Type: Scn, Score: 5},
	{Name: "Guardfile", Type: Scn, Score: 5},
	{Name: "Login.action", Type: Scn, Score: 5},
	{Name: "Pipfile", Type: Scn, Score: 5},
	{Name: "Pipfile.lock", Type: Scn, Score: 5},
	{Name: "Procfile", Type: Scn, Score: 5},
	{Name: "apps.tidy.infinity.json", Type: Scn, Score: 5},
	{Name: "auth.js", Type: Scn, Score: 5},
	{Name: "aws.sh", Type: Scn, Score: 5},
	{Name: "awstats.conf", Type: Scn, Score: 5},
	{Name: "azure-pipelines.yml", Type: Scn, Score: 5},
	{Name: "azuredeploy.json", Type: Scn, Score: 5},
	{Name: "bitbucket-pipelines.yml", Type: Scn, Score: 5},
	{Name: "config.html", Type: Scn, Score: 5},
	{Name: "config.inc.php", Type: Scn, Score: 5},
	{Name: "config.js", Type: Scn, Score: 5},
	{Name: "credentials.db", Type: Scn, Score: 5},
	{Name: "crossdomain.xml", Type: Scn, Score: 5},
	{Name: "dataOrigin.ashx", Type: Scn, Score: 5},
	{Name: "databases.yml", Type: Scn, Score: 5},
	{Name: "debug.log", Type: Scn, Score: 5},
	{Name: "default.json", Type: Scn, Score: 5},
	{Name: "deploy.sh", Type: Scn, Score: 5},
	{Name: "desktop.ini", Type: Scn, Score: 5},
	{Name: "development.log", Type: Scn, Score: 5},
	{Name: "development.log", Type: Scn, Score: 5},
	{Name: "env.js", Type: Scn, Score: 5},
	{Name: "env.sh", Type: Scn, Score: 5},
	{Name: "error.log", Type: Scn, Score: 5},
	{Name: "errors.log", Type: Scn, Score: 5},
	{Name: "firewall.log", Type: Scn, Score: 5},
	{Name: "go.mod", Type: Scn, Score: 5},
	{Name: "htaccess.txt", Type: Scn, Score: 5},
	{Name: "init.sh", Type: Scn, Score: 5},
	{Name: "karma.conf.js", Type: Scn, Score: 5},
	{Name: "karma.js", Type: Scn, Score: 5},
	{Name: "keycloak.json", Type: Scn, Score: 5},
	{Name: "keys.js", Type: Scn, Score: 5},
	{Name: "kustomization.yml", Type: Scn, Score: 5},
	{Name: "login.action", Type: Scn, Score: 5},
	{Name: "mobile.log", Type: Scn, Score: 5},
	{Name: "privatekey.key", Type: Scn, Score: 5},
	{Name: "production.log", Type: Scn, Score: 5},
	{Name: "run.sh", Type: Scn, Score: 5},
	{Name: "secret_token.rb", Type: Scn, Score: 5},
	{Name: "security.yaml", Type: Scn, Score: 5},
	{Name: "sendgrid.ev", Type: Scn, Score: 5},
	{Name: "settings.py", Type: Scn, Score: 5},
	{Name: "settings.yml", Type: Scn, Score: 5},
	{Name: "user.xml", Type: Scn, Score: 5},
	{Name: "vpn.log", Type: Scn, Score: 5},

	// Admin panel scans.
	// As of right now, I'm conservative.
	// The idea is to flag paths that aren't used by legit users.
	// I might change that later, but given the simplist nature of the detectiong logic,
	// I'm keeping it simple.
	// In any case, attackers will almost always scan weird paths that nobody uses.
	{Name: "/cgi-bin", Type: Scn, Score: 5},
	{Name: "/setup.php", Type: Scn, Score: 5},
	{Name: "CgiStart", Type: Scn, Score: 5},
	{Name: "FNDWRR.exe", Type: Scn, Score: 5},
	{Name: "Telerik.Web.UI.DialogHandler", Type: Scn, Score: 5},
	{Name: "b9dmx.php", Type: Scn, Score: 5},
	{Name: "carbon/admin/login.jsp", Type: Scn, Score: 5},
	{Name: "cfg/login", Type: Scn, Score: 5},
	{Name: "phpmyadmin", Type: Scn, Score: 5},
	{Name: "qmailadmin.cgi", Type: Scn, Score: 5},
	{Name: "backup/auto.php", Type: Scn, Score: 5},
	{Name: "backup2.cgi", Type: Scn, Score: 5},
	{Name: "cache/backup", Type: Scn, Score: 5},

	// Data dump scans.
	{Name: ".sql", Type: Scn, Score: 3},
	{Name: ".7z", Type: Scn, Score: 3},
	{Name: ".bz2", Type: Scn, Score: 3},
	{Name: ".gz", Type: Scn, Score: 3},
	{Name: ".lz", Type: Scn, Score: 3},
	{Name: ".rar", Type: Scn, Score: 3},
	{Name: ".tar", Type: Scn, Score: 3},
	{Name: ".zip", Type: Scn, Score: 3},

	{Name: "/backup", Type: Scn, Score: 2},
	{Name: "/bak", Type: Scn, Score: 2},
	{Name: "/db", Type: Scn, Score: 2},
	{Name: "/db_backup", Type: Scn, Score: 2},
	{Name: "/dbdump", Type: Scn, Score: 2},
	{Name: "/dump", Type: Scn, Score: 2},
	{Name: "/html", Type: Scn, Score: 2},
	{Name: "/old", Type: Scn, Score: 2},
	{Name: "/public", Type: Scn, Score: 2},
	{Name: "/release", Type: Scn, Score: 2},
	{Name: "/root", Type: Scn, Score: 2},
	{Name: "/test", Type: Scn, Score: 2},
	{Name: "/www", Type: Scn, Score: 2},
}

// String translate an attack type byte to a human readable string
func String(atk AttackType) string {
	switch atk {
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
	case Scn:
		return "scan"
	default:
		return ""
	}
}
