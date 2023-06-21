package threat

import "testing"

type result struct {
	To uint32
	Ty byte
}

func TestDetect(t *testing.T) {
	// Just testing one result. Multiple results tested in sub functions.
	td := struct {
		Path  string
		Query string
		Agent string
		Exp   Threat
	}{
		"user/",
		"id=1'PG_SLEEP(1000)",
		"EdgeScan/1.1 (+https://www.edgescan.com/)",
		Threat{
			Attack{
				2,
				5,
				"sql_injection",
			},
			"EdgeScan",
		},
	}

	got, err := Detect(td.Path, td.Query, td.Agent)
	if err != nil {
		t.Error(err)
	}

	if got.Attack.Id != td.Exp.Attack.Id {
		t.Errorf("attack id doesn't match, got %d, exp %d", got.Attack.Id, td.Exp.Attack.Id)
	}

	if got.Attack.Score != td.Exp.Attack.Score {
		t.Errorf("attack score doesn't match, got %d, exp %d", got.Attack.Score, td.Exp.Attack.Score)
	}

	if got.Attack.Name != td.Exp.Attack.Name {
		t.Errorf("attack name doesn't match, got %s, exp %s", got.Attack.Name, td.Exp.Attack.Name)
	}

	if got.Bot != td.Exp.Bot {
		t.Errorf("bot doesn't match, got %s, exp %s", got.Bot, td.Exp.Bot)
	}
}

func TestAttacks(t *testing.T) {
	testMatch := func(t *testing.T, hay string, exp result) {
		ty, to := attackScan(hay, "")
		got := result{Ty: ty, To: to}
		if got != exp {
			t.Errorf("query total doesn't match, got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp result
	}{
		{"hello(", result{0, 0}},
		{"write(", result{5, 5}},
		{"\x22id\x22:\x221\x22", result{12, 1}},
		{"path/stuff../../etc/passwd", result{13, 4}},
		{"hello'boba_sleep(", result{5, 2}},
		{"hello'boba_SLEEP(", result{5, 2}},
		{"./././././././././././././././././././././././././././././././././", result{132, 4}},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestSearchAttackPts(t *testing.T) {
	testMatch := func(t *testing.T, hay string, exp result) {
		attack, tot := attackAhoSearch(testAC, testPatterns, hay)
		if attack.Type != exp.Ty {
			t.Errorf("type doesn't match, got %d, exp %d", attack.Type, exp.Ty)
		}
		if tot != exp.To {
			t.Errorf("total doesn't match, got %d, exp %d", tot, exp.To)
		}
	}

	data := []struct {
		In  string
		Exp result
	}{
		{"\x99()", result{255, Unk}},
		{"'SLEEP(0)", result{255, SQL}},
		{"socket.hitmycommandcenter", result{255, CMD}},
		{"../etc/passwd", result{255, Dir}},
		{"<script>", result{255, XSS}},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestBotScan(t *testing.T) {
	testMatch := func(t *testing.T, agent string, exp string) {
		got, err := botScan(agent)
		if err != nil {
			t.Error(err)
		}
		if got != exp {
			t.Errorf("bot doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		UA  string
		Exp string
	}{
		{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Googlebot"},
	}

	for _, d := range data {
		testMatch(t, d.UA, d.Exp)
	}
}

func TestBotAhoSearch(t *testing.T) {
	testMatch := func(t *testing.T, agent string, exp string) {
		got, err := botAhoSearch(botAC, agent)
		if err != nil {
			t.Error(err)
		}
		if got != exp {
			t.Errorf("bot doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		UA  string
		Exp string
	}{
		{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Googlebot"},
	}

	for _, d := range data {
		testMatch(t, d.UA, d.Exp)
	}
}

// === benchmarks ===

// BenchmarkGetStringFromSublice-8   	1000000000	         0.3137 ns/op
func BenchmarkGetStringFromSublice(b *testing.B) {
	got := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	nul := func(s string) {}

	for i := 0; i < b.N; i++ {
		bot := got[146:155]
		nul(bot)
	}
}

// BenchmarkGetStringFromSlice-8     	1000000000	         0.6072 ns/op
func BenchmarkGetStringFromSlice(b *testing.B) {
	nul := func(s string) {}

	for i := 0; i < b.N; i++ {
		bot := botPatterns[8]
		nul(bot)
	}
}
