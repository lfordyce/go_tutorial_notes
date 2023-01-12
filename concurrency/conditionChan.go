package concurrency

type Locker interface {
	Lock()
	Unlock()
}

type barrier chan struct{}

type Cond struct {
	L  Locker
	bs chan barrier
}

func NewCond(l Locker) Cond {
	c := Cond{
		L:  l,
		bs: make(chan barrier, 1),
	}
	// Waits will block until signalled to continue.
	c.bs <- make(barrier)
	return c
}

func (c Cond) Broadcast() {
	// Acquire barrier
	b := <-c.bs
	// release all waiter
	close(b)
	// create a new barrier for future calls.
	c.bs <- make(barrier)
}

func (c Cond) Signal() {
	// Acquire barrier
	b := <-c.bs
	// Release one waiter if there are any waiting.
	select {
	case b <- struct{}{}:
	default:
	}
	// Release barrier.
	c.bs <- b
}

// Wait performs two actions atomically:
// * Call Unlock
// * Suspend execution
// To do so we receive the current barrier, call Unlock while we still
// hold it and release it. This guarantees that nothing else has happened
// in the meantime.
// After this operation we wait on the barrier we received, which
// might not reflect the current one (as intended).
func (c Cond) Wait() {
	// Acquire barrier.
	b := <-c.bs
	// Unlock while in critical section.
	c.L.Unlock()
	// Release barrier.
	c.bs <- b
	// Wait for release on the value of barrier that was valid during
	// the call to Unlock.
	<-b
	// We were unblocked, acquire lock.
	c.L.Lock()
}
