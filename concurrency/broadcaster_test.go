package concurrency

import (
	"testing"
)

func TestBroadCaster(t *testing.T) {
	w := &bworker{
		source: make(chan interface{}),
	}

	globalQuit := make(chan struct{})
	w.quit = globalQuit

	bcast := &threadSafeSlice{
		workers: make([]*bworker, 0, 1),
	}

	quit := make(chan bool)

	w.Start()
	bcast.Push(w)

	values := make(chan interface{})

	go broadcaster(bcast, values)
	go generator(values, quit)

	<-quit

}

func broadcaster(bcast *threadSafeSlice, ch chan interface{}) {
	for {
		msg := <-ch
		bcast.Iter(func(w *bworker) {
			w.source <- msg
		})
	}
}

func generator(ch chan interface{}, quit chan bool) {

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(quit)

	//v := 0
	//for {
	//	ch <- v
	//	v++
	//}
}
