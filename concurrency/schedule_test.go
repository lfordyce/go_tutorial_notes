package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	t1 := shcedule(oneSec, time.Second)
	t2 := shcedule(twoSec, 2*time.Second)

	time.Sleep(time.Second * 10)
	t1.Stop()
	t2.Stop()
}

type taskWithChan struct {
}

func (t taskWithChan) doStuff() {
	queue := make(chan struct {
		string
		int
	})
	go sendPair(queue)
	pair := <-queue
	fmt.Println(pair.string, pair.int)
}

func sendPair(queue chan struct {
	string
	int
}) {
	queue <- struct {
		string
		int
	}{string: "http:...", int: 3}
}

func TestQueue(t *testing.T) {
	q := taskWithChan{}
	q.doStuff()
}

func _TestHeartBeat(t *testing.T) {
	done := make(chan bool)
	heartbeatA := createHeartbeat(500 * time.Millisecond)
	heartbeatB := createHeartbeat(1 * time.Second)
	s := &server{
		done:       done,
		heartbeatA: *heartbeatA,
		heartbeatB: *heartbeatB,
		reset:      make(chan time.Duration),
	}
	go s.listener()
	<-done
}

func TestNewChanHolder(t *testing.T) {
	for _, h := range []TimeHolder{
		NewChanHolder(),
	} {
		h.Set(1 * time.Millisecond)
		if got, want := h.Get(), 1*time.Millisecond; got != want {
			t.Errorf("Get() got %d, want %d", got, want)
		}

		if ch, ok := h.(holder); ok {
			ch.Set(0)
			done := make(chan struct{})
			go func() {
				defer close(done)
				ch.Get()
				//t.Log(get.String())
			}()
			select {
			case <-time.After(500 * time.Millisecond):
				s := h.(holder)
				s.Set(time.Second * 2)
			case <-done:
				t.Errorf("Get() retured, expected it to block")
			}
			get := ch.Get()
			t.Log(get.String())
			ch.Close()
		}
	}
}

type server struct {
	done       chan bool
	reset      chan time.Duration
	heartbeatA heartbeat
	heartbeatB heartbeat
}

func NewServer(heartbeatA heartbeat, heartbeatB heartbeat) TimeHolder {
	h := server{
		done:       make(chan bool),
		reset:      make(chan time.Duration),
		heartbeatA: heartbeatA,
		heartbeatB: heartbeatB,
	}
	go h.listener()
	return h
}

func (s *server) listener() {
	start := time.Now()
	tickACount := 0
	fmt.Println("Elapsed: 0")
	for {
		select {
		case <-s.heartbeatA.ticker.C:
			elapsed := time.Since(start)
			fmt.Println("Elapsed: ", elapsed, " Heartbeat A")
			tickACount++
			if tickACount == 5 {
				s.done <- true
			}
		case <-s.heartbeatB.ticker.C:
			s.heartbeatA.restHeartbeat(500 * time.Millisecond)
			elapsed := time.Since(start)
			fmt.Println("Elapsed: ", elapsed, " Heartbeat B - Going to reset heartbeat A")
		case resetPeriod := <-s.reset:
			s.heartbeatA.restHeartbeat(resetPeriod)
			fmt.Println("heartbeat period reset to: ", resetPeriod)
		}
	}
}

func (s server) Get() time.Duration {
	return <-s.reset
}

func (s server) Set(duration time.Duration) {
	s.reset <- duration
}

type holder struct {
	resetPeriod chan time.Duration
	inputPeriod chan time.Duration
	closeCh     chan struct{}
}

func NewChanHolder() TimeHolder {
	h := holder{
		resetPeriod: make(chan time.Duration),
		inputPeriod: make(chan time.Duration),
		closeCh:     make(chan struct{}),
	}
	go h.mux()
	return h
}

func (h *holder) mux() {
	var period time.Duration
	for {
		if period == 0 {
			select {
			case <-h.closeCh:
				close(h.resetPeriod)
				close(h.inputPeriod)
				return
			case period = <-h.inputPeriod:
				continue
			}
		}
		select {
		case period = <-h.inputPeriod:
		case h.resetPeriod <- period:
		case <-h.closeCh:
			close(h.resetPeriod)
			close(h.inputPeriod)
			return
		}
	}
}

func (h holder) Get() time.Duration {
	return <-h.resetPeriod
}

func (h holder) Set(p time.Duration) {
	h.inputPeriod <- p
}

func (h holder) Close() {
	close(h.closeCh)
}

//func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(filepath.Base(caller.FullPath()))
//}

//func MyCaller2(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(filepath.Base(caller.FullPath()))
//}
