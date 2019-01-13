package circular

// Value stored in circular.Buffer.
type Value = int

// Buffer interface. Note Buffer return values are just for chaining, not required "to reset the pointer".
type Buffer interface {
	Skip(steps int) Buffer
	InsertBefore(Value) Buffer
	InsertAfter(Value) Buffer
	Get() Value
	Set(Value)
	Length() int
	Reverse(length int) Buffer
	Pos() int
	Goto(pos int)
}
