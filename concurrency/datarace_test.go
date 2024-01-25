package concurrency

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	data := Data{readRequest: make(chan chan int)}
	go data.Run()
	time.Sleep(time.Second)
	secret := data.Get()
	fmt.Printf("Fisrt Secret : %d\n", secret)
	time.Sleep(time.Second)
	second := data.Get()
	fmt.Printf("Second Secret : %d\n", second)
	time.Sleep(time.Second)
	third := data.Get()
	fmt.Printf("Third Secret : %d\n", third)
}

// You can send a channel over a channel to protect the single access of a selected value.
// Only a receiver will do the work and the others will leave when they received the closing signal.
func TestChanOfChans(t *testing.T) {
	//this is overkill, it's the default on go1.5+
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	st := make(chan string)
	//send the st over the channel
	owner := make(chan (<-chan string))
	//send the channel over owner, only one channel will receive it
	go func() {
		owner <- st
		//only one goroutine will get the st channel
		//the others will get an empty value and a flag that indicates
		//that the channel is closed
		close(owner)
	}()
	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Println("me", x)
			//only one goroutine will receive from st
			fmt.Println("waiting receiver", x)
			st, ok := <-owner
			if !ok {
				fmt.Println("channel is closed", x)
				return
			}
			fmt.Println("ok here we go", x)
			for val := range st {
				seconds := rand.Intn(3)
				sleepTime := time.Second * time.Duration(seconds)
				time.Sleep(sleepTime)
				fmt.Println("I", x, "fell asleep for ", seconds, "seconds", "and received ", val)
			}
			fmt.Println("done with the channel")
		}(i)
	}
	for i := 0; i < 10; i++ {
		st <- fmt.Sprintf("n: %d", i)
	}
	//you need to close to say the goroutine(s) that is listening the channel
	//"there are no new messages, work is over"
	close(st)
	fmt.Println("waiting..")
	wg.Wait()
	fmt.Println("all done")
}
