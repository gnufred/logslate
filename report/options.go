package report

import (
	"fmt"
	"sort"
	"strconv"
)

type filterData struct {
	List     map[string]ListRD
	Filtered uint
	Cnt      []Count
}

func execOptions(fd filterData, options Opts) (filterData, error) {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			switch val {
			default:
				return fd, fmt.Errorf("invalid option value \"%s\"", val)
			case "count":
				fd = filterOrder(fd)
			case "name":
			}
		case "min":
			min, err := strconv.Atoi(val)
			if err != nil {
				return fd, fmt.Errorf("invalid option value \"%s\"", val)
			}
			fd = filterMin(min, fd)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return fd, err
}

func filterOrder(fd filterData) filterData {
	count := []Count{}
	for name, entry := range fd.List {
		count = append(count, Count{entry.Cnt, name})
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i].Num > count[j].Num
	})
	fd.Cnt = count
	return fd
}

func filterMin(min int, fd filterData) filterData {
	filtered := make(map[string]ListRD, 0)
	var filteredNum uint

	for name, entry := range fd.List {
		if entry.Cnt >= uint(min) {
			filtered[name] = ListRD{entry.Cnt}
		} else {
			filteredNum++
		}
	}

	filteredCount := []Count{}
	for _, entry := range fd.Cnt {
		if entry.Num >= uint(min) {
			filteredCount = append(filteredCount, entry)
		}
	}

	return filterData{filtered, filteredNum, filteredCount}
}
