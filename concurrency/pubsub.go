package concurrency

import (
	"fmt"
	"time"
)

func consumer(in chan int, notifyCh chan struct{}) {
	fmt.Printf("Start consumer\n")
	for i := range in {
		<-time.After(100 * time.Millisecond)
		if i == 42 {
			fmt.Printf("%d fails\n", i)
			notifyCh <- struct{}{}
			return
		} else {
			fmt.Printf("%d\n", i)
		}

	}
	fmt.Printf("Consumer stopped working\n")
}

func producer(N int, out chan int, notifyCh chan struct{}) {
	for i := 0; i < N; i++ {
		select {
		case out <- i:
		case <-notifyCh:
			close(out)
			return
		}
	}
	close(out)
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

type control int

const (
	sleep control = iota
	die           // receiver will close the control chan in response to die, to ack.
)

func (cmd control) String() string {
	switch cmd {
	case sleep:
		return "sleep"
	case die:
		return "die"
	}
	return fmt.Sprintf("%d", cmd)
}

func ProduceTo(writechan chan<- int, ctrl chan control, done chan bool) {
	var product int
	go func() {
		for {
			select {
			case writechan <- product:
				fmt.Printf("Producer produced %v\n", product)
				product++
			case cmd := <-ctrl:
				fmt.Printf("Producer got control cmd: %v\n", cmd)
				switch cmd {
				case sleep:
					fmt.Printf("Producer sleeping 2 sec.\n")
					time.Sleep(2000 * time.Millisecond)
				case die:
					fmt.Printf("Producer dies.\n")
					close(done)
					return
				}
			}
		}
	}()
}

func ConsumeFrom(readchan <-chan int, ctrl chan control, done chan bool) {
	go func() {
		var product int
		for {
			select {
			case product = <-readchan:
				fmt.Printf("Consumer consumed %v\n", product)
			case cmd := <-ctrl:
				fmt.Printf("Consumer got control cmd: %v\n", cmd)
				switch cmd {
				case sleep:
					fmt.Printf("Consumer sleeping 2 sec.\n")
					time.Sleep(2000 * time.Millisecond)
				case die:
					fmt.Printf("Consumer dies.\n")
					close(done)
					return
				}

			}
		}
	}()
}
