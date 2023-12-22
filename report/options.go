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
			switch val {
			default:
				return fmt.Errorf("invalid option value \"%s\"", val)
			case "count":
				t.AttackersCount = filterOrder(t.Attackers)
			case "name":
			}
		case "min":
			min, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid option value \"%s\"", val)
			}
			t.Attackers, t.AttackersFiltered, t.AttackersCount = filterMin(
				min, t.Attackers, t.AttackersCount)
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
			switch val {
			default:
				return fmt.Errorf("invalid option value \"%s\"", val)
			case "count":
				t.AdsCount = filterOrder(t.Ads)
			case "name":
			}
		case "min":
			min, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid option value \"%s\"", val)
			}
			t.Ads, t.AdsFiltered, t.AdsCount = filterMin(
				min, t.Ads, t.AdsCount)
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
			switch val {
			default:
				return fmt.Errorf("invalid option value \"%s\"", val)
			case "count":
				t.BotsCount = filterOrder(t.Bots)
			case "name":
			}
		case "min":
			min, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid option value \"%s\"", val)
			}
			t.Bots, t.BotsFiltered, t.BotsCount = filterMin(
				min, t.Bots, t.BotsCount)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func filterOrder(report map[string]ListRD) []Count {
	count := []Count{}
	for name, entry := range report {
		count = append(count, Count{entry.Cnt, name})
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i].Num > count[j].Num
	})
	return count
}

func filterMin(min int, report map[string]ListRD, reportCount []Count) (map[string]ListRD, uint, []Count) {
	filtered := make(map[string]ListRD)
	var filteredNum uint

	for name, entry := range report {
		if entry.Cnt >= uint(min) {
			filtered[name] = ListRD{entry.Cnt}
		} else {
			filteredNum++
		}
	}

	filteredCount := []Count{}
	for _, entry := range reportCount {
		if entry.Num >= uint(min) {
			filteredCount = append(filteredCount, entry)
		}
	}

	return filtered, filteredNum, filteredCount
}
