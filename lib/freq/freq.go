package freq

import "sort"

// Interface for frequency map.
type Interface interface {
	Add(interface{})
	Sorted() []Record
}

// NewMap to count frequencies.
func NewMap() Interface {
	return &Map{}
}

// Map values to frequencies.
type Map map[interface{}]int

// Add value.
func (m Map) Add(v interface{}) {
	m[v]++
}

// Record to get values from map into slice.
type Record struct {
	V interface{}
	F int
}

// Sorted slice of records.
func (m Map) Sorted() (out []Record) {
	for v, f := range m {
		out = append(out, Record{v, f})
	}
	sort.Sort(byFreq(out))
	return
}

type byFreq []Record

func (v byFreq) Len() int {
	return len(v)
}

func (v byFreq) Less(i, j int) bool {
	return v[i].F < v[j].F
}

func (v byFreq) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
