package main

import (
	"fmt"
	"metalim/advent/2016/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `zzz`

func main() {
	var ins source.Inputs

	ins.Test(1|2, test1, `1`, `2`)
	ins.Advent(2016, xxx)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		ssw := par.Lines().Words().Data().([][]string)
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		if par.Part(1) {
			par.DrySubmitInt(1, 1)
		}

		// if par.Part(2) {
		// 	par.DrySubmitInt(2, 2)
		// }

	}
}
