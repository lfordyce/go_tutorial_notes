package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	data := Data{readRequest: make(chan chan int)}
	go data.Run()
	time.Sleep(time.Second)
	secret := data.Get()
	fmt.Println(secret)
}
