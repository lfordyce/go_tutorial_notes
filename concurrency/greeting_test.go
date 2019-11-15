package concurrency

import (
	"context"
	"testing"
	"time"
)

func TestGreeter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	names := make(chan []byte)
	greetings := make(chan []byte)
	errs := make(chan error)

	go Greeter(ctx, names, greetings, errs)

	// send name and get the greeting
	names <- []byte("Batman")
	//greeting := <-greetings

	//if string(greeting) != "Hello Batman" {
	//	t.Error("unexpected message")
	//}

	select {
	case <-time.After(1 * time.Second):
		t.Error("timed out waiting for greeting")
	case greeting := <-greetings:
		if string(greeting) != "Hello Batman" {
			t.Error("unexpected message")
		}
	}
}
