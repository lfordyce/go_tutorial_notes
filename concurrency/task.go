package concurrency

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func DoSomeWork() {

	// Define three slow functions which will be runnning concurrently
	slowFunction := func() interface{} {
		time.Sleep(time.Second * 10)
		fmt.Println("slow function")
		return 2
	}

	verySlowFunction := func() interface{} {
		time.Sleep(time.Second * 5)
		fmt.Println("very slow function")
		return "I'm ready"
	}

	faskFunction := func() interface{} {
		time.Sleep(time.Nanosecond * 5)
		fmt.Println("fast function")
		return "I'm fast!!"
	}

	averageFunction := func() interface{} {
		time.Sleep(time.Millisecond * 25)
		fmt.Println("average speed")
		return 5
	}

	// One function returns an error
	errorFunction := func() interface{} {
		time.Sleep(time.Second * 11)
		fmt.Println("function with an error")
		return errors.New("error in function")
	}

	tasks := []TaskFunction{slowFunction, verySlowFunction, errorFunction, faskFunction, averageFunction}
	//tasks := []TaskFunction{slowFunction, verySlowFunction}

	// Use context to cancel goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultChannel := PerformTasks(ctx, tasks)

	// Print value from first goroutine and cancel others
	for result := range resultChannel {
		switch result.(type) {
		case error:
			fmt.Println("Received error")
			cancel()
			return
		case string:
			fmt.Println("Here is a string:", result.(string))
		case int:
			fmt.Println("Here is an integer:", result.(int))
		default:
			fmt.Println("Some unknown type ")
		}
	}
}
