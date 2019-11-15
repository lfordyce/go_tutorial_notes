package concurrency

import (
	"fmt"
	"sync"
)

type bworker struct {
	source chan interface{}
	quit   chan struct{}
}

func (w *bworker) Start() {
	w.source = make(chan interface{}, 10) // some buffer size to avoid blocking
	go func() {
		for {
			select {
			case msg := <-w.source:
				// do something with msg
				fmt.Println(msg)

			case <-w.quit: // will explain this in the last section
				return
			}
		}
	}()
}

type threadSafeSlice struct {
	sync.Mutex
	workers []*bworker
}

func (slice *threadSafeSlice) Push(w *bworker) {
	slice.Lock()
	defer slice.Unlock()

	slice.workers = append(slice.workers, w)
}

func (slice *threadSafeSlice) Iter(routine func(*bworker)) {
	slice.Lock()
	defer slice.Unlock()

	for _, worker := range slice.workers {
		routine(worker)
	}
}
