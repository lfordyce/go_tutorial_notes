package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Init concurrency patterns
func Init() {
	// create the channel
	//var c chan int
	//c = make(chan int)

	// alternatively
	//c := make(chan int)
	// Sending on a channel
	//c <- 1

	// Receiving from a channel
	// the "arrow" indicates the direction of data flow
	//value = <- c

	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() { // launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can ben any suitable value.
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Process function that reads from an input and executes the operation,
// another routine that waits for the operation to complete, and return immediately the output channel
func Process(input <-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(1)

	output := make(chan string)

	go func() {
		for str := range input {
			output <- doHeavyOperation(str)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func doHeavyOperation(str string) string {
	return "(" + str + ")"
}
