package main

import (
	"fmt"
	"math"
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
				max := 0
				fs := map[byte]int{}
				for _, l := range sl {
					b := l[i]
					fs[b]++
					if max < fs[b] {
						max = fs[b]
						out[i] = b
					}
				}
			}
			par.Submit(1, string(out))
		}

		if par.Part(2) {
			out := make([]byte, len(sl[0]))
			for i := range sl[0] {
				fs := map[byte]int{}
				for _, l := range sl {
					b := l[i]
					fs[b]++
				}
				min := math.MaxInt64
				for b, f := range fs {
					if min > f {
						min = f
						out[i] = b
					}
				}
			}
			par.Submit(2, string(out))
		}

	}
}
