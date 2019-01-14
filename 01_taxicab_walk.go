package main

import (
	"fmt"
	"metalim/advent/2016/lib/field"
	"metalim/advent/2016/lib/numbers"
	"metalim/advent/2016/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins.Test(1, `R2, L3`, `5`)
	ins.Test(1, `R2, R2, R2`, `2`)
	ins.Test(1, `R5, L5, R5, R3`, `12`)
	ins.Test(2, `R8, R4, R4, R8`, `4`)
	ins.Advent(2016, 1)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		sw := par.Words().Data().([]string)
		fmt.Println(len(sw), Black(sw[0]).Bold())

		if par.Part(1) {
			d := 3
			var x, y int
			for _, w := range sw {
				switch w[0] {
				case 'L':
					d = (d + 3) & 3
				case 'R':
					d = (d + 1) & 3
				}
				n, _ := strconv.Atoi(w[1:])
				x += (1 - d) % 2 * n
				y += (2 - d) % 2 * n
			}
			manh := numbers.Abs(x) + numbers.Abs(y)
			par.SubmitInt(1, manh)
		}

		if par.Part(2) {
			f := field.Map{}
			d := field.Dir40N
			var p field.Pos
			f.Set(p, 1)

		LOOP:
			for _, w := range sw {
				switch w[0] {
				case 'L':
					d = (d + 3) & 3
				case 'R':
					d = (d + 1) & 3
				}
				n, _ := strconv.Atoi(w[1:])
				for i := 0; i < n; i++ {
					p = field.Step4(p, d)
					if f.Get(p) != 0 {
						break LOOP
					}
					f.Set(p, 1)
				}
			}
			manh := field.Manh(p, field.Pos{})
			par.SubmitInt(2, manh)
		}

	}
}
