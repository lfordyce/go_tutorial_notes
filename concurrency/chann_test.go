package concurrency

import "testing"

func TestChann(t *testing.T) {
	in, out := MakeChan[int](10)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	for c := range out {
		println(c)
	}
}
