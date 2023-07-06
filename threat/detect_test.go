package threat

import (
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

type result struct {
	To uint32
	Ty AttackType
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
				SQL,
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
		{"replace(", result{5, XSS}},
		{"\x22id\x22:\x221\x22", result{12, Unk}},
		{"path/stuff../../etc/passwd", result{9, Dir}},
		{"hello'boba_sleep(", result{5, SQL}},
		{"hello'boba_SLEEP(", result{5, SQL}},
		{"./././././././././././././././././././././././././././././././././", result{66, Dir}},
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
	validBotScans(t)
	moreThanOneBotFound(t)
}

func validBotScans(t *testing.T) {
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
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/114.0.5735.99 Mobile/15E148 Safari/604.1", ""},
		{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Googlebot"},
		{"FeedBurner/1.0 (http://www.FeedBurner.com)", "FeedBurner"},
		{"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Chrome/103.0.5060.134 Safari/537.36", "bingbot"},
	}

	for _, d := range data {
		testMatch(t, d.UA, d.Exp)
	}
}

func moreThanOneBotFound(t *testing.T) {
	testMatch := func(t *testing.T, agent string, exp string) {
		got, err := botScan(agent)
		test.ErrorExpected(t, err, "more than one")
		if got != exp {
			t.Errorf("bot doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		UA  string
		Exp string
	}{
		{"Bingbot/2.0; http://Bingbot.com", "Bingbot"},
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
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/114.0.5735.99 Mobile/15E148 Safari/604.1", ""},
		{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Googlebot"},
		{"FeedBurner/1.0 (http://www.FeedBurner.com)", "FeedBurner/"},
	}

	for _, d := range data {
		testMatch(t, d.UA, d.Exp)
	}
}
