package generics

import (
	"sync"
	"testing"
)

func TestPriorityChannels(t *testing.T) {
	var wg sync.WaitGroup
	produceDone := make(chan struct{})
	wg.Add(1)

	high := make(chan int)
	low := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				low <- i
			} else {
				high <- i
			}
		}
		close(produceDone)
		close(high)
		close(low)
		wg.Done()
	}()

	for item := range Plex(high, low) {
		t.Logf("received item: %d", item)
	}
	<-produceDone
	t.Logf("producer done, closing channel")
	wg.Wait()
}
