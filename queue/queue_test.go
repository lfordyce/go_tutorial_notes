package queue

import (
	"fmt"
	"testing"
	"time"
)

type event int

func (e event) OnTimer(t time.Time) {
	fmt.Printf("  Event %d executed at %v\n", int(e), t)
}

func TestQueue(t *testing.T) {
	queue := New()

	// Schedule an event each day from Jan 1 to Jan 7, 2015.
	tm := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 1; i <= 7; i++ {
		queue.Schedule(event(i), tm)
		tm = tm.Add(24 * time.Hour)
	}

	fmt.Println("Advancing to Jan 4...")
	queue.Advance(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC))

	fmt.Println("Advancing to Jan 10...")
	queue.Advance(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC))
}
