package report

import (
	"reflect"
	"testing"
)

func TestParseReportLayout(t *testing.T) {
	testMatch := func(t *testing.T, in string, exp []section) {
		got := parseLayout(in)
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("custom config doesn't match, got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp []section
	}{
		{"section0,section1:opt1_val1,section2:opt1_val1:opt2_val2",
			[]section{
				{"section0", Opts{}},
				{"section1", Opts{
					"opt1": "val1",
				}},
				{"section2", Opts{
					"opt1": "val1",
					"opt2": "val2",
				}},
			},
		},
		{"parsed",
			[]section{
				{"parsed", Opts{}},
			},
		},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestParseSection(t *testing.T) {
	testMatch := func(t *testing.T, in string, exp section) {
		got := parseSection(in)
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("options don't match, got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp section
	}{
		{"section:opt1_val1:opt2_val2",
			section{"section",
				Opts{"opt1": "val1", "opt2": "val2"}},
		},
		{"parsed",
			section{"parsed",
				Opts{}},
		},
		{"attackers:order_name",
			section{"attackers",
				Opts{"order": "name"}},
		},
		{"bots:order_count",
			section{"bots",
				Opts{"order": "count"}},
		},
		{":",
			section{"",
				Opts{"": ""}},
		},
		{"",
			section{"",
				Opts{}},
		},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestParseKeyValue(t *testing.T) {
	testMatch := func(t *testing.T, in string, exp []string) {
		k, v := parseKeyValue(in)
		got := []string{k, v}
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("key value doesn't match, got %v, exp %v", got, exp)
		}
	}

	data := []struct {
		In  string
		Exp []string
	}{
		{"key", []string{"key", ""}},
		{"key_value", []string{"key", "value"}},
		{"key_val_ue", []string{"key", "val_ue"}},
		{"order_name", []string{"order", "name"}},
		{"order_count", []string{"order", "count"}},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
