package main

import (
	"fmt"
	"metalim/advent/2016/lib/source"
	"regexp"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins.Test(1, `abba[mnop]qrst`, `1`)
	ins.Test(1, `abcd[bddb]xyyx`, `0`)
	ins.Test(1, `aaaa[qwer]tyui`, `0`)
	ins.Test(1, `ioxxoj[asdfgh]zxcvbn`, `1`)
	ins.Test(1, `ioxxoj[asdfgh]zxcvbnabcd[bddb]xyyx`, `0`)
	ins.Test(2, `aba[bab]xyz`, `1`)
	ins.Test(2, `xyx[xyx]xyx`, `0`)
	ins.Test(2, `aaa[kek]eke`, `1`)
	ins.Test(2, `zazbz[bzb]cdb`, `1`)
	ins.Advent(2016, 7)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		sl := par.Lines().Data().([]string)
		fmt.Println(len(sl), Black(sl[0]).Bold())

		if par.Part(1) {
			n := 0
		NEXT:
			for _, l := range sl {
				sm := reg.FindAllStringSubmatch(l, -1)
				for _, m := range sm {
					if hasABBA(m[3]) {
						continue NEXT
					}
				}
				for _, m := range sm {
					if hasABBA(m[1]) {
						n++
						continue NEXT
					}
				}
			}
			par.SubmitInt(1, n)
		}

		if par.Part(2) {
			n := 0
		NEXT2:
			for _, l := range sl {
				sm := reg.FindAllStringSubmatch(l, -1)

				for _, m := range sm {
					s := m[1]
					for i := 0; i < len(s)-2; i++ {
						if s[i] != s[i+1] && s[i] == s[i+2] {
							aba := s[i : i+3]
							bab := string([]byte{aba[1], aba[0], aba[1]})
							for _, m2 := range sm {
								if strings.Contains(m2[3], bab) {
									n++
									continue NEXT2
								}
							}
						}
					}
				}
			}
			par.SubmitInt(2, n)
		}

	}
}

var reg = regexp.MustCompile("(\\w+)(\\[([^\\]]+)\\])?") // all out[in] and out

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}
