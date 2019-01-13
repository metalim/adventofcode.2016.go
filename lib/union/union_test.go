package union

import "testing"

var _t *testing.T

func verify(v, ex interface{}) {
	if v != ex {
		_t.Fatalf("got:%v expected:%v", v, ex)
	}
}

func TestLink(t *testing.T) {
	_t = t
	u := New()

	t.Log("single node")
	u.Link(1)
	verify(len(u.Nodes), 1)
	verify(len(u.Unions), 1)

	t.Log("first link", u)
	u.Link(1, 2)
	verify(len(u.Nodes), 2)
	verify(len(u.Unions), 1)
	verify(u.Nodes[1] == u.Nodes[2], true)

	t.Log("same link", u)
	u.Link(2, 1)
	verify(len(u.Nodes), 2)
	verify(len(u.Unions), 1)
	verify(u.Nodes[1] == u.Nodes[2], true)

	t.Log("link 2 3", u)
	u.Link(2, 3)
	verify(len(u.Nodes), 3)
	verify(len(u.Unions), 1)
	verify(u.Nodes[1] == u.Nodes[2], true)
	verify(u.Nodes[1] == u.Nodes[3], true)

	t.Log("new union", u)
	u.Link(5, 4)
	verify(len(u.Nodes), 5)
	verify(len(u.Unions), 2)
	verify(u.Nodes[4] == u.Nodes[5], true)

	t.Log("link unions", u)
	u.Link(1, 5)
	verify(len(u.Nodes), 5)
	verify(len(u.Unions), 1)
	verify(u.Nodes[2] == u.Nodes[4], true)
}
