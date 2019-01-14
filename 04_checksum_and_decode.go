package main

import (
	"fmt"
	"metalim/advent/2016/lib/source"
	"regexp"
	"sort"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var test1 = `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

func main() {
	var ins source.Inputs

	ins.Test(1, test1, `1514`)
	ins.Test(2, `qzmt-zixmtkozy-ivhz-343[aaaaa]`, `343`) // just to test decrypt
	ins.Advent(2016, 4)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		sw := par.Lines().Data().([]string)
		fmt.Println(len(sw), Black(sw[0]).Bold())

		if par.Part(1) {
			reg := regexp.MustCompile("(.*)-(\\d+)\\[(.*)\\]")
			var sum int
			for _, s := range sw {
				m := reg.FindStringSubmatch(s)
				if getChecksum(m[1]) == m[3] {
					n, _ := strconv.Atoi(m[2])
					sum += n
				}
			}
			par.SubmitInt(1, sum)
		}

		if par.Part(2) {
			reg := regexp.MustCompile("(.*)-(\\d+)\\[(.*)\\]")
			var n int
			for _, s := range sw {
				m := reg.FindStringSubmatch(s)
				n, _ = strconv.Atoi(m[2])
				name := decrypt(m[1], n)
				if strings.Contains(name, "north") {
					break
				}
			}
			par.SubmitInt(2, n)
		}

	}
}

func decrypt(s string, key int) string {
	var out strings.Builder
	for _, r := range s {
		switch r {
		case '-':
			out.WriteRune(' ')
		default:
			out.WriteRune(((r-'a')+rune(key))%('z'-'a'+1) + 'a')
		}
	}
	return out.String()
}

func getChecksum(s string) string {
	f := map[rune]int{}
	for _, c := range s {
		f[c]++
	}
	fs := []freq{}
	for r := rune('a'); r <= 'z'; r++ {
		if f[r] > 0 {
			fs = append(fs, freq{r, f[r]})
		}
	}
	sort.Stable(byFreq(fs))
	rs := [5]rune{}
	for i := 0; i < 5; i++ {
		rs[i] = fs[i].r
	}
	return string(rs[:])
}

type freq struct {
	r rune
	f int
}

type byFreq []freq

func (f byFreq) Len() int {
	return len(f)
}

func (f byFreq) Less(i, j int) bool {
	return f[i].f > f[j].f
}

func (f byFreq) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
