package concurrency

import (
	"fmt"
	"sync"
)

var accountBalance = 0
var mutex = &sync.Mutex{}

func updateBalance(amt int) {
	mutex.Lock()
	defer mutex.Unlock()
	accountBalance += amt
}

func reportAndExit(msg string) {
	fmt.Println(msg)
}

type bankOp struct {
	howMuch int
	confirm chan int
}

var accountBalanceChan = 0
var bankRequests chan *bankOp

func updateBalanceChan(amt int) int {
	update := &bankOp{howMuch: amt, confirm: make(chan int)}
	bankRequests <- update
	newBalance := <-update.confirm
	return newBalance
}

func logBalance(current int) {
	fmt.Println(current)
}

func reportAndExitChan(msg string) {
	fmt.Println(msg)
}
