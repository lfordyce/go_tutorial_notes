package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestPubSubConsumer(t *testing.T) {
	N := 10
	q := make(chan int, N)

	prodCtrl := make(chan control)
	consCtrl := make(chan control)

	prodDone := make(chan bool)
	consDone := make(chan bool)

	ProduceTo(q, prodCtrl, prodDone)
	ConsumeFrom(q, consCtrl, consDone)

	// wait for a moment, to let them produce and consume
	timer := time.NewTimer(10 * time.Millisecond)
	<-timer.C

	// tell producer to pause
	fmt.Printf("telling producer to pause\n")
	prodCtrl <- sleep

	// wait for a second
	timer = time.NewTimer(1 * time.Second)
	<-timer.C

	// tell consumer to pause
	fmt.Printf("telling consumer to pause\n")
	consCtrl <- sleep

	// tell them both to finish
	prodCtrl <- die
	consCtrl <- die

	// wait for that to actually happen
	<-prodDone
	<-consDone
}
