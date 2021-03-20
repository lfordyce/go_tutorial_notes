package watchtower

import (
	"fmt"
	"sync"
)

type Observer interface {
	OnNotify(Event)
}

type Observer2 interface {
	NotifyCallback(Event)
}

type Notifier interface {
	Register(Observer)
	Deregister(observer Observer)
	Notify(Event)
}

type Event struct {
	Data int64
}

type WatchTower struct {
	observer sync.Map
}

//func NewWatchTower() *WatchTower {
//	return &WatchTower{observer: new(sync.Map)}
//}

func (wt *WatchTower) Register(l Observer) {
	wt.observer.Store(l, struct{}{})
}

func (wt *WatchTower) Deregister(l Observer) {
	wt.observer.Delete(l)
}

func (wt *WatchTower) Notify(e Event) {
	wt.observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).OnNotify(e)
		return true
	})
}

type eventObserver struct {
	id   int
	zone string
}

type eventNotifier struct {
	observers map[Observer]struct{}
}

func (o *eventObserver) OnNotify(e Event) {
	//if e.Data == o.id {
	//	fmt.Printf()
	//}

	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (n *eventNotifier) Register(l Observer) {
	n.observers[l] = struct{}{}
}

func (n *eventNotifier) Deregister(l Observer) {
	delete(n.observers, l)
}

func (n *eventNotifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}
