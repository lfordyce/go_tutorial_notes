package handler

import (
	"fmt"
	"testing"
)

func TestIteratorClosure(t *testing.T) {
	c := NewCollection(5)
	for next, hasNext := c.ClosureIter(); hasNext; {
		var v int
		v, hasNext = next()
		fmt.Printf("%d \n", v)
	}

	c2 := NewCollection(5)
	for next, hasNext := c2.ClosureIter2(); hasNext(); {
		fmt.Printf("%d \n", next())
	}

	c3 := NewCollection(5)
	iter := c3.ClosureIter3()
	for v, err := iter(); err == nil; v, err = iter() {
		fmt.Printf("%d \n", v)
	}
}

func TestStatefulIterator(t *testing.T) {
	c := NewCollection(5)
	for iter := c.StatefulIter(); iter.HasNext(); {
		fmt.Printf("%d ", iter.Next())
	}
}
