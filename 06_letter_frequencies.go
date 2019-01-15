package main

import (
	"fmt"
	"metalim/advent/2016/lib/freq"
	"metalim/advent/2016/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func main() {
	var ins source.Inputs

	ins.Test(1|2, test1, `easter`, `advent`)
	ins.Advent(2016, 6)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		sl := par.Lines().Data().([]string)
		fmt.Println(len(sl), Black(sl[0]).Bold())

		if par.Part(1) {
			out := make([]byte, len(sl[0]))
			for i := range sl[0] {
				fs := freq.Map{}
				for _, l := range sl {
					fs.Add(l[i])
				}
				s := fs.Sorted()
				out[i] = s[len(s)-1].V.(byte)
			}
			par.Submit(1, string(out))
		}

		if par.Part(2) {
			out := make([]byte, len(sl[0]))
			for i := range sl[0] {
				fs := freq.Map{}
				for _, l := range sl {
					fs.Add(l[i])
				}
				s := fs.Sorted()
				out[i] = s[0].V.(byte)
			}
			par.Submit(2, string(out))
		}

	}
}
