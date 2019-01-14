package source

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func newParser(in input) Parser {
	return Parser{input: in, data: in.String()}
}

// Parser extend input with parsing capabilities.
type Parser struct {
	input
	data interface{}
}

// Words split by all non-word characters.
func (par *Parser) Words() *Parser {
	// log.Printf("Words for %#v\n", par.data)
	reg := regexp.MustCompile("\\w+")
	switch v := par.data.(type) {
	case string:
		par.data = reg.FindAllString(v, -1)

	case []string:
		data := make([][]string, 0, len(v))
		for _, s := range v {
			data = append(data, reg.FindAllString(s, -1))
		}
		par.data = data
	default:
		log.Fatalf("unsupported split level %T", v)
	}
	// log.Println("Words complete", par.data)
	return par
}

// Split by strict separator.
func (par *Parser) Split(sep string) *Parser {
	// log.Printf("Splitting %#v by %#v\n", par.data, sep)
	switch v := par.data.(type) {
	case string:
		par.data = strings.Split(v, sep)
	case []string:
		data := make([][]string, 0, len(v))
		for _, s := range v {
			data = append(data, strings.Split(s, sep))
		}
		par.data = data
	default:
		log.Fatalf("unsupported split level %T", v)
	}
	// log.Println("Split complete", par.data)
	return par
}

// Lines split by \n.
func (par *Parser) Lines() *Parser {
	return par.Split("\n")
}

// Int converts all values to integers.
func (par *Parser) Int() *Parser {
	// log.Printf("Int for %#v\n", par.data)
	switch v := par.data.(type) {
	case string:
		par.data, _ = strconv.Atoi(v)

	case []string:
		data := make([]int, len(v))
		for i, s := range v {
			data[i], _ = strconv.Atoi(s)
		}
		par.data = data
	case [][]string:
		data := make([][]int, 0, len(v))
		for _, ss := range v {
			r := make([]int, len(ss))
			for i, s := range ss {
				r[i], _ = strconv.Atoi(s)
			}
			data = append(data, r)
		}
		par.data = data
	default:
		log.Fatalf("unsupported split level %T", v)
	}
	// log.Println("Int complete", par.data)
	return par
}

// Data returns result.
func (par *Parser) Data() interface{} {
	return par.data
}
