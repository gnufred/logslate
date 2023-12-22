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
	rep.Attackers["127.0.0.1"] = ListRD{10}
	rep.Attackers["127.0.0.2"] = ListRD{100}
	data := NewText(rep)

	err := atkrOptions(data, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if data.AttackersCount != nil {
		t.Error("attacker count is not nil")
	}

	err = atkrOptions(data, Opts{"order": "count"})
	if err != nil {
		t.Error(err)
	}
	if data.AttackersCount[0].Item != "127.0.0.2" {
		t.Error("attacker order by count data is incorrect")
	}
}

func TestAdOrder(t *testing.T) {
	rep := New("screen", "text", "ads")
	rep.Ads["goog"] = ListRD{10}
	rep.Ads["face"] = ListRD{100}
	data := NewText(rep)

	err := adOptions(data, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if data.AdsCount != nil {
		t.Error("ad count is not nil")
	}

	err = adOptions(data, Opts{"order": "count"})
	if err != nil {
		t.Error(err)
	}
	if data.AdsCount[0].Item != "face" {
		t.Error("ad ordered by count data is incorrect")
	}
}

func TestBotOrder(t *testing.T) {
	rep := New("screen", "text", "attackers")
	rep.Bots["one"] = ListRD{10}
	rep.Bots["two"] = ListRD{100}
	data := NewText(rep)

	err := botOptions(data, Opts{"order": "invalid"})
	test.ErrorExpected(t, err, "invalid option")
	if data.BotsCount != nil {
		t.Error("bot count is not nil")
	}

	err = botOptions(data, Opts{"order": "count"})
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
	rep.Attackers["127.0.0.0"] = ListRD{0}
	rep.Attackers["127.0.0.1"] = ListRD{1}
	rep.Attackers["127.0.0.2"] = ListRD{5}
	rep.Attackers["127.0.0.3"] = ListRD{10}
	rep.Attackers["127.0.0.4"] = ListRD{20}
	data := NewText(rep)

	// Check if a non integer string make it fail
	err := atkrOptions(data, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	// Order the current count, which atrkMin needs to update.
	// Ignoring error because we test them in it's own test.
	data.AttackersCount = filterOrder(data.Attackers)
	// Then, filter by minimum.
	min := 10
	data.Attackers, data.AttackersFiltered, data.AttackersCount = filterMin(
		min, data.Attackers, data.AttackersCount)

	testMinimum(t, len(data.Attackers), 2)
	testMinimum(t, len(data.AttackersCount), 2)
	testFiltered(t, data.AttackersFiltered, 3)
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

	rep := New("screen", "text", "ads")
	rep.Ads["goog"] = ListRD{1}
	rep.Ads["face"] = ListRD{10}
	rep.Ads["pin"] = ListRD{20}
	rep.Ads["x"] = ListRD{50}
	rep.Ads["y"] = ListRD{100}
	data := NewText(rep)

	// Check if a non integer string make it fail
	err := adOptions(data, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	// Order the current count, which atrkMin needs to update.
	// Ignoring error because we test them in it's own test.
	adOptions(data, Opts{"order": "count"})
	// Then, filter by minimum.
	min := 30
	data.Ads, data.AdsFiltered, data.AdsCount = filterMin(
		min, data.Ads, data.AdsCount)

	testMinimum(t, len(data.Ads), 2)
	testMinimum(t, len(data.AdsCount), 2)
	testFiltered(t, data.AdsFiltered, 3)
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
	rep.Bots["zero"] = ListRD{1}
	rep.Bots["one"] = ListRD{10}
	rep.Bots["two"] = ListRD{20}
	rep.Bots["three"] = ListRD{50}
	rep.Bots["four"] = ListRD{100}
	data := NewText(rep)

	// Check if a non integer string make it fail
	err := botOptions(data, Opts{"min": "invalid_value"})
	test.ErrorExpected(t, err, "invalid option")

	// Order the current count, which atrkMin needs to update.
	// Ignoring error because we test them in it's own test.
	botOptions(data, Opts{"order": "count"})
	// Then, filter by minimum.
	min := 30
	data.Bots, data.BotsFiltered, data.BotsCount = filterMin(
		min, data.Bots, data.BotsCount)

	testMinimum(t, len(data.Bots), 2)
	testMinimum(t, len(data.BotsCount), 2)
	testFiltered(t, data.BotsFiltered, 3)
}
