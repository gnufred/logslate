package report

import (
	"testing"

	"gitlab.com/gnufred/logslate/test"
)

func TestAtkrOptions(t *testing.T) {
	testOptions := func(t *testing.T, d *TextReport, o Opts) {
		err := atkrOptions(d, o)
		if err != nil {
			t.Error(err)
		}
	}

	rep := New("screen", "text", "attackers")
	data := NewText(rep)

	err := atkrOptions(data, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")

	testOptions(t, data, Opts{"order": "name"})
	testOptions(t, data, Opts{"order": "count"})
	testOptions(t, data, Opts{"min": "10"})
}

func TestBotOptions(t *testing.T) {
	testOptions := func(t *testing.T, d *TextReport, o Opts) {
		err := botOptions(d, o)
		if err != nil {
			t.Error(err)
		}
	}

	rep := New("screen", "text", "bots")
	data := NewText(rep)

	err := botOptions(data, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")

	testOptions(t, data, Opts{"order": "name"})
	testOptions(t, data, Opts{"order": "count"})
	testOptions(t, data, Opts{"min": "10"})
}

func TestAtkrOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Atkr["127.0.0.1"] = Atkrs{10}
	rep.Atkr["127.0.0.2"] = Atkrs{100}
	data := NewText(rep)

	err := atkrOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.AtkrCnt != nil {
		t.Error("attacker count is not nil")
	}

	err = atkrOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.AtkrCnt[0].Item != "127.0.0.2" {
		t.Error("attacker order by count data is incorrect")
	}
}
func TestBotOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Bot["one"] = Bots{10}
	rep.Bot["two"] = Bots{100}
	data := NewText(rep)

	err := botOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.BotCnt != nil {
		t.Error("bot count is not nil")
	}

	err = botOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.BotCnt[0].Item != "two" {
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

	rep := New("screen", "text", "attackers")
	rep.Atkr["127.0.0.0"] = Atkrs{0}
	rep.Atkr["127.0.0.1"] = Atkrs{1}
	rep.Atkr["127.0.0.2"] = Atkrs{5}
	rep.Atkr["127.0.0.3"] = Atkrs{10}
	rep.Atkr["127.0.0.4"] = Atkrs{20}
	data := NewText(rep)

	err := atkrMin(data, "")
	test.ErrorExpected(t, err, "invalid option")

	// Order the current count, which atrkMin needs to update.
	// Ignoring error because we test them in it's own test.
	atkrOrder(data, "count")
	// Then, filter by minimum.
	err = atkrMin(data, "10")

	if err != nil {
		t.Error(err)
	}

	testMinimum(t, len(data.Atkr), 2)
	testMinimum(t, len(data.AtkrCnt), 2)
	testFiltered(t, data.AtkrFiltered, 3)
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

	rep := New("screen", "text", "bots")
	rep.Bot["zero"] = Bots{1}
	rep.Bot["one"] = Bots{10}
	rep.Bot["two"] = Bots{20}
	rep.Bot["three"] = Bots{50}
	rep.Bot["four"] = Bots{100}
	data := NewText(rep)

	err := botMin(data, "")
	test.ErrorExpected(t, err, "invalid option")

	// Order the current count, which atrkMin needs to update.
	// Ignoring error because we test them in it's own test.
	botOrder(data, "count")
	// Then, filter by minimum.
	err = botMin(data, "30")

	if err != nil {
		t.Error(err)
	}

	testMinimum(t, len(data.Bot), 2)
	testMinimum(t, len(data.BotCnt), 2)
	testFiltered(t, data.BotFiltered, 3)
}
