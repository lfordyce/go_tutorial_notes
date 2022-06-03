package concurrency

import (
	"math/rand"
	"time"
)

type Data struct {
	secretOfTheSecond int
	readRequest       chan chan int
}

// Run writes the data periodically to the struct. This meant to run in
// it's own goroutine
func (d *Data) Run() {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			d.secretOfTheSecond = gen.Int()
		case responseChan := <-d.readRequest:
			responseChan <- d.secretOfTheSecond
		}
	}
}

// Get does an ad-hoc read of the data. This is the user interface
func (d *Data) Get() int {
	responseChan := make(chan int)
	d.readRequest <- responseChan
	response := <-responseChan
	return response
}
