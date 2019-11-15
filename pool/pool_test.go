package pool

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {

	pool := NewPool(1)

	ctx := context.Background()
	for _, task := range []int{636, 878, 150, 904} {
		task := task
		pool.Add(ctx, func() {
			i := digits(task)
			fmt.Printf("sum of digits %d\n", i)
		})
	}
	pool.Finish()
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	//fmt.Printf("sum of digits %d\n", sum)
	return sum
}

func TestCheckStatus(t *testing.T) {

	done := make(chan struct{})
	defer close(done)

	urls := []string{
		"https://golang.org/",
		"https://medium.com/",
		"https://www.rust-lang.org/",
		"https://www.google.com",
		"https://badhost",
	}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

func checkStatus(done <-chan struct{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			select {
			case <-done:
				return
			case responses <- resp:
			}
		}
	}()
	return responses
}
