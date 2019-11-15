package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestPlex(t *testing.T) {
	done := make(chan bool)
	defer close(done)

	start := time.Now()
	items := PrepareItems(done)

	workers := make([]<-chan int, 8)
	for i := 0; i < 8; i++ {
		workers[i] = PackItems(done, items, i)
	}

	numPackages := 0
	for range mergePlexer(done, workers...) {
		numPackages++
	}
	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
