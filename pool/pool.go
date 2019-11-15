package pool

import (
	"context"
	"errors"
	"sync"
)

type Pool struct {
	finish chan struct{}
	work   chan<- func()
	wg     sync.WaitGroup
}

// ErrFinished indicates that Finished has been called on a Pool.
var ErrFinished = errors.New("Add: Pool is finished")

// NewPool starts a pool of n workers.
func NewPool(n int) *Pool {
	work := make(chan func())

	p := &Pool{work: work, finish: make(chan struct{})}

	for ; n > 0; n-- {
		p.wg.Add(1)
		go func() {
			for {
				select {
				case f := <-work:
					f()
				case <-p.finish:
					p.wg.Done()
					return
				}
			}
		}()
	}

	return p
}

// Add sends f to a worker goroutine, which then executes it.
// If Finish has been called, Add returns ErrFinished.
// If ctx is done and no worker is available, Add returns ctx.Err().
//
// f must not call Add or MustAdd, but may call TryAdd.
func (p *Pool) Add(ctx context.Context, f func()) error {
	select {
	case <-p.finish:
		return ErrFinished
	default:
	}
	select {
	case <-p.finish:
		return ErrFinished
	case p.work <- f:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TryAdd is like Add, but does not block.
// The return value reports whether f was added.
func (p *Pool) TryAdd(f func()) bool {
	select {
	case p.work <- f:
		return true
	default:
		return false
	}
}

// MustAdd is like Add, but blocks indefinitely.
// If Finish has been called, MustAdd panics.
func (p *Pool) MustAdd(f func()) {
	select {
	case <-p.finish:
		panic(ErrFinished)
	case p.work <- f:
	}
}

// Finish causes pending and future calls to Add to return ErrFinished,
// then blocks until all workers have returned.
func (p *Pool) Finish() {
	select {
	case <-p.finish:
	default:
		close(p.finish)
	}
	p.wg.Wait()
}
