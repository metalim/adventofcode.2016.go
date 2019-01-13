package numbers

import (
	"strconv"
	"strings"
)

// MinMax values of []int.
func MinMax(sn []int) (min, max int) {
	if len(sn) == 0 {
		return
	}
	min = sn[0]
	max = min
	for _, n := range sn[1:] {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
	}
	return
}

// Join numbers into a string.
func Join(sn []int, sep string) string {
	var out strings.Builder
	for _, n := range sn {
		out.WriteString(strconv.Itoa(n))
		out.WriteString(sep)
	}
	return out.String()
}

// Abs value.
func Abs(n int) int {
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
}
