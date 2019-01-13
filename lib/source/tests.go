package source

import (
	"fmt"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

// Parts specifies which part should be submitted/tested.
type Parts uint

// We can have part 1, 2 or 1+2.
const (
	Part1 Parts = 1 << iota
	Part2
	Part12 Parts = 3
)

// Test adds test input (with expected outputs) to the queue.
func (ins *Inputs) Test(part Parts, in string, ex ...string) {
	*ins = append(*ins, newTest(part, in, ex))
}

////////////////////////////////////////////////////////////////////////
// Implementation

var tests int

func newTest(parts Parts, in string, ex []string) input {
	tests++
	test := &test{inputBase{name: "test" + strconv.Itoa(tests)}, parts, in, ex}
	test.input = test

	last := 1
	for parts = parts >> 1; parts > 0; parts = parts >> 1 {
		last++
	}

	if len(ex) < last {
		test.ex = append(make([]string, last-len(ex)), ex...)
	}

	return test
}

type test struct {
	inputBase
	parts Parts // 1, 2, 1+2=3
	in    string
	ex    []string
}

func (t *test) Part(part uint) bool {
	return t.canProcess(part) && t.parts&(1<<(part-1)) != 0
}

func (t *test) String() string {
	return t.in
}

func (t *test) Submit(part uint, val string) bool {
	if !t.Part(part) {
		panic("part should not be submitted")
	}
	if val != t.ex[part-1] {
		fmt.Printf("part%d: %s %s\n", part, Cyan(val), Red("✗ expected "+t.ex[part-1]))
		return false
	}
	t.valid[part-1] = true
	fmt.Printf("part%d: %s %s\n", part, Cyan(val), Green("✓"))
	return true
}
