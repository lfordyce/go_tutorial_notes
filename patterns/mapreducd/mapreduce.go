package mapreducd

import (
	"fmt"
	"sync"
)

var list []string

func length(s string) int {
	return len(s)
}

func måp(list []string, fn func(string) int) []int {
	res := make([]int, len(list))
	for i, elem := range list {
		res[i] = fn(elem)
	}
	return res
}

func reduce(list []int, fn func(int, int) int) (res int) {
	for _, elem := range list {
		res = fn(res, elem)
	}
	return res
}

func sum(a, b int) int {
	return a + b
}

// mapper receives a channel of strings and counts the occurrence of each unique word read from this channel. It sends the resulting map to the output channel.
func mapper(in <-chan string, out chan<- map[string]int) {
	count := map[string]int{}
	for word := range in {
		count[word] = count[word] + 1
	}
	out <- count
	close(out)
}

// reducer receives a channel of ints and adds up all ints until the channel is closed. Then it divides through the number of received ints to calculate the average.
func reducer(in <-chan int, out chan<- float32) {
	sum, count := 0, 0
	for n := range in {
		sum += n
		count++
	}
	out <- float32(sum) / float32(count)
	close(out)
}

// inputDistributor receives three output channels and sends each of them some input.
func inputReader(out [3]chan<- string) {
	// "Read" some input.
	input := [][]string{
		{"noun", "verb", "verb", "noun", "noun"},
		{"verb", "verb", "verb", "noun", "noun", "verb"},
		{"noun", "noun", "verb", "noun"},
	}

	for i := range out {
		go func(ch chan<- string, word []string) {
			for _, w := range word {
				ch <- w
			}
			close(ch)
		}(out[i], input[i])
	}
}

// shuffler gets a list of input channels containing key/value pairs like
// "noun: 5, verb: 4". For each "noun" key, it sends the corresponding value
// to out[0], and for each "verb" key to out[1].
// The input channles are multiplexed into one, based on the `merge` function
// from the [Pipelines article](https://blog.golang.org/pipelines) of the
// Go Blog.
func shuffler(in []<-chan map[string]int, out [2]chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {
		go func(c <-chan map[string]int) {
			for m := range c {
				nc, ok := m["noun"]
				if ok {
					out[0] <- nc
				}
				vc, ok := m["verb"]
				if ok {
					out[1] <- vc
				}
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out[0])
		close(out[1])
	}()
}

// outputWriter starts a goroutine for each input channel and writes out
// the averages that it receives from each channel.
func outputWriter(in []<-chan float32) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	// `out[0]` contains the nouns, `out[1]` the verbs.
	name := []string{"noun", "verb"}
	for i := 0; i < len(in); i++ {
		go func(n int, c <-chan float32) {
			for avg := range c {
				fmt.Printf("Average number of %ss per input text: %f\n", name[n], avg)
			}
			wg.Done()
		}(i, in[i])
	}
	wg.Wait()
}
