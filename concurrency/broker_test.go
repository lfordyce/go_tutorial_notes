package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBroker(t *testing.T) {
	// Create and start a broker:
	b := NewBroker()
	go b.Start()

	// Create and subscribe 3 clients:
	clientFunc := func(id int) {
		msgCh := b.Subscribe()
		for {
			fmt.Printf("Client %d got message: %v\n", id, <-msgCh)
		}
	}
	for i := 0; i < 3; i++ {
		go clientFunc(i)
	}

	// Start publishing messages:
	go func() {
		for msgId := 0; ; msgId++ {
			b.Publish(fmt.Sprintf("msg#%d", msgId))
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
}
