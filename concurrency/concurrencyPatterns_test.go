package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	// GIVEN
	input := make(chan string)
	defer close(input)

	done := make(chan bool)
	defer close(done)

	go func() {
		input <- "hello world"
		done <- true
	}()

	// WHEN
	output := Process(input)
	<-done // blocks until the input write routine is finished

	// THEN
	expected := "(hello world)"
	found := <-output // blocks until the output has contents

	if found != expected {
		t.Errorf("Expected %s, found %s", expected, found)
	}
}

func TestAnalyzeJobs(t *testing.T) {
	jobs := make(chan *Job, 100)    // Buffered channel
	results := make(chan *Job, 100) // Buffered channel

	// Start consumers:
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go comsume(i, jobs, results)
	}
	// start producer
	go produce(jobs)

	//start analyzer
	wg2.Add(1)
	go analyze(results)

	wg.Wait()

	// All jobs processed, no more values will be send on result;
	close(results)

	// Wait analyzer to analyze all results
	wg2.Wait()
}

func TestPubSub(t *testing.T) {
	ch := make(chan int)
	clients := 4
	// make it buffered, so all clients can fail without hanging
	notifyCh := make(chan struct{}, clients)
	go producer(100, ch, notifyCh)

	var wg sync.WaitGroup
	wg.Add(clients)
	for i := 0; i < clients; i++ {
		go func() {
			consumer(ch, notifyCh)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestConcurrentPrime(t *testing.T) {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func TestNonBlocking(t *testing.T) {
	var wg sync.WaitGroup

	i := 0
	fmt.Println("Starting...")
	for i <= 3 {
		fmt.Println("Loop: ", i)
		go long(&wg, i)
		time.Sleep(1 * time.Second)
		i = i + 1
	}
	wg.Wait()
	fmt.Println("Done...")
}

func TestNoBlockingWithChan(t *testing.T) {
	ok := make(chan bool, 1)
	done := make(chan bool)

	i := 0
	j := 0
	fmt.Println("Starting...")

	for i <= 3 {
		fmt.Println("Loop: ", i)
		go longChan(ok, i)
		time.Sleep(1 * time.Second)
		i = i + 1
	}

	go func() {
		for {
			select {
			case _ = <-ok:
				j++
				if j == 4 {
					done <- true
					return
				}
			}
		}
	}()

	<-done
	fmt.Println("Done...")
}

func TestParallelize(t *testing.T) {
	var functions []func()
	for i := 1; i <= 100; i++ {
		function := func(i int) func() {
			return func() {
				fmt.Println(i)
			}
		}(i)
		functions = append(functions, function)
	}

	Parallelize(functions...)
	fmt.Println("DONE")
}

func TestPerformTasks(t *testing.T) {
	DoSomeWork()
}

func TestStateMonitor(t *testing.T) {
	// Create our input and output channels.

	pending := make(chan *Resource)
	complete := make(chan *Resource)

	// Launch the StateMonitor.
	status := StateMonitor(statusInterval)

	// Launch some Poller goroutines.
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}
	//go Poller(pending, complete, status)

	// Send some Resources to the pending queue.
	go func() {
		for _, url := range urls2 {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}
}

func perform(work int) {
	fmt.Println(work)
}

var sendWork = make(chan int)

func receiveWork() int {
	return <-sendWork
}

func worker(ready chan struct{}, work chan int, wg *sync.WaitGroup) {
	for range ready {
		w, ok := <-work
		if !ok {
			break
		}
		perform(w)
	}
	wg.Done()
}

func queue(ready chan struct{}, work chan int, stop chan struct{}) {
	defer close(ready)
	defer close(work)
	for {
		select {
		case ready <- struct{}{}:
			work <- receiveWork()
		case <-stop:
			return
		}
	}
}

func TestSendRecieve(t *testing.T) {
	work := make(chan int)
	ready := make(chan struct{})
	stop := make(chan struct{})

	go queue(ready, work, stop)

	const parallel = 4

	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go worker(ready, work, &wg)
	}

	for i := 0; i < 10; i++ {
		sendWork <- i
	}
	close(stop)
	wg.Wait()
}

// wrap a timeout channel in a generic interface channel
//func makeDefaultTimeoutChan() <-chan interface{} {
//	channel := make(chan interface{})
//	go func() {
//		<-time.After(30 * time.Second)
//		channel <- struct{}{}
//	}()
//	return channel
//}

// usage
//func main() {
//	resultChannel := doOtherThingReturningAsync()
//	cancel := makeDefaultTimeoutChan()
//	select {
//	case <-cancel:
//		fmt.Println("cancelled!")
//	case results := <-resultChannel:
//		fmt.Printf("got result: %#v\n", results)
//	}
//}
