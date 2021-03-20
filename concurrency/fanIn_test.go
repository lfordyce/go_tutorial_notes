package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepeatFunc(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	randFn := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFunc(done, randFn), 10) {
		fmt.Println(num)
	}
}

func TestTeeOperation(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeater(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
