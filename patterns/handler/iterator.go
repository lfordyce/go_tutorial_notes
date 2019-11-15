package handler

import "errors"

var errDone = errors.New("Done")

// Collection is a dummy collection.
// For the sake of testing, we'll store items in reverse order
type Collection struct {
	items []int
}

// Push adds an item at the beginning of the collection
func (c *Collection) Push(n int) {
	// Items stored in reverse order, so append(.., n) makes it the first item
	c.items = append(c.items, n)
}

// NewCollection creates a new collection initialized with numbers from 1 to n.
func NewCollection(n int) *Collection {
	var c Collection
	for i := n; i > 0; i-- {
		c.Push(i)
	}
	return &c
}

// ------ Array-style accessors ------

// ValueAt returns the Nth item.
func (c *Collection) ValueAt(n int) int {
	return c.items[len(c.items)-n-1] // let it panic if out of bounds
}

// Len returns the number of items in the collection
func (c *Collection) Len() int { return len(c.items) }

// ------ Closure iterators ------

// ClosureIter returns a closure based iterator and a boolean set to true if there
// are values to be read.
func (c *Collection) ClosureIter() (f func() (int, bool), hasNext bool) {
	l := len(c.items)
	hasNext = l > 0
	f = func() (int, bool) {
		l--
		return c.items[l], l > 0
	}
	return
}

// ClosureIter2 is almost the same as above but returns a next() and hasNext() function (just to
// show that this approach is slower).
func (c *Collection) ClosureIter2() (next func() int, hasNext func() bool) {
	l := len(c.items)
	next = func() int { l--; return c.items[l] }
	hasNext = func() bool { return l > 0 }
	return
}

// ClosureIter3 is a variation on ClosureIter that returns an error when trying to read past
// the end of the collection. It's slower than ClosureIter that uses a "predictive" hasNext, thus
// skipping a conditional branch.
func (c *Collection) ClosureIter3() (next func() (int, error)) {
	l := len(c.items)
	return func() (int, error) {
		if l > 0 {
			l--
			return c.items[l], nil
		}
		return 0, errDone
	}
}

// ------ Stateful iterator ------

// this iterator struct does not have to be exported, but can be if you have some
// generic Iterator interface that you want/need to implement.
type iterator struct {
	*Collection // neat!
	index       int
}

// Next returns the next item in the collection.
func (i *iterator) Next() int {
	i.index--
	return i.items[i.index]
}

// NextErr is a variation of HasNext that returns an error when trying to read past
// the end of the collection (no need to use HasNext). It is slower with this particular
// Collection implementation since it introduces an additional conditional branch.
// More complex collections may not suffer that much.
func (i *iterator) NextErr() (int, error) {
	if i.index == 0 {
		return 0, errDone
	}
	i.index--
	return i.items[i.index], nil
}

// HasNext return true if there are values to be read.
func (i *iterator) HasNext() bool {
	return i.index > 0
}

// StatefulIter returns our stateful iterator.
func (c *Collection) StatefulIter() *iterator {
	return &iterator{c, len(c.items)}
}
