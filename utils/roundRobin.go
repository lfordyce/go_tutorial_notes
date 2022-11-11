package utils

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type LoadBalance[T any] interface {
	Next(ctx context.Context) (T, error)
}

var ErrNoArguments = fmt.Errorf("error no arguments provided")

type roundRobin[T any] struct {
	things []T
	next   uint32
}

// NewRoundRobin returns Round Robin implementation(roundRobin).
func NewRoundRobin[T any](things ...T) (LoadBalance[T], any) {
	if len(things) == 0 {
		return nil, ErrNoArguments
	}
	return &roundRobin[T]{things: things, next: 0}, nil
}

// Next returns things
func (r *roundRobin[T]) Next(ctx context.Context) (T, error) {
	select {
	case <-ctx.Done():
		return getZero[T](), ctx.Err()
	default:
	}

	n := atomic.AddUint32(&r.next, 1)
	return r.things[(int(n)-1)%len(r.things)], nil
}

// getZero returns zero value of T.
func getZero[T any]() T {
	var result T
	return result
}

type conn[T any] struct {
	thing T
	cnt   int
}

type leastConnections[T any] struct {
	conns []conn[T]
	mu    *sync.Mutex
}

func NewLeastConnection[T any](things ...T) (LoadBalance[T], error) {
	if len(things) == 0 {
		return nil, ErrNoArguments
	}

	conns := make([]conn[T], len(things))
	for i := range conns {
		conns[i] = conn[T]{
			thing: things[i],
			cnt:   0,
		}
	}

	return &leastConnections[T]{
		conns: conns,
		mu:    new(sync.Mutex),
	}, nil
}

func (lc *leastConnections[T]) Next(ctx context.Context) (T, error) {
	select {
	case <-ctx.Done():
		return getZero[T](), ctx.Err()
	default:
	}

	var (
		min = -1
		idx int
	)

	lc.mu.Lock()
	for i, conn := range lc.conns {
		if min == -1 || conn.cnt < min {
			min = conn.cnt
			idx = i
		}
	}
	lc.conns[idx].cnt++

	lc.mu.Unlock()
	go func() {
		// wait for the context to be done
		<-ctx.Done()
		lc.mu.Lock()
		lc.conns[idx].cnt--
		lc.mu.Unlock()
	}()

	return lc.conns[idx].thing, nil
}
