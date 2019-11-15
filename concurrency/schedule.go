package concurrency

import (
	"log"
	"sync"
	"time"
)

func shcedule(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	return ticker
}

func oneSec() {
	log.Println("one Sec")
}

func twoSec() {
	log.Println("two Sec")
}

// TimeHolder interface to dynamically get and set time duration interval
type TimeHolder interface {
	Get() time.Duration
	Set(duration time.Duration)
}

type heartbeat struct {
	period time.Duration
	ticker time.Ticker
	mu     *sync.Mutex
}

func createHeartbeat(period time.Duration) *heartbeat {
	return &heartbeat{period: period, ticker: *time.NewTicker(period), mu: &sync.Mutex{}}
}

func (h *heartbeat) restHeartbeat(period time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.ticker.Stop()
	h.ticker = *time.NewTicker(period)
}
