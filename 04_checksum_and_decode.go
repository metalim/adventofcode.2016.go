package main

import (
	"fmt"
	"metalim/advent/2016/lib/freq"
	"metalim/advent/2016/lib/source"
	"regexp"
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
			var sum int
			for _, s := range sw {
				m := reg.FindStringSubmatch(s)
				if getChecksum(strings.Replace(m[1], "-", "", -1)) == m[3] {
					n, _ := strconv.Atoi(m[2])
					sum += n
				}
			}
			par.SubmitInt(1, sum)
		}

		if par.Part(2) {
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

var reg = regexp.MustCompile("(.*)-(\\d+)\\[(.*)\\]")

func getChecksum(s string) string {
	f := freq.RuneMap{}
	for _, c := range s {
		f.Add(c)
	}
	fs := f.Sorted()
	rs := [5]rune{}
	for i := 0; i < 5; i++ {
		rs[i] = fs[i].V
	}
	return string(rs[:])
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

// type freq struct {
// 	r rune
// 	f int
// }

// type byFreq []freq

// func (f byFreq) Len() int {
// 	return len(f)
// }

// func (f byFreq) Less(i, j int) bool {
// 	return f[i].f > f[j].f
// }

// func (f byFreq) Swap(i, j int) {
// 	f[i], f[j] = f[j], f[i]
// }
