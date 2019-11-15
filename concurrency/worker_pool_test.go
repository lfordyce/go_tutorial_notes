package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	tasks := []*Task{
		NewTask(func() error {
			fmt.Println("task 1")
			return nil
		}),
		NewTask(func() error {
			fmt.Println("task 2")
			time.Sleep(time.Second * 4)
			return nil
		}),
		NewTask(func() error {
			fmt.Println("task 3")
			return nil
		}),
		NewTask(func() error {
			fmt.Println("task 4")
			return nil
		}),
		NewTask(func() error {
			fmt.Println("task 5")
			return nil
		}),
	}

	p := NewPool(tasks, 1)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			t.Log(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			t.Log("Too many errors.")
			break
		}
	}
}
