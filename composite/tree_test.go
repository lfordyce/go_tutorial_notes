package composite

import (
	"fmt"
	"testing"
)

func TestBuildTree_ImmutableNodes(t *testing.T) {

	ct, r := NewTree()
	a := mys{"a"}
	ct.Add(r, &a)
	b := mys{"b"}
	c := mys{"c"}
	ct.Add(&a, &b)
	ct.Add(&a, &c)
	ct.Add(&a, &c) // ignored as c already in the tree
	e := mys{"e"}
	f := mys{"f"}
	h := mys{"h"}
	ct.Add(&c, &e)
	ct.Add(&c, &f)
	ct.Add(&c, &h)
	x := 1
	ct.Add(r, &x)
	ct.Walk(p)

	fmt.Println()

	d := "d"
	d2 := "d"
	slice := &[]int{1, 2, 3, 4}
	BuildTree_ImmutableNodes().
		Add(mys{"a"}).Down().
		Add(&mys{"b"}).
		Add(mys{"c"}).
		Add(&d2).
		Add(d).
		Add("d").Up().
		Add(slice).
		Build().
		Walk(p)

}
