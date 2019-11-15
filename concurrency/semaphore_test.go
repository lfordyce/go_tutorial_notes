package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for _, i := range ints {
		wg.Add(1)
		sem <- struct{}{}
		go func(id int) {
			defer wg.Done()
			printWithSleep(id)
			//fmt.Printf("ID: %d\n", i)
			//fmt.Println(i*2)
			<-sem
		}(i)
	}
	wg.Wait()
}

func printWithSleep(id int) {
	time.Sleep(time.Second)
	fmt.Printf("ID: %d\n", id*2)
}
