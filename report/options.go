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
		for ip, atkr := range r.Atkr {
			count = append(count, Count{atkr.Cnt, ip})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	r.AtkrCnt = count
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
		for name, bot := range t.Bot {
			count = append(count, Count{bot.Cnt, name})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	t.BotCnt = count
	return nil
}

func atkrMin(t *TextReport, min string) error {
	imin, err := strconv.Atoi(min)
	if err != nil {
		return fmt.Errorf("invalid option value \"%s\"", min)
	}

	filtered := make(map[string]Atkrs)
	for ip, atkr := range t.Atkr {
		if atkr.Cnt >= uint(imin) {
			filtered[ip] = Atkrs{atkr.Cnt}
		} else {
			t.AtkrFiltered++
		}
	}
	t.Atkr = filtered

	if t.AtkrCnt != nil {
		count := []Count{}
		for _, atrk := range t.AtkrCnt {
			if atrk.Num >= uint(imin) {
				count = append(count, atrk)
			}
		}
		t.AtkrCnt = count
	}

	return nil
}

func botMin(t *TextReport, min string) error {
	imin, err := strconv.Atoi(min)
	if err != nil {
		return fmt.Errorf("invalid option value \"%s\"", min)
	}

	filtered := make(map[string]Bots)
	for name, bot := range t.Bot {
		if bot.Cnt >= uint(imin) {
			filtered[name] = Bots{bot.Cnt}
		} else {
			t.BotFiltered++
		}
	}
	t.Bot = filtered

	if t.BotCnt != nil {
		count := []Count{}
		for _, bot := range t.BotCnt {
			if bot.Num >= uint(imin) {
				count = append(count, bot)
			}
		}
		t.BotCnt = count
	}

	return nil
}
