package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg, wg2 sync.WaitGroup

type Job struct {
	ID     int
	Work   string
	Result string
}

func produce(jobs chan<- *Job) {
	// Generate jobs
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++
		jobs <- &Job{ID: id, Work: fmt.Sprintf("%c", c)}
	}
	close(jobs)
}

func comsume(id int, jobs <-chan *Job, results chan<- *Job) {
	defer wg.Done()
	for job := range jobs {
		sleepMs := rand.Intn(1000)
		fmt.Printf("worker #%d received: '%s', sleep %dms\n", id, job.Work, sleepMs)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		job.Result = job.Work + fmt.Sprintf("-%dms", sleepMs)
		results <- job
	}
}

func analyze(results <-chan *Job) {
	defer wg2.Done()
	for job := range results {
		fmt.Printf("result: %s\n", job.Result)
	}
}

func long(wg *sync.WaitGroup, i int) {
	fmt.Println("Inside long: ", i)
	time.Sleep(3 * time.Second)
	fmt.Println("Done with loop: ", i)
	wg.Done()
}

func longChan(c chan bool, i int) {
	fmt.Println("Inside long: ", i)
	time.Sleep(3 * time.Second)
	fmt.Println("Done with loop: ", i)
	c <- true
}

// ASYNC TEST:
type TaskFunction func() interface{}

// PerformTasks is a function which will be called by the client to perform
// multiple task concurrently.
// Input:
// tasks: the slice with functions (type TaskFunction)
// done:  the channel to trigger the end of task processing and return
// Output: the channel with results
func PerformTasks(ctx context.Context, tasks []TaskFunction) chan interface{} {

	// Create a worker for each incoming task
	workers := make([]chan interface{}, 0, len(tasks))

	for _, task := range tasks {
		resultChannel := newWorker(ctx, task)
		workers = append(workers, resultChannel)
	}

	// Merge results from all workers
	out := merge(ctx, workers)
	return out
}

// newWorker: Commonly “select” statement is used for non-blocking reading from
// channels when it is not necessary to wait the result.
// It can be achieved using “default” statement.
// But it can be used to read results from several channels.
// In our example firstly it reads from “done” channel.
// It is indicating that we have to leave worker before task is actually done.
// Secondly we execute the task and send the result to output channel which is accessible from the worker (goroutine).
func newWorker(ctx context.Context, task TaskFunction) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)

		select {
		case <-ctx.Done():
			// Received a signal to abandon further processing
			return
		case out <- task():
			// Got some result
		}
	}()

	return out
}

func merge(ctx context.Context, workers []chan interface{}) chan interface{} {
	// Merged channel with results
	out := make(chan interface{})

	// Synchronization over channels: do not close "out" before all tasks are completed
	var wg sync.WaitGroup

	// Define function which waits the result from worker channel
	// and sends this result to the merged channel.
	// Then it decreases the counter of running tasks via wg.Done().
	output := func(c <-chan interface{}) {
		defer wg.Done()
		for result := range c {
			select {
			case <-ctx.Done():
				// Received a signal to abandon further processing
				return
			case out <- result:
				// some message or nothing
			}
		}
	}

	wg.Add(len(workers))
	for _, workerChannel := range workers {
		go output(workerChannel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
