package concurrency

import (
	"fmt"
	"math/big"
	"sync"
)

var (
	once sync.Once
)

func Mux(channels []chan big.Int) chan big.Int {
	n := len(channels)

	ch := make(chan big.Int, n)

	// Make one go per channel.
	for _, c := range channels {
		go func(c <-chan big.Int) {
			for x := range c {
				ch <- x
			}
			n -= 1
			//close output if all close now.
			if n == 0 {
				close(ch)
			}
		}(c)
	}

	return ch
}

func MuxSync(consumers []chan []byte) {

	logCh := make(chan []byte, 0)

	//group := sync.WaitGroup{}

	once.Do(func() {
		go func() {
			for v := range logCh {
				for _, cons := range consumers {
					cons <- v
				}
			}
		}()
	})

}

func fromTo(f, t int) chan big.Int {
	ch := make(chan big.Int)

	go func() {
		for i := f; i < t; i++ {
			fmt.Println("Feed: ", i)
			ch <- *big.NewInt(int64(i))
		}
		close(ch)
	}()
	return ch
}
