package concurrency

import (
	"fmt"
	"sync"
)

func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatFunc(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func repeater(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func typicalDone(done <-chan interface{}, input <-chan interface{}) {
	//outValues := make(chan interface{})
	go func() {
	loop:
		for {
			select {
			case <-done:
				break loop
			case maybeVal, ok := <-input:
				if !ok {
					return // or maybe break
				}
				// do something with val
				fmt.Println(maybeVal)
			}
		}
	}()
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func returnOrDone() func(done, c <-chan interface{}) <-chan interface{} {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
	return orDone
}

func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range orDone(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func returnTee() func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {

	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}
	return tee
}

func Blocking(dst chan<- string, s string) {
	dst <- s
}
func NonBlocking(dst chan<- string, s string) {
	select {
	case dst <- s:
	default:
	}
}

func TeeChan(src <-chan string, fn func(chan<- string, string)) func(...chan<- string) {
	return func(destinations ...chan<- string) {
		go func() {
			for s := range src {
				for _, d := range destinations {
					fn(d, s)
				}
			}
			for _, d := range destinations {
				close(d)
			}
		}()
	}
}

func TeeChan2(src <-chan string, fn func(chan<- string, string)) (<-chan string, <-chan string) {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		defer close(out1)
		defer close(out2)
		for s := range src {
			fn(out1, s)
			fn(out2, s)
		}
	}()
	return out1, out2
}

func TeeCh3(source chan string) (tee1, tee2 chan string) {

	tee1 = make(chan string)
	tee2 = make(chan string)

	go func() {

		defer func() {
			close(tee1)
			close(tee2)
		}()

		for s := range source {
			select {
			case tee1 <- s:
			default:
				fmt.Println("blocked sending to tee1")
				// log.Println("blocked sending to tee1")
			}
			select {
			case tee2 <- s:
			default:
				fmt.Println("blocked sending to tee2")
				// log.Println("blocked sending to tee2")
			}
		}
	}()
	return tee1, tee2
}
