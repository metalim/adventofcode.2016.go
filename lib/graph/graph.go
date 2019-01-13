package graph

type idt int

// Graph .
type Graph struct {
	nodes    map[idt]*Node
	links    map[idt]*Link
	idn, idl idt
}

// NewGraph creates maps.
func NewGraph() *Graph {
	return &Graph{nodes: map[idt]*Node{}, links: map[idt]*Link{}}
}

// Link - not using name "edge", as the name is stupid.
type Link struct {
	id    idt
	Data  interface{}
	nodes [2]*Node // we need it for g.Unlink, and if we're interested in nodes in general.
	// or
	// nodes map[idt]*nodes // can be done for multilinks ("hyperedges" in "hypergraphs").
}

// Path of Links.
type Path []*Link

type linksMap map[*Link]*Node

// Node - not using name "vertice", for same reasons.
type Node struct {
	id    idt
	Data  interface{}
	links linksMap
}

// Node returns existing node or creates new one.
func (g *Graph) Node(id int) *Node {
	n, ok := g.nodes[idt(id)]
	if !ok {
		n = &Node{id: g.idn, links: linksMap{}}
		g.idn++
		g.nodes[idt(id)] = n
	}
	return n
}

// Link bidirectional.
func (g *Graph) Link(a, b *Node) *Link {
	ln := &Link{id: g.idl, nodes: [2]*Node{a, b}}
	g.links[g.idl] = ln
	g.idl++
	// every to every.
	a.links[ln] = b // works for self-linking loops,
	b.links[ln] = a // as this just replaces with same value.
	return ln
}

// LinkTo one direction. Note this makes node b to not know anything about the link.
func (g *Graph) LinkTo(a, b *Node) *Link {
	ln := &Link{id: g.idl, nodes: [2]*Node{a, b}}
	g.links[g.idl] = ln
	g.idl++

	a.links[ln] = b // a to b only.
	return ln
}

func (n *Node) copyLinks() linksMap {
	out := linksMap{}
	for l, n := range n.links {
		out[l] = n
	}
	return out
}

// IteratePathsFrom starts from specified node and iterates all possible paths in DFS.
// Path is followed only if callback returns true.
func (g *Graph) IteratePathsFrom(start *Node, fn func(Path) bool) {
	lns := start.copyLinks()
	used := map[*Link]bool{}
	var path Path
	var queues []linksMap // expensive.

APPEND:
	for {
		for ln, n2 := range lns {
			delete(lns, ln) // it's a queue in map type :-)
			if used[ln] {   // skip already used links
				continue
			}
			next := append(path, ln)
			if !fn(next) { // follow this path only if callback returns true.
				continue
			}
			path = next
			used[ln] = true
			queues = append(queues, lns)
			lns = n2.copyLinks()
			continue APPEND // with new set of links
		}
		// no more appends -> pop last link, and continue with remaining links.
		if len(queues) == 0 { // nothing to pop -> exit.
			break
		}
		lns, queues = queues[len(queues)-1], queues[:len(queues)-1]
		var released *Link
		released, path = path[len(path)-1], path[:len(path)-1]
		delete(used, released)
	}
}
