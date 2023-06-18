package report

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Prct output a human readable percentage.
func (r *Report) Prct(divident, divisor uint) float64 {
	if divident == 0 || divisor == 0 {
		return 0
	}
	prc := float64(divident) / float64(divisor) * 100
	rounded := math.Round(prc*10) / 10
	return rounded
}

// HumanTime format a time.Time to a human readlable string.
func (r *Report) HumanTime(tm time.Time) string {
	layout := "2006-01-02 15:04:05 MST"
	return tm.Format(layout)
}

// HumanDuration format a time.Duration to a human readable string.
func (r *Report) HumanDuration(d time.Duration) string {
	ds := fmt.Sprint(d)
	if strings.HasSuffix(ds, "ms") {
		if i := strings.Index(ds, "."); i > -1 {
			return ds[:i] + "ms"
		}
	}
	if strings.HasSuffix(ds, "s") {
		if strings.Contains(ds, "m") {
			if i := strings.Index(ds, "."); i > -1 {
				return ds[:i] + "s"
			}
		} else {
			if i := strings.Index(ds, "."); i > -1 {
				return ds[:i+2] + "s"
			}
		}
	}
	return ds
}

// BytesToMB convert a Bytes number into MB.
func (r *Report) BytesToMB(bytes uint) float64 {
	bts := float64(bytes)
	ratioMB := float64(1048576)
	MB := bts / ratioMB
	roundedMB := math.Ceil(MB*100) / 100
	return roundedMB
}

// AddCommas add commas to number. See this post where I took the source code:
// https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
func (r *Report) AddCommas(num uint) string {
	n := int64(num)
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
