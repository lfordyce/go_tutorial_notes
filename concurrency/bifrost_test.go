package concurrency

import (
	"fmt"
	"github.com/serdmanczyk/bifrost"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBifrostDispatcher(t *testing.T) {

	const iter int64 = 10
	const start int64 = 10
	const numWorkers int = 10
	const numJobs int64 = 20

	var wg sync.WaitGroup

	dispatcher := bifrost.NewWorkerDispatcher(
		bifrost.Workers(4),
		bifrost.JobExpiry(time.Second),
	)

	a := start
	jobfunc := func() (err error) {
		atomic.AddInt64(&a, iter)
		wg.Done()
		return
	}

	wg.Add(1)
	job := dispatcher.QueueFunc(jobfunc)
	status := job.Status()
	fmt.Println(status)

	<-job.Done()
	<-job.Done()
	status = job.Status()
	fmt.Println(status)

	wg.Add(1)
	job = dispatcher.QueueFunc(jobfunc)
	status = job.Status()
	fmt.Println(status)

	for i := 0; int64(i) < numJobs-2; i++ {
		wg.Add(1)
		dispatcher.Queue(bifrost.JobRunnerFunc(jobfunc))
	}
	wg.Wait()
	dispatcher.Stop()

	expectedValue := (start + numJobs*iter)
	if a != expectedValue {
		t.Errorf("Expected final a value %d, got %d", expectedValue, a)
	}

	status = job.Status()
	fmt.Println(status)
}
