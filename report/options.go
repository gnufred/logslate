package report

import (
	"fmt"
	"sort"
	"strconv"
)

func atkrOptions(t *TextReport, options Opts) error {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			err = atkrOrder(t, val)
		case "min":
			err = atkrMin(t, val)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func adOptions(t *TextReport, options Opts) error {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			err = adOrder(t, val)
		case "min":
			err = adMin(t, val)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func botOptions(t *TextReport, options Opts) error {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			err = botOrder(t, val)
		case "min":
			err = botMin(t, val)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func atkrOrder(r *TextReport, order string) error {
	switch order {
	default:
		return fmt.Errorf("invalid option value \"%s\"", order)
	case "count":
	case "name":
	}

	count := []Count{}
	if order == "count" {
		for ip, atkr := range r.Attackers {
			count = append(count, Count{atkr.Cnt, ip})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	r.AttackersCount = count
	return nil
}

func adOrder(t *TextReport, order string) error {
	switch order {
	default:
		return fmt.Errorf("invalid option value \"%s\"", order)
	case "count":
	case "name":
	}

	count := []Count{}
	if order == "count" {
		for name, ad := range t.Ads {
			count = append(count, Count{ad.Cnt, name})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	t.AdsCount = count
	return nil
}

func botOrder(t *TextReport, order string) error {
	switch order {
	default:
		return fmt.Errorf("invalid option value \"%s\"", order)
	case "count":
	case "name":
	}

	count := []Count{}
	if order == "count" {
		for name, bot := range t.Bots {
			count = append(count, Count{bot.Cnt, name})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	t.BotsCount = count
	return nil
}

func atkrMin(t *TextReport, min string) error {
	imin, err := strconv.Atoi(min)
	if err != nil {
		return fmt.Errorf("invalid option value \"%s\"", min)
	}

	filtered := make(map[string]AttackersRD)
	for ip, atkr := range t.Attackers {
		if atkr.Cnt >= uint(imin) {
			filtered[ip] = AttackersRD{atkr.Cnt}
		} else {
			t.AttackersFiltered++
		}
	}
	t.Attackers = filtered

	if t.AttackersCount != nil {
		count := []Count{}
		for _, atrk := range t.AttackersCount {
			if atrk.Num >= uint(imin) {
				count = append(count, atrk)
			}
		}
		t.AttackersCount = count
	}

	return nil
}

func adMin(t *TextReport, min string) error {
	imin, err := strconv.Atoi(min)
	if err != nil {
		return fmt.Errorf("invalid option value \"%s\"", min)
	}

	filtered := make(map[string]AdsRD)
	for name, ad := range t.Ads {
		if ad.Cnt >= uint(imin) {
			filtered[name] = AdsRD{ad.Cnt}
		} else {
			t.AdsFiltered++
		}
	}
	t.Ads = filtered

	if t.AdsCount != nil {
		count := []Count{}
		for _, ad := range t.AdsCount {
			if ad.Num >= uint(imin) {
				count = append(count, ad)
			}
		}
		t.AdsCount = count
	}

	return nil
}

func botMin(t *TextReport, min string) error {
	imin, err := strconv.Atoi(min)
	if err != nil {
		return fmt.Errorf("invalid option value \"%s\"", min)
	}

	filtered := make(map[string]BotsRD)
	for name, bot := range t.Bots {
		if bot.Cnt >= uint(imin) {
			filtered[name] = BotsRD{bot.Cnt}
		} else {
			t.BotsFiltered++
		}
	}
	t.Bots = filtered

	if t.BotsCount != nil {
		count := []Count{}
		for _, bot := range t.BotsCount {
			if bot.Num >= uint(imin) {
				count = append(count, bot)
			}
		}
		t.BotsCount = count
	}

	return nil
}
