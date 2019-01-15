package freq

import "sort"

// IMap values to frequencies.
type IMap map[interface{}]int

// Add value.
func (m IMap) Add(v interface{}) {
	m[v]++
}

// IRecord to get values from generic map into slice.
type IRecord struct {
	V interface{}
	F int
}

// Sorted slice of records.
func (m IMap) Sorted() (out []IRecord) {
	for v, f := range m {
		out = append(out, IRecord{v, f})
	}
	sort.Sort(byFreq(out))
	return
}

type byFreq []IRecord

func (v byFreq) Len() int {
	return len(v)
}

func (v byFreq) Less(i, j int) bool {
	return v[i].F > v[j].F
}

func (v byFreq) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
