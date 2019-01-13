package circular

// NewList with values.
func NewList(vals ...Value) Buffer {
	l := &list{}
	for _, val := range vals {
		l.InsertAfter(val)
	}
	l.Skip(1) // back to cur
	return l
}

// Length of the list.
func (l *list) Length() int {
	return l.len
}

// Skip steps.
func (l *list) Skip(steps Value) Buffer {
	if l.cur == nil {
		return l
	}
	for ; steps > 0; steps-- {
		l.cur = l.cur.next
	}
	for ; steps < 0; steps++ {
		l.cur = l.cur.prev
	}
	return l
}

// InsertBefore cursor.
func (l *list) InsertBefore(val Value) Buffer {
	if l.cur == nil {
		l.cur = &node{val: val}
		l.cur.next = l.cur
		l.cur.prev = l.cur
		l.len = 1
		return l
	}
	l.cur = &node{val: val, next: l.cur, prev: l.cur.prev}
	l.cur.prev.next = l.cur
	l.cur.next.prev = l.cur
	l.len++
	return l
}

// InsertAfter cursor, and select new node (== pos+1 in array).
func (l *list) InsertAfter(val Value) Buffer {
	if l.cur == nil {
		l.cur = &node{val: val}
		l.cur.next = l.cur
		l.cur.prev = l.cur
		l.len = 1
		return l
	}
	l.cur = &node{val: val, next: l.cur.next, prev: l.cur}
	l.cur.prev.next = l.cur
	l.cur.next.prev = l.cur
	l.len++
	return l
}

// Get value at cursor.
func (l *list) Get() (val Value) {
	if l.cur == nil {
		return
	}
	return l.cur.val
}

// Set value at cursor.
func (l *list) Set(val Value) {
	l.cur.val = val // panic if list is empty
}

func (l *list) Reverse(steps int) Buffer {
	panic("Reverse is not implemented for List. Use Slice")
}
func (l *list) Pos() int {
	panic("Pos is not implemented yet.")
}

func (l *list) Goto(pos int) {
	panic("Goto is not implemented yet.")
}

type list struct {
	cur *node
	len int
}

type node struct {
	val        Value
	next, prev *node
}
