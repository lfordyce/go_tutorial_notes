package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewHeap(t *testing.T) {
	heap := NewHeap()

	go func() {
		for range time.Tick(3 * time.Second) {
			for n := 0; n < 3; n++ {
				x := rand.Intn(100)
				fmt.Println("push:", x)
				heap.Push(x)
			}
		}
	}()

	for {
		item := heap.Pop()
		fmt.Println("pop:", item)
	}
}
