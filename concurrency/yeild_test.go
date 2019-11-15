package concurrency

import (
	"fmt"
	"testing"
)

func TestYield(t *testing.T) {
	myMapper := func(yield yieldFn) {
		for i := 0; i < 5; i++ {
			if keepGoing := yield(i); !keepGoing {
				return
			}
		}
	}
	iterator, cancel := mapperToIterator(myMapper)
	defer cancel()
	for value, notDone := iterator(); notDone; value, notDone = iterator() {
		fmt.Printf("value %d\n", value.(int))
	}
}
