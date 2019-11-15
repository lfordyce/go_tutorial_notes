package concurrency

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrInboxTerminated = errors.New("Inbox terminated")

type Inbox struct {
	mu          sync.Mutex
	nextMessage chan interface{}
	messages    []interface{}
}

func (m *Inbox) ReceiveNext(ctx context.Context) (msg interface{}, err error) {
	var ok bool
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg, ok = <-m.nextMessage:
		if !ok {
			return nil, ErrMailboxTerminated
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.messages) > 0 {
		select {
		case m.nextMessage <- m.messages[0]:
			m.messages = m.messages[1:]
		default:
		}
	}

	return msg, nil
}

func (m *Inbox) Terminate() {
	m.mu.Lock()
	defer m.mu.Unlock()

	select {
	case _, ok := <-m.nextMessage:
		if !ok {
			return
		}
	default:
	}

	close(m.nextMessage)
	m.messages = nil
}

func (m *Inbox) send(msg interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	select {
	case prevMsg, ok := <-m.nextMessage:
		if !ok {
			return
		}
		m.messages = append(m.messages, msg)
		m.nextMessage <- prevMsg
	default:
		m.messages = append(m.messages, msg)
		m.nextMessage <- m.messages[0]
		m.messages = m.messages[1:]
	}
}

type Mail struct {
	inbox *Inbox
}

func (m *Mail) Send(msg interface{}) {
	m.inbox.send(msg)
}

func RunInbox() {
	inbox := &Inbox{nextMessage: make(chan interface{}, 1)}
	mail := &Mail{inbox: inbox}
	wg := new(sync.WaitGroup)
	received := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for {
				msg, err := inbox.ReceiveNext(context.Background())
				if err != nil {
					fmt.Printf("[%d] Received: %q\n", i, err)
					return
				}
				fmt.Printf("[%d] Received: %q\n", i, msg)

				received <- struct{}{}
			}
		}(i)
	}

	mail.Send("Test message 1")
	//<-received
	mail.Send("Test message 2")
	//<-received

	inbox.Terminate()
	wg.Wait()
}
