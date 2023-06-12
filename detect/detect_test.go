package detect

import "testing"

type result struct {
	To uint32
	Ty byte
}

func TestAttacks(t *testing.T) {
	testMatch := func(t *testing.T, hay string, exp result) {
		to, ty := Attacks(hay, "")
		got := result{To: to, Ty: ty}
		if got != exp {
			t.Errorf("query total doesn't match, got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp result
	}{
		{"hello(", result{0, 0}},
		{"print(", result{5, 5}},
		{"\x22id\x22:\x221\x22", result{20, 1}},
		{"path/stuff../../etc/passwd", result{13, 4}},
		{"hello''world", result{8, 1}},
		{"hello'boba_sleep(", result{9, 2}},
		{"hello'boba_SLEEP(", result{9, 2}},
		{"./././././././././././././././././././././././././././././././././", result{132, 4}},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestSearchAttackPts(t *testing.T) {
	testMatch := func(t *testing.T, hay string, exp result) {
		tot, typ := searchAttackPts(testAC, testPts, hay)
		if tot != exp.To {
			t.Errorf("total doesn't match, got %d, exp %d", tot, exp.To)
		}
		if typ != exp.Ty {
			t.Errorf("type doesn't match, got %d, exp %d", typ, exp.Ty)
		}
	}

	data := []struct {
		In  string
		Exp result
	}{
		{"\x99()", result{255, Gen}},
		{"'SLEEP(0)", result{255, SQL}},
		{"socket.hitmycommandcenter", result{255, CMD}},
		{"../etc/passwd", result{255, Dir}},
		{"<script>", result{255, XSS}},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
