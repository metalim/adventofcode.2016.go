package main

import (
	"fmt"
	"metalim/advent/2016/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `zzz`

func main() {
	var ins source.Inputs

	ins.Advent(2016, 3)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		ssi := par.Lines().Words().Int().Data().([][]int)
		fmt.Println(len(ssi), Black(ssi[0]).Bold())

		if par.Part(1) {
			var num int
			for _, t := range ssi {
				if t[0] < t[1]+t[2] && t[1] < t[2]+t[0] && t[2] < t[0]+t[1] {
					num++
				}
			}
			par.SubmitInt(1, num)
		}

		if par.Part(2) {
			var num int
			for i := range ssi {
				it, d := i/3, i%3
				a := ssi[it*3][d]
				b := ssi[it*3+1][d]
				c := ssi[it*3+2][d]
				if a < b+c && b < c+a && c < a+b {
					num++
				}
			}
			par.SubmitInt(2, num)
		}

	}
}
