package main

import (
	"fmt"
	"metalim/advent/2016/lib/field"
	"metalim/advent/2016/lib/source"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var test1 = `ULL
RRDDD
LURDL
UUUUD`

func main() {
	var ins source.Inputs

	ins.Test(1|2, test1, `1985`, `5DB3`)
	ins.Advent(2016, 2)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		sl := par.Lines().Data().([]string)
		fmt.Println(len(sl), Black(sl[0]).Bold())

		if par.Part(1) {
			f := &field.Map{}
			f.SetDefault(' ')
			field.FillFromString(f, field.Pos{}, "123\n456\n789")
			p := field.Pos{1, 1}
			par.Submit(1, getCode(f, p, sl))
		}

		if par.Part(2) {
			f := &field.Slice{}
			f.SetDefault(' ')
			field.FillFromString(f, field.Pos{}, "  1\n 234\n56789\n ABC\n  D")
			p := field.Pos{0, 2}
			par.Submit(2, getCode(f, p, sl))
		}

	}
}

func getCode(f field.Interface, p field.Pos, sl []string) string {
	var out strings.Builder
	var p2 field.Pos
	for _, l := range sl {
		for _, d := range l {
			switch d {
			case 'U':
				p2 = p.Add(field.Pos{0, -1})
			case 'D':
				p2 = p.Add(field.Pos{0, 1})
			case 'L':
				p2 = p.Add(field.Pos{-1, 0})
			case 'R':
				p2 = p.Add(field.Pos{1, 0})
			}
			if f.Get(p2) != ' ' {
				p = p2
			}
		}
		out.WriteRune(rune(f.Get(p)))
	}
	return out.String()
}
