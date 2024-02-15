package threat

import "testing"

var testPatterns = []ptrn{
	{Name: "x", Type: Unk, Score: 1},
	{Name: "xy", Type: Unk, Score: 2},
	{Name: "xyz", Type: Unk, Score: 3},

	{Name: "\x99()", Type: Unk, Score: 255},
	{Name: "'SLEEP(0)", Type: SQL, Score: 255},
	{Name: "socket.hitmycommandcenter", Type: CMD, Score: 255},
	{Name: "../etc/passwd", Type: Dir, Score: 255},
	{Name: "<script>", Type: XSS, Score: 255},
	{Name: "/dump.sql", Type: Scn, Score: 255},
}

// TestFalsePositives tests squashed false positives after 2023-07-06. This will avoid
// detecting them again when adding new attack patterns.
// IMPORTANT: Listed queries need to be "anonymized" as best as possible.
func TestFalsePositives(t *testing.T) {
	detectionShouldNotTrigger := func(t *testing.T, id int, query, path string) {
		atkType, atkScore := attackScan(path, query)
		if atkType != 0 {
			t.Errorf(
				"%d: false positive found, attak type != 0, got %d",
				id, atkType,
			)
		}
		if atkScore != 0 {
			t.Errorf(
				"%d: false positive found, attak score != 0, got %d",
				id, atkScore,
			)
		}
	}

	falsePositives := []struct {
		// The ID is used in failure output to make it easier to identify what failed.
		// Format is: wX (week X) . int increment.
		Id    int
		Path  string
		Query string
	}{
		{1, "/gql", "q={p(cP%3A1%2CpS%3A60%2Cf%3A%7Bcid%3A%7Beq%3A%2217076%22%7D%2Cdata%3A%7Bin%3A%5B%22%22%5D%7D%7D%2Cs%3A%22l%27%C3%A9du%20me%22%2Cresolver3A1){tc,ges,nst,dm,ag{o{c,l,v},a,c,l,i,e},pi{c{l,v},c{l,v},s%20{l,v},p,e,t,__t},i%20{__t,i,s,n,e,i{n,u},s{u},k,k{u},u,t,e,c,c,e,r,n,f,a,m{o{o,r,p,t,s,o},t,t,l{v},l{v},s{v},s{v}},m,m%20{l,v},g{g,m},mc_:%20s_m,s_v_c_:%20s_v,f_b_c_:%20f,b_i_c_:%20b_i,s_p_:%20s_p}}}"},
		{2, "/", "p[c][b]=R%7CMr%20%26amp%3B%20%C3%A9t%20CHR%7CM%20de%20s&p[c][i]=5&e[c][l]=P%20de%20t%20int%C3%A9"},
		{3, "/", "d=%7B%22d%22%3A%22S%20%20Up%20-%20h%20s%201%20%5Cu00e0%202%20j%20r%5Cu00e9s2C%22c%22al%22%2C%22c%22%3A%22ed%22%2C%22display_iframe%22%3A0%7D&118&t=0&a=4&au=&co=F&sc=0&fr=ACD&re=&av=&cvk=&pp=&eci3ds=5&v=&ct=&pa=Y&cb=fd&cp=1111111******1111&ce=111111&cc=F&cid=111111&id=11111111l&r=1111111111&c=E&a=11.00&h="},
		{4, "/", "j-m=sx&aL=s&ar=yCE&bS=sx&d=ssx&ex=s"},
		{5, "/gql", "{p(f:+{s:+{in:[%22ER%22,%22EER%22,%22R%22,%22E1AER%22,%221AER%22,%222AER%22,%22E1AER%22]}},"},
		{6, "/test/acc/test-l-p-c/test-l-s/index.htm", ""},
		{7, "/index.htm", "%20http://www.test.com/ridgid/video/reels/index.htm,%20http://www.test.com/drain/index.htm,%20http://www.test.com/drain/index.htm"},
		{8, "/c/p/v/id/12345", "'%20+%20p.key%20+%20(p.suffix%20%7C%7C%20'')"},
	}

	for _, d := range falsePositives {
		detectionShouldNotTrigger(t, d.Id, d.Path, d.Query)
	}
}

func TestAttackTypeString(t *testing.T) {
	testMatch := func(t *testing.T, in AttackType, exp string) {
		got := String(in)
		if got != exp {
			t.Errorf("attack type string doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		In  AttackType
		Exp string
	}{
		{Nil, ""},
		{Unk, "unknown"},
		{SQL, "sql_inject"},
		{CMD, "cmd_inject"},
		{Dir, "dir_travers"},
		{XSS, "xss"},
		{Scn, "scan"},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestDataDumpPatterns(t *testing.T) {
	pts := DataDumpPatterns()
	found := false
	search := "www.zip"

	for _, v := range pts {
		if v.Name == search {
			found = true
		}
	}
	if !found {
		t.Errorf("did not find %s", search)
	}
}
