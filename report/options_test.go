package report

import (
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

func TestExecOptionsErrors(t *testing.T) {
	testOptions := func(t *testing.T, fd filterData, o Opts) {
		_, err := execOptions(fd, o)
		if err != nil {
			t.Error(err)
		}
	}

	err := error(nil)
	rep := New("screen", "text", "ads")
	data := NewText(rep)
	fd := filterData{List: data.Ads}

	fd, err = execOptions(fd, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")

	testOptions(t, fd, Opts{"order": "name"})
	testOptions(t, fd, Opts{"order": "count"})
	testOptions(t, fd, Opts{"min": "10"})
}

func TestAtkrFilterOrder(t *testing.T) {
	err := error(nil)
	rep := New("screen", "text", "attackers")
	rep.Attackers["127.0.0.1"] = ListRD{10}
	rep.Attackers["127.0.0.2"] = ListRD{100}
	data := NewText(rep)
	fd := filterData{List: data.Attackers}

	fd, err = execOptions(fd, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if fd.Cnt != nil {
		t.Error("attacker count is not nil")
	}

	fd, err = execOptions(fd, Opts{"order": "count"})
	if err != nil {
		t.Error(err)
	}
	if fd.Cnt[0].Item != "127.0.0.2" {
		t.Error("attacker order by count data is incorrect")
	}
}

func TestAdOrder(t *testing.T) {
	err := error(nil)
	rep := New("screen", "text", "ads")
	rep.Ads["goog"] = ListRD{10}
	rep.Ads["face"] = ListRD{100}
	data := NewText(rep)
	fd := filterData{List: data.Ads}

	fd, err = execOptions(fd, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if fd.Cnt != nil {
		t.Error("ad count is not nil")
	}

	fd, err = execOptions(fd, Opts{"order": "count"})
	if err != nil {
		t.Error(err)
	}
	if fd.Cnt[0].Item != "face" {
		t.Error("ad ordered by count data is incorrect")
	}
}

func TestBotOrder(t *testing.T) {
	err := error(nil)
	rep := New("screen", "text", "bots")
	rep.Bots["one"] = ListRD{10}
	rep.Bots["two"] = ListRD{100}
	data := NewText(rep)
	fd := filterData{List: data.Bots}

	fd, err = execOptions(fd, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if fd.Cnt != nil {
		t.Error("bot count is not nil")
	}

	fd, err = execOptions(fd, Opts{"order": "count"})
	if err != nil {
		t.Error(err)
	}
	if fd.Cnt[0].Item != "two" {
		t.Error("bot ordered by count data is incorrect")
	}
}

func TestAtkrMin(t *testing.T) {
	testMinimum := func(t *testing.T, got, exp int) {
		if got != exp {
			t.Errorf(
				"attacker minimum is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}
	testFiltered := func(t *testing.T, got, exp uint) {
		if got != exp {
			t.Errorf(
				"attackers filtered is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}

	err := error(nil)
	rep := New("screen", "text", "attackers")
	rep.Attackers["127.0.0.0"] = ListRD{0}
	rep.Attackers["127.0.0.1"] = ListRD{1}
	rep.Attackers["127.0.0.2"] = ListRD{5}
	rep.Attackers["127.0.0.3"] = ListRD{10}
	rep.Attackers["127.0.0.4"] = ListRD{20}
	data := NewText(rep)
	fd := filterData{List: data.Attackers}

	// Check if a non integer string make it fail
	fd, err = execOptions(fd, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	min := 10
	fd = filterOrder(fd)
	fd = filterMin(min, fd)

	testMinimum(t, len(fd.List), 2)
	testMinimum(t, len(fd.Cnt), 2)
	testFiltered(t, fd.Filtered, 3)
}

func TestAdMin(t *testing.T) {
	testMinimum := func(t *testing.T, got, exp int) {
		if got != exp {
			t.Errorf(
				"ad minimum is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}
	testFiltered := func(t *testing.T, got, exp uint) {
		if got != exp {
			t.Errorf(
				"ad filtered is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}

	err := error(nil)
	rep := New("screen", "text", "ads")
	rep.Ads["goog"] = ListRD{1}
	rep.Ads["face"] = ListRD{10}
	rep.Ads["pin"] = ListRD{20}
	rep.Ads["x"] = ListRD{50}
	rep.Ads["y"] = ListRD{100}
	data := NewText(rep)
	fd := filterData{List: data.Ads}

	// Check if a non integer string make it fail
	fd, err = execOptions(fd, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	min := 30
	fd = filterOrder(fd)
	fd = filterMin(min, fd)

	testMinimum(t, len(fd.List), 2)
	testMinimum(t, len(fd.Cnt), 2)
	testFiltered(t, fd.Filtered, 3)
}

func TestBotMin(t *testing.T) {
	testMinimum := func(t *testing.T, got, exp int) {
		if got != exp {
			t.Errorf(
				"bot minimum is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}
	testFiltered := func(t *testing.T, got, exp uint) {
		if got != exp {
			t.Errorf(
				"bot filtered is not the expected length, got %d, exp %d",
				got, exp,
			)
		}
	}

	err := error(nil)
	rep := New("screen", "text", "bots")
	rep.Bots["zero"] = ListRD{1}
	rep.Bots["one"] = ListRD{10}
	rep.Bots["two"] = ListRD{20}
	rep.Bots["three"] = ListRD{50}
	rep.Bots["four"] = ListRD{100}
	data := NewText(rep)
	fd := filterData{List: data.Bots}

	// Check if a non integer string make it fail
	fd, err = execOptions(fd, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	min := 30
	fd = filterOrder(fd)
	fd = filterMin(min, fd)

	testMinimum(t, len(fd.List), 2)
	testMinimum(t, len(fd.Cnt), 2)
	testFiltered(t, fd.Filtered, 3)
}
