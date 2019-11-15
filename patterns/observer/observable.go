package observer

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

type Observable struct {
	observers []chan int
	mu        *sync.Mutex
}

func (o *Observable) Attach(c chan int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.observers = append(o.observers, c)
}

func (o *Observable) Detach(c chan int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for i, v := range o.observers {
		if v == c {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			return
		}
	}
}

func (o *Observable) Notify(evt int) {
	for _, v := range o.observers {
		v <- evt
	}
}

type Observer struct {
	Food string
	ch   chan int
}

func (obs *Observer) Observe() {
	evt := <-obs.ch
	fmt.Println("Food:", obs.Food, "Event:", evt)
	wg.Done()
}
