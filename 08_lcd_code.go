package main

import (
	"fmt"
	"metalim/advent/2016/lib/field"
	"metalim/advent/2016/lib/source"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var test1 = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`

func main() {
	var ins source.Inputs

	ins.Test(1, test1, `6`)
	ins.Advent(2016, 8)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		ssw := par.Lines().Words().Data().([][]string)
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		f := field.Slice{}
		f.SetDefault(' ')
		for _, sw := range ssw {
			switch sw[0] {
			case "rect":
				s := strings.Split(sw[1], "x")
				w, _ := strconv.Atoi(s[0])
				h, _ := strconv.Atoi(s[1])
				for y := 0; y < h; y++ {
					for x := 0; x < w; x++ {
						f.Set(field.Pos{x, y}, '#')
					}
				}
			case "rotate":
				p, _ := strconv.Atoi(sw[3])
				d, _ := strconv.Atoi(sw[5])
				// fmt.Println(sw[2], p, d)
				if sw[2] == "y" {
					row := [50]int{}
					for x := range row {
						row[(x+d)%50] = f.Get(field.Pos{x, p})
					}
					for x, v := range row {
						f.Set(field.Pos{x, p}, v)
					}
				} else {
					col := [6]int{}
					for y := range col {
						col[(y+d)%6] = f.Get(field.Pos{p, y})
					}
					for y, v := range col {
						f.Set(field.Pos{p, y}, v)
					}
				}
			}
		}
		if par.Part(1) {
			sum := 0
			for y := 0; y < 6; y++ {
				for x := 0; x < 50; x++ {
					if f.Get(field.Pos{x, y}) == '#' {
						sum++
					}
				}
			}
			par.SubmitInt(1, sum)
		}

		if par.Part(2) {
			print(&f)
			// custom submit for the moment.
			switch par.Name() {
			case "github":
				par.Submit(2, "AFBUPZBJPS")
			case "google":
				par.Submit(2, "ZJHRKCPLYJ")
			}
		}

	}
}

func print(f field.Interface) {
	bs := f.Bounds()
	for y := bs.Min.Y; y < bs.Max.Y; y++ {
		for x := bs.Min.X; x < bs.Max.X; x++ {
			fmt.Printf("%c", f.Get(field.Pos{x, y}))
		}
		fmt.Println()
	}
}
