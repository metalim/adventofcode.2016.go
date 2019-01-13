package field

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

// Cell = int
type Cell = int

// Pos = image.Point
type Pos = image.Point

// Rect = image.Rectangle
type Rect = image.Rectangle

// Field is two-dimensional.
type Field interface {
	Get(Pos) Cell
	Set(Pos, Cell)
	Bounds() Rect
	Default() Cell
	SetDefault(Cell)
}

////////////////////////////////////////////////////////////////////////
// Base methods, can't use methods of derived classes.

type fieldBase struct {
	b   Rect
	def Cell
}

// Default cell value.
func (f *fieldBase) Default() Cell {
	return f.def
}

// SetDefault cell value.
func (f *fieldBase) SetDefault(c Cell) {
	f.def = c
}

// Bounds AABB.
func (f *fieldBase) Bounds() Rect {
	return f.b
}

////////////////////////////////////////////////////////////////////////
// Common "methods", actually functions.

// Print field
func Print(f Field) {
	bs := f.Bounds()
	r := make([]string, 0, bs.Dx())
	for y := bs.Min.Y; y < bs.Max.Y; y++ {
		r = r[:0]
		for x := bs.Min.X; x < bs.Max.X; x++ {
			c := f.Get(Pos{x, y})
			if c == 0 {
				r = append(r, ".")
				continue
			}
			r = append(r, strconv.Itoa(int(c)))
		}
		fmt.Println(strings.Join(r, " "))
	}
}

// Rotate90 rectangle 90°.
func Rotate90(f Field, rect Rect) {
	cy := rect.Dy()
	cx := rect.Dx()
	if cx != cy {
		panic("rectangle must be square!")
	}
	if cx%2 == 1 { // odd? -> rotate (N x N+1)/(N+1 x N) rectangles around center piece.
		cx += 2
	}
	cx /= 2
	cy /= 2
	for y := 0; y < cy; y++ {
		for x := 0; x < cx; x++ {
			p1 := Pos{rect.Min.X + x, rect.Min.Y + y}
			p2 := Pos{rect.Max.X - 1 - y, rect.Min.Y + x}
			p3 := Pos{rect.Max.X - 1 - x, rect.Max.Y - 1 - y}
			p4 := Pos{rect.Min.X + y, rect.Max.Y - 1 - x}

			v1, v2, v3, v4 := f.Get(p1), f.Get(p2), f.Get(p3), f.Get(p4)
			f.Set(p2, v1)
			f.Set(p3, v2)
			f.Set(p4, v3)
			f.Set(p1, v4)
		}
	}
}

// FlipVert a rectangle.
func FlipVert(f Field, rect Rect) {
	dy := rect.Dy()
	for y := 0; y < dy/2; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			p1, p2 := Pos{x, rect.Min.Y + y}, Pos{x, rect.Max.Y - 1 - y}
			v1, v2 := f.Get(p1), f.Get(p2)
			f.Set(p1, v2)
			f.Set(p2, v1)
		}
	}
}

// FillFromString from start position.
func FillFromString(f Field, start Pos, s string) {
	for y, l := range strings.Split(s, "\n") {
		for x, r := range l {
			f.Set(start.Add(Pos{x, y}), Cell(r))
		}
	}
}

// ToString rect.
func ToString(f Field, rect Rect) string {
	var s strings.Builder
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		if s.Len() > 0 {
			s.WriteByte('\n')
		}
		for x := rect.Min.X; x < rect.Max.X; x++ {
			s.WriteRune(rune(f.Get(Pos{x, y})))
		}
	}
	return s.String()
}

////////////////////////////////////////////////////////////////////////

// Dir4 0..3 for axis aligned
type Dir4 uint

// Dir44 is Dir4 extended to 4..7 for diagonals
type Dir44 Dir4

