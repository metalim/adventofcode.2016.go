package union

// Unions set.
type Unions struct {
	Nodes  map[int]int
	Unions map[int][]int
}

// New inits maps.
func New() Unions {
	return Unions{map[int]int{}, map[int][]int{}}
}

// Link a to b.
func (o *Unions) Link(sn ...int) {

	// create nodes if necessary.
	for _, a := range sn {
		if _, ok := o.Nodes[a]; !ok {
			o.Nodes[a] = a
			o.Unions[a] = []int{a}
		}
	}

	// join unions
	a := sn[0]
	u := o.Nodes[a]
	for _, b := range sn[1:] {
		v := o.Nodes[b]
		if u == v {
			continue
		}
		o.Unions[u] = append(o.Unions[u], o.Unions[v]...)
		for _, n := range o.Unions[v] {
			o.Nodes[n] = u
		}
		delete(o.Unions, v)
	}
}
