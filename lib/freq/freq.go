package freq

// Interface for frequency map.
type Interface interface {
	Add(interface{})
	Sorted() []Record
	Sorted2() []Record
}

// Record for docs.
type Record interface{}
