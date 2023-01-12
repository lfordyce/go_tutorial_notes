package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPrintConditionChan(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := NewCond(&sync.Mutex{})
	cond.L.Lock()

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // simulate a workload

			// Changes must be made when
			// cond.L is locked.
			cond.L.Lock()
			values[i] = string(rune('a' + i))

			// Notify when cond.L lock is acquired.
			cond.Broadcast()
			cond.L.Unlock()

			// "cond.Broadcast()" can also be put
			// here, when cond.L lock is released.
			//cond.Broadcast()
		}(i)
	}

	// This function must be called when
	// cond.L is locked.
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}
	for !checkCondition() {
		// Must be called when cond.L is locked.
		cond.Wait()
	}
	cond.L.Unlock()
}

var n int64

func dostuff() {
	atomic.AddInt64(&n, 1)
}

type looper struct {
	pause  chan struct{}
	paused sync.WaitGroup
	resume chan struct{}
}

func (l *looper) loop() {
	for {
		select {
		case <-l.pause:
			l.paused.Done()
			<-l.resume
		default:
			dostuff()
		}
	}
}

func (l *looper) whilePaused(fn func()) {
	l.paused.Add(32)
	l.resume = make(chan struct{})
	close(l.pause)
	l.paused.Wait()
	fn()
	l.pause = make(chan struct{})
	close(l.resume)
}

func TestLooper(t *testing.T) {
	l := &looper{
		pause: make(chan struct{}),
	}
	var init sync.WaitGroup
	init.Add(32)
	for i := 0; i < 32; i++ {
		go func() {
			init.Done()
			l.loop()
		}()
	}
	init.Wait()
	for i := 0; i < 100; i++ {
		l.whilePaused(func() { fmt.Printf("%d ", i) })
	}
	fmt.Printf("\n%d\n", atomic.LoadInt64(&n))
}

func TestChannelOrDonePattern(t *testing.T) {

	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}

type startGoroutineFn func(done <-chan struct{}, pulseInterval time.Duration) (heartbeat <-chan struct{})

func TestSteward(t *testing.T) {

	var or func(channels ...<-chan struct{}) <-chan struct{}

	or = func(channels ...<-chan struct{}) <-chan struct{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan struct{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	newSteward := func(timeout time.Duration, fn startGoroutineFn) startGoroutineFn {
		return func(done <-chan struct{}, pulseInterval time.Duration) <-chan struct{} {
			heartbeat := make(chan struct{})
			go func() {
				defer close(heartbeat)

				var wardDone chan struct{}
				var wardHeartbeat <-chan struct{}

				startWard := func() {
					wardDone = make(chan struct{})
					wardHeartbeat = fn(or(wardDone, done), timeout/2)
				}

				startWard()
				pulse := time.Tick(pulseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat:
							continue monitorLoop
						case <-timeoutSignal:
							t.Log("steward: unhealthy; restarting")
							//fmt.Println("steward: unhealthy; restarting")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}

			}()
			return heartbeat
		}
	}

	doWork := func(done <-chan struct{}, _ time.Duration) <-chan struct{} {
		t.Log("ward: Hello I'm irresponsible!")
		//fmt.Println("ward: Hello I'm irresponsible!")
		go func() {
			<-done
			t.Log("ward: I am halting")
			//fmt.Println("ward: I am halting")
		}()
		return nil
	}

	doWorkWithSteward := newSteward(4*time.Second, doWork)

	done := make(chan struct{})
	time.AfterFunc(9*time.Second, func() {
		t.Log("main: halting steward and ward.")
		//fmt.Println("main: halting steward and ward.")
		close(done)
	})
	for range doWorkWithSteward(done, 4*time.Second) {
	}
	t.Log("DONE")
	//fmt.Println("done")

}
