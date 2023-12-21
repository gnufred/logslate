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

func TestAdOptions(t *testing.T) {
	testOptions := func(t *testing.T, d *TextReport, o Opts) {
		err := adOptions(d, o)
		if err != nil {
			t.Error(err)
		}
	}

	rep := New("screen", "text", "ads")
	data := NewText(rep)

	err := adOptions(data, Opts{"banana!": ""})
	test.ErrorExpected(t, err, "invalid option")

	testOptions(t, data, Opts{"order": "name"})
	testOptions(t, data, Opts{"order": "count"})
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
	rep.Attackers["127.0.0.1"] = AttackersRD{10}
	rep.Attackers["127.0.0.2"] = AttackersRD{100}
	data := NewText(rep)

	err := atkrOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.AttackersCount != nil {
		t.Error("attacker count is not nil")
	}

	err = atkrOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.AttackersCount[0].Item != "127.0.0.2" {
		t.Error("attacker order by count data is incorrect")
	}
}

func TestAdOrder(t *testing.T) {
	rep := New("screen", "text", "ads")
	rep.Ads["goog"] = AdsRD{10}
	rep.Ads["face"] = AdsRD{100}
	data := NewText(rep)

	err := adOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.AdsCount != nil {
		t.Error("ad count is not nil")
	}

	err = adOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.AdsCount[0].Item != "face" {
		t.Error("ad ordered by count data is incorrect")
	}
}

func TestBotOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Bots["one"] = BotsRD{10}
	rep.Bots["two"] = BotsRD{100}
	data := NewText(rep)

	err := botOrder(data, "")
	test.ErrorExpected(t, err, "invalid option")
	if data.BotsCount != nil {
		t.Error("bot count is not nil")
	}

	err = botOrder(data, "count")
	if err != nil {
		t.Error(err)
	}
	if data.BotsCount[0].Item != "two" {
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
	rep.Attackers["127.0.0.0"] = AttackersRD{0}
	rep.Attackers["127.0.0.1"] = AttackersRD{1}
	rep.Attackers["127.0.0.2"] = AttackersRD{5}
	rep.Attackers["127.0.0.3"] = AttackersRD{10}
	rep.Attackers["127.0.0.4"] = AttackersRD{20}
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

	testMinimum(t, len(data.Attackers), 2)
	testMinimum(t, len(data.AttackersCount), 2)
	testFiltered(t, data.AttackersFiltered, 3)
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
	rep.Bots["zero"] = BotsRD{1}
	rep.Bots["one"] = BotsRD{10}
	rep.Bots["two"] = BotsRD{20}
	rep.Bots["three"] = BotsRD{50}
	rep.Bots["four"] = BotsRD{100}
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

	testMinimum(t, len(data.Bots), 2)
	testMinimum(t, len(data.BotsCount), 2)
	testFiltered(t, data.BotsFiltered, 3)
}
