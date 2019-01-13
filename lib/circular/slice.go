package circular

// NewSlice with values
func NewSlice(vals ...Value) Buffer {
	return &slice{data: vals}
}

func (s *slice) Skip(steps int) Buffer {
	mod := len(s.data)
	if mod > 0 {
		s.cur = ((s.cur+steps)%mod + mod) % mod
	}
	return s
}

func (s *slice) InsertBefore(val Value) Buffer {
	if len(s.data) == 0 {
		s.data = append(s.data, val)
		return s
	}
	var empty Value
	s.data = append(s.data, empty)
	copy(s.data[s.cur+1:], s.data[s.cur:])
	s.data[s.cur] = val
	return s
}

func (s *slice) InsertAfter(val Value) Buffer {
	if len(s.data) == 0 {
		s.data = append(s.data, val)
		return s
	}
	var empty Value
	s.data = append(s.data, empty)
	s.cur++
	copy(s.data[s.cur+1:], s.data[s.cur:])
	s.data[s.cur] = val
	return s
}

func (s *slice) Get() (val Value) {
	return s.data[s.cur]
}

func (s *slice) Set(val Value) {
	s.data[s.cur] = val
}

func (s *slice) Length() int {
	return len(s.data)
}

func (s *slice) Reverse(steps int) Buffer {
	for i := 0; i < steps/2; i++ {
		a := (s.cur + i) % len(s.data)
		b := (s.cur + steps - 1 - i) % len(s.data)
		s.data[a], s.data[b] = s.data[b], s.data[a]
	}
	return s
}

func (s *slice) Pos() int { return s.cur }
func (s *slice) Goto(pos int) {
	s.cur = pos
}

type slice struct {
	data []Value
	cur  int
}
