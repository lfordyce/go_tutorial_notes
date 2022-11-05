package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMailbox_ReceiveNext(t *testing.T) {
	mbox := &Mailbox{cond: sync.NewCond(new(sync.Mutex))}
	addr := &Address{mailbox: mbox}

	wg := new(sync.WaitGroup)
	received := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for {
				msg := mbox.ReceiveNext()
				fmt.Printf("[%d] Received: %q\n", i, msg)

				if msg == ErrMailboxTerminated {
					return
				}

				received <- struct{}{}
			}
		}(i)
	}

	addr.Send("Test message 1")
	time.Sleep(2 * time.Second)
	<-received
	addr.Send("Test message 2")
	time.Sleep(2 * time.Second)
	<-received
	addr.Send("Test message 3")
	time.Sleep(2 * time.Second)
	<-received

	mbox.Terminate()
	wg.Wait()
}
