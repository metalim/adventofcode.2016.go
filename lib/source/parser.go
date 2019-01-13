package source

import (
	"log"
	"regexp"
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

// Data returns result.
func (par *Parser) Data() interface{} {
	return par.data
}