// Dir4 and Dir44 constants
// Since we don't define direction of Y, but only it's value,
// P = +1, N = -1, and not North/Up or South/Down.
const (
	Dir4P0 Dir4 = iota
	Dir40P
	Dir4N0
	Dir40N
	Dir44PP Dir44 = iota
	Dir44NP
	Dir44NN
	Dir44PN
	Dir4Mask  = 3
	Dir44Mask = 7
	Dir4All   = 0x0f
	Dir44All  = 0xff
)

// Step for backwards compatibility.
func Step(p Pos, d4 Dir4) Pos {
	return Step4(p, d4)
}

// Step4 in specified direction:
// * 0123 -> ESWN / RDLU, for x axis to the right, and y axis down.
// * 0123 -> ENWS / RULD, for x axis to the right, and y axis up.
// * 4567, etc == 0123
// == dir * 90°
func Step4(p Pos, d4 Dir4) Pos {
	return Pos{p.X + (1-int(d4&Dir4Mask))%2, p.Y + (2-int(d4&Dir4Mask))%2}
}

// Step44 in specified direction, including diagonals **after** the main sequence:
// 0123:4567 -> E S W N : SE SW NW NE, so 0-3 directions mean the same as in Step.
func Step44(p Pos, d44 Dir44) Pos {
	out := Step(p, Dir4(d44))
	if d44&4 != 0 { // diagonal? -> additional step in +1 direction.
		out = Step(out, Dir4(d44+1))
	}
	return out
}

// Dir8 is 0..7 including diagonals: 0 1 2 3 4 5 6 7 -> E SE S SW W NW N NE
type Dir8 uint

// Dir8 constants
const (
	Dir8P0 Dir8 = iota
	Dir8PP
	Dir80P
	Dir8NP
	Dir8N0
	Dir8NN
	Dir80N
	Dir8PN
	Dir8Mask = 7
	Dir8All  = 0xff
)

// Step8 in specified direction, including diagonals **inside** main sequence:
// 0 1 2 3 4 5 6 7 -> E SE S SW W NW N NE
// 0246 -> ESWN
// == dir * 45°
// To step in axis aligned directions add steps of 2: 0 2 4 6.
func Step8(p Pos, d8 Dir8) Pos {
	d4 := Dir4(d8 >> 1)
	out := Step4(p, d4)
	if d8&1 != 0 { // diagonal
		out = Step4(out, d4+1)
	}
	return out
}

func abs(n int) int {
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
}

// Manh return manhattan distance between points.
func Manh(p1, p2 Pos) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// Walk on arbitrary 2d grid with callbacks for tile check, direction get, and actual step,
// and return end position and distance walked.
// Callbacks:
// * canStepOn(p) should return true if you can Walk onto the tile.
// * walk(p) should store path (if needed) and return true to continue walking.
// * getDirections(p,d) should return bitmask, where each bit from bits 0-7 represents allowed direction.
//   0 1 2 3 -> E S W N, 4 5 6 7 -> SE SW NW NE
//   * each allowed direction will then be checked with canStepOn.
func Walk(p Pos, d Dir8, canStepOn func(Pos) bool, stepOn func(Pos, Dir8) int) (end Pos, steps int) {
	if canStepOn != nil && !canStepOn(p) {
		return
	}

WALKING:
	for {
		steps++
		dirs := Dir8All
		if stepOn != nil {
			dirs = stepOn(p, d)
		}
		if dirs != 0 {
			// start with same direction as before, then turn 45° right on each try.
			dirs = (dirs>>d | dirs<<(8-d)) & Dir8All
			for ; dirs > 0; dirs, d = dirs>>1, (d+1)%8 {

				if dirs&1 == 0 { // direction not allowed? -> skip it.
					continue
				}

				p1 := Step8(p, d)
				if canStepOn != nil && !canStepOn(p1) { // can't step in this direction? -> skip it.
					continue
				}

				p = p1
				continue WALKING
			}
		}

		// no more moves -> exit.
		return p, steps
	}
}
