package watchtower

import (
	"testing"
	"time"
)

func TestWatchTower_Notify(t *testing.T) {
	//watchTower := WatchTower{}
	watchTower := new(WatchTower)

	watchTower.Register(&eventObserver{id: 1})
	watchTower.Register(&eventObserver{id: 2})
	watchTower.Register(&eventObserver{id: 3})
	watchTower.Register(&eventObserver{id: 4})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	ticker := time.NewTicker(time.Second)
	//tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			ticker.Stop()
			return
		case t := <-ticker.C:
			watchTower.Notify(Event{Data: t.UnixNano()})
		}
	}
}
