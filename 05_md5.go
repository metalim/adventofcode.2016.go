package main

import (
	"crypto/md5"
	"fmt"
	"metalim/advent/2016/lib/source"

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
			key := make([]byte, 0, 20)
			key = append(key, (id + "0")...)
			first := len(id)

			var pass [8]byte
			var filled int
			for filled < 8 {
				hash := md5.Sum(key)
				if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
					pass[filled] = hex[hash[2]]
					filled++
				}
				key = inc(key, first)
			}
			par.Submit(1, string(pass[:]))
		}

		if par.Part(2) {
			key := make([]byte, 0, 20)
			key = append(key, (id + "0")...)
			first := len(id)

			var pass [8]byte
			var filled int
			for filled < 8 {
				hash := md5.Sum(key)
				if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
					index := hash[2]
					if index < 8 && pass[index] == 0 {
						pass[index] = hex[hash[3]>>4]
						filled++
					}
				}
				key = inc(key, first)
			}
			par.Submit(2, string(pass[:]))
		}

	}
}

const hex = "0123456789abcdef"

// Just for practice.
// Manually increment decimal string, up to specified position.
// Can be replaced by key = id + strconv.Itoa(i), which is slower.
func inc(key []byte, first int) []byte {
	for i := len(key) - 1; ; i-- {
		if key[i] != '9' {
			key[i]++
			break
		}
		key[i] = '0'
		if i == first { // all 0, overflow from all 9, need 1 more digit.
			key = append(key, '0')
			key[first] = '1'
			break
		}
	}
	return key
}
