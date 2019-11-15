package concurrency

import (
	"errors"
	"sync"
)

var ErrMailboxTerminated = errors.New("Mailbox terminated")

type Mailbox struct {
	cond       *sync.Cond
	messages   []interface{}
	terminated bool
}

func (m *Mailbox) ReceiveNext() interface{} {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	for len(m.messages) == 0 && !m.terminated {
		// Mailbox is empty.  Wait for a Signal or Broadcast signalling
		// the receipt of a message or termination of the mailbox before
		// attempting to acquire the mutex and proceed.
		m.cond.Wait()
	}

	if m.terminated {
		return ErrMailboxTerminated
	}

	msg := m.messages[0]
	if len(m.messages) == 1 {
		m.messages = m.messages[:0]
	} else {
		m.messages = m.messages[1:]
	}

	return msg
}

func (m *Mailbox) Terminate() {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	if m.terminated {
		return
	}

	m.terminated = true
	m.messages = nil

	// Wake up all goroutines so they'll immediately realize the
	// mailbox is terminated.
	m.cond.Broadcast()
}

func (m *Mailbox) send(msg interface{}) {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	if m.terminated {
		return
	}

	m.messages = append(m.messages, msg)

	// Wake up a single goroutine to receive the message.
	m.cond.Signal()
}

type Address struct {
	mailbox *Mailbox
}

func (a *Address) Send(msg interface{}) {
	a.mailbox.send(msg)
}
