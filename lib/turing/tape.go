package turing

// Tape for turing machine.
type Tape struct {
	d [2][]int
	p int
}

// GetAt - for iteration.
func (t *Tape) GetAt(pos int) int {
	i, p := t.sel(pos)
	if p < len(t.d[i]) {
		return t.d[i][p]
	}
	return 0
}

// Get value at current position.
func (t *Tape) Get() int {
	return t.GetAt(t.p)
}

// Set value at current position.
func (t *Tape) Set(v int) {
	i, p := t.sel(t.p)
	if p >= len(t.d[i]) {
		t.d[i] = append(t.d[i], make([]int, p+1-len(t.d[i]))...) // extra space
	}
	t.d[i][p] = v
}

// Left 1 position.
func (t *Tape) Left() {
	t.p--
}

// Right 1 position.
func (t *Tape) Right() {
	t.p++
}

// Goto position.
func (t *Tape) Goto(p int) {
	t.p = p
}

// GoRight by distance.
func (t *Tape) GoRight(d int) {
	t.p += d
}

// Bounds inclusive
func (t *Tape) Bounds() (int, int) {
	return -len(t.d[1]), len(t.d[0]) - 1
}

// select slice and position.
func (t *Tape) sel(p int) (int, int) {
	if p < 0 {
		return 1, -p - 1
	}
	return 0, p
}
