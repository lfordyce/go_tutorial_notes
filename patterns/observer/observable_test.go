package observer

import (
	"sync"
	"testing"
	"time"
)

func TestObserver_Observe(t *testing.T) {
	o := &Observable{}
	o.mu = &sync.Mutex{}

	obs := []Observer{
		{"Eggs", make(chan int, 3)},
		{"Bacon", make(chan int, 2)},
	}

	wg = &sync.WaitGroup{}
	wg.Add(len(obs))
	for _, v := range obs {
		o.Attach(v.ch)
		go v.Observe()
	}

	go func() {
		<-time.After(1 * time.Second)
		o.Notify(3)
		o.Notify(4)
	}()

	go func() {
		<-time.After(3 * time.Second)
		o.Notify(5)
	}()

	wg.Wait()
}
