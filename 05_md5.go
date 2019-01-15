package main

import (
	"crypto/md5"
	"fmt"
	"metalim/advent/2016/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `abc`

func main() {
	var ins source.Inputs

	ins.Test(1|2, test1, `18f47a30`, `05ace8e3`)
	ins.Advent(2016, 5)

	for par := range ins.Iterate() {
		fmt.Println(Brown("\n" + par.Name()).Bold())
		id := par.Data().(string)
		fmt.Println(Black(id).Bold())

		if par.Part(1) {
			var pass [8]byte
			var filled int
			for i := 0; filled < 8; i++ {
				data := id + strconv.Itoa(i)
				hash := md5.Sum([]byte(data))
				if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
					pass[filled] = hex[hash[2]]
					filled++
				}
			}
			par.Submit(1, string(pass[:]))
		}

		if par.Part(2) {
			var pass [8]byte
			var filled int
			for i := 0; filled < 8; i++ {
				data := id + strconv.Itoa(i)
				hash := md5.Sum([]byte(data))
				if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
					index := hash[2]
					if index < 8 && pass[index] == 0 {
						pass[index] = hex[hash[3]>>4]
						filled++
					}
				}
			}
			par.Submit(2, string(pass[:]))
		}

	}
}

const hex = "0123456789abcdef"
