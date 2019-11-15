package concurrency

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

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

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
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

func orDones(done, c <-chan interface{}) <-chan interface{} {
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

func TestTake(t *testing.T) {

	done := make(chan interface{})
	defer close(done)

	randFn := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFn(done, randFn), 30) {
		fmt.Println(num)
	}
}

func ProcessWithErrorGroup(ctx context.Context, input <-chan string) (<-chan string, error) {
	group, ctx := errgroup.WithContext(ctx)

	out := make(chan string)

	group.Go(func() error {
		for str := range input {
			out <- doHeavyOperation(str)
		}
		return nil
	})

	group.Go(func() error {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var err error
	inner:
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println("tick", tick)
			case <-ctx.Done():
				err = ctx.Err()
				break inner
				//return ctx.Err()
			}
		}
		return err
	})

	go func() {
		group.Wait()
		close(out)
	}()

	return out, group.Wait()
}

func Processing(ctx context.Context, input <-chan string) (<-chan string, <-chan error) {
	var wg sync.WaitGroup
	wg.Add(1)

	out := make(chan string)
	errc := make(chan error, 1)

	go func(out chan<- string) {
		for str := range input {
			//if str == "first" {
			//	errc <- errors.New("some other error")
			//}
			out <- doHeavyOperation(str)
		}
		//wg.Done()
	}(out)

	go func(errc chan<- error) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
	loop:
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println("tick", tick)
			case <-ctx.Done():
				errc <- ctx.Err()
				wg.Done()
				break loop
			}
		}
	}(errc)

	go func() {
		wg.Wait()
		close(out)
		close(errc)
	}()
	return out, errc
}

func doHeavyMockOperation(str string) string {
	return "(" + str + ")"
}

func TestPubSubChannelPSubscribe(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	input := make(chan string)
	defer close(input)

	done := make(chan struct{})
	defer close(done)

	//go func() {
	//	input <- "first"
	//	<-time.After(time.Second * 5)
	//	cancelFunc()
	//	input <- "second"
	//	<-time.After(time.Second * 5)
	//	input <- "third"
	//	<-time.After(time.Second * 5)
	//	//cancelFunc()
	//	//close(done)
	//}()

	go func() {
		//defer close(input)
		for {
			select {
			case <-time.After(time.Second * 5):
				input <- "first"
			case <-ctx.Done():
				return
			}
		}
	}()

	process, errc := Processing(ctx, input)
outer:
	for {
		select {
		case resp, ok := <-process:
			if !ok {
				//if err := <-errc; err != nil {
				//	fmt.Println(err)
				//}
				select {
				case err := <-errc:

					fmt.Println(err)
					if err == context.DeadlineExceeded {
						fmt.Println("Timeout reached")
					}

					if err == context.Canceled {
						fmt.Println()
					}
					//default:
					//	fmt.Println("error done")
					//	break
				}
				break outer
			}
			fmt.Println(resp)
		}
	}
	fmt.Println("all done")
}

func TestPSubscribeContext(t *testing.T) {

	gen := func(ctx context.Context) <-chan string {
		dst := make(chan string)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				//case dst <- strconv.Itoa(n):
				//	n++
				case <-time.After(1 * time.Second):
					dst <- strconv.Itoa(n)
					n++

				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	process, errc := Processing(ctx, gen(ctx))
outer:
	for {
		select {
		case resp, ok := <-process:
			if !ok {
				if err := <-errc; err != nil {
					fmt.Println(err)

					if err == context.DeadlineExceeded {
						fmt.Println("TIMEOUT")
					}

					if err == context.Canceled {
						fmt.Println("CANCELLATION")
					}
				}
				//select {
				//case err := <-errc:
				//	fmt.Println(err)
				//	//default:
				//	//	fmt.Println("error done")
				//	//	break
				//}
				break outer
			}
			fmt.Println(resp)
			if resp == "(5)" {
				cancel()
			}
		}
	}
	fmt.Println("all done")
}

func TestPSubscribeErrorGroupj(t *testing.T) {
	gen := func(ctx context.Context) <-chan string {
		dst := make(chan string)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				//case dst <- strconv.Itoa(n):
				//	n++
				case <-time.After(1 * time.Second):
					dst <- strconv.Itoa(n)
					n++
				}
			}
		}()
		return dst
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	strings, err := ProcessWithErrorGroup(ctx, gen(ctx))

loop:
	for {
		select {
		case resp, ok := <-strings:
			if !ok {
				if err != nil {
					fmt.Println("error reached")
					fmt.Println(err)
				}
				break loop
			}
			fmt.Println(resp)
		}
	}
	fmt.Println("all done")
}

func TestSelectWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go sendRegularHeartbeats(ctx)
	time.Sleep(time.Second * 5)
}

func sendRegularHeartbeats2(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		//select as usual
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			//give priority to a possible concurrent Done() event non-blocking way
			select {
			case <-ctx.Done():
				return
			default:
			}
			sendHeartbeat()
		}
	}
}

func sendRegularHeartbeats(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			sendHeartbeat()
		}
	}
}

func sendHeartbeat() {
	fmt.Println("Heartbeat!")
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
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

	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
