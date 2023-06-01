package report

import (
	"strings"
	"testing"
)

func TestCreateTextOutput(t *testing.T) {
	testOut := func(t *testing.T, gs, exp string) {
		if !strings.Contains(gs, exp) {
			t.Errorf("text output didn't contain %s", exp)
		}
	}

	r := New()
	got, err := createTextOutput(r)
	if err != nil {
		t.Error(err)
	}
	gs := got.String()

	data := []string{
		"=== [[ Summary ]] ===",
		"=== [[ Time ]] ===",
		"=== [[ Report ]] ===",
	}

	for _, s := range data {
		testOut(t, gs, s)
	}
}
