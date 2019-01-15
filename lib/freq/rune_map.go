package freq

import "sort"

// RuneMap values to frequencies.
type RuneMap map[rune]int

// Add value.
func (m RuneMap) Add(v rune) {
	m[v]++
}

// RuneRecord to get values from map into slice.
type RuneRecord struct {
	V rune
	F int
}

// Sorted slice of RuneRecords, additionally sorted by rune.
func (m RuneMap) Sorted() (out []RuneRecord) {
	for v, f := range m {
		out = append(out, RuneRecord{v, f})
	}
	sort.Sort(byFreqAndValue(out))
	return
}

type byFreqAndValue []RuneRecord

func (v byFreqAndValue) Len() int {
	return len(v)
}

func (v byFreqAndValue) Less(i, j int) bool {
	return v[i].F > v[j].F || v[i].F == v[j].F && v[i].V < v[j].V
}

func (v byFreqAndValue) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
