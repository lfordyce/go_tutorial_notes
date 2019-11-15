package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestBankChan(t *testing.T) {
	bankRequests = make(chan *bankOp, 8)

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case request := <-bankRequests:
				accountBalanceChan += request.howMuch
				request.confirm <- accountBalanceChan
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 8; i++ {
			newBalance := updateBalanceChan(1)
			logBalance(newBalance)
			runtime.Gosched() // be nice --yield
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			newBalance := updateBalanceChan(-1)
			logBalance(newBalance)
			runtime.Gosched() // be nice --yield
		}
	}()

	wg.Wait()
	fmt.Println("final balance: ", accountBalanceChan)
}
