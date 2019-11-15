package executor

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

type contextKey struct{}

var executeKey = contextKey{}

var (
	ID     = 0
	IdLock = new(sync.Mutex)
)

func NextID() int {
	IdLock.Lock()
	defer IdLock.Unlock()
	ID++
	n := ID
	return n
}

type PrintTask struct {
	ID int
}

func (t PrintTask) Run(c context.Context) error {
	gid := getGID()
	fmt.Printf("Task: %d, GO:%d : IN\n", t.ID, gid)
	defer func() {
		fmt.Printf("Task: %d, GO:%d : OUT\n", t.ID, gid)
	}()
	//<-time.After(time.Duration(rand.Intn(10)) * time.Second)
	<-time.After(6 * time.Second)
	return nil
}

func TaskFromContext(ctx context.Context) Task {
	value := ctx.Value(executeKey)
	if task, ok := value.(Task); ok {
		return task
	}
	return nil
}

func NewPrintTask() *PrintTask {
	t := &PrintTask{NextID()}
	return t
}

func GenerateTask(num int, x Executor) {
	for i := 0; i < num; i++ {
		pt := NewPrintTask()
		x.Send(pt)
	}
}

// getGID returns current goroutine ID. **DO NOT USE** Only for debug
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func TestNewSpawnExecutor(t *testing.T) {
	x := NewSpawnExecutor()
	x.Start()
	go GenerateTask(5, x)
	go func() {
		<-time.After(20 * time.Second)
		x.Stop()
	}()
	x.Wait()
}

type TestTask struct {
	err      error
	delay    time.Duration
	executed bool
}

func (t *TestTask) Run(ctx context.Context) error {
	<-time.After(t.delay)
	t.executed = true
	return t.err
}

func TestSpawnExecutor(t *testing.T) {
	q := make(chan Task)
	e := NewSpawnExecutorForQueue(q)
	t.Log("Starting executor")
	e.Start()
	go func() {
		<-time.After(10 * time.Second)
		t.Log("Stopping executor")
		e.Stop()
	}()

	ts := &TestTask{}
	// TODO
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				// ts := &TestTask{}
				t.Log("sending task to task queue")
				e.Send(ts)
			}
		}
	}()

	// ts := &TestTask{}
	// e.Send(ts)
	e.Wait()
	if !ts.executed {
		t.Errorf("Task not executed, got: %v, want: %v", ts.executed, true)
	}

}
