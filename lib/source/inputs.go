package source

import (
	"fmt"
	"strconv"
)

// exported via embedding in Parser.
type input interface {
	Part(uint) bool // is it Part(N)?
	Name() string   // test1, test2, github.
	String() string // input as a string.

	Submit(uint, string) bool    // Submit an answer.
	SubmitInt(uint, int) bool    // Submit an answer.
	DrySubmit(uint, string) bool // Print an answer.
	DrySubmitInt(uint, int) bool // Print an answer.
}

// Inputs is queue of inputs.
type Inputs []input

// Iterate creates Parser for each input and returns them in channel.
func (ins *Inputs) Iterate() <-chan Parser {
	ch := make(chan Parser)
	go func() {
		for _, in := range *ins {
			ch <- newParser(in)
		}
		close(ch)
	}()
	return ch
}

////////////////////////////////////////////////////////////////////////
// Implementation

// DANGER! input HAS to be initialized and have ALL interface methods implemented.
// Forget to initialize -> runtime panic.
// Forget to implement interface method and then call it -> loop leading to stack overflow.
type inputBase struct {
	input
	name  string // test1, or github
	valid [2]bool
}

func todo(s string) {
	panic(fmt.Sprint("Not implemented yet: ", s))
}

func (i *inputBase) Name() string {
	return i.name
}

func (i *inputBase) SubmitInt(part uint, val int) bool {
	return i.Submit(part, strconv.Itoa(val))
}

var dry bool

func (i *inputBase) DrySubmit(part uint, val string) bool {
	var prev bool
	prev, dry = dry, true
	out := i.Submit(part, val)
	dry = prev
	return out
}

func (i *inputBase) DrySubmitInt(part uint, val int) bool {
	return i.DrySubmit(part, strconv.Itoa(val))
}

func (i *inputBase) canProcess(part uint) bool {
	switch part {
	case 2:
		return !i.Part(1) || i.valid[0]
	}
	return true
}
