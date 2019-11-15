package mapreducd

import (
	"fmt"
	"testing"
)

func TestMapreduce(t *testing.T) {
	list = []string{"a", "abc", "ef", "g", "hijk"}
	res := reduce(måp(list, length), sum)
	fmt.Println(res)
}

func TestConcurrentMapReduce(t *testing.T) {
	size := 10
	text1 := make(chan string, size)
	text2 := make(chan string, size)
	text3 := make(chan string, size)
	map1 := make(chan map[string]int, size)
	map2 := make(chan map[string]int, size)
	map3 := make(chan map[string]int, size)
	reduce1 := make(chan int, size)
	reduce2 := make(chan int, size)
	avg1 := make(chan float32, size)
	avg2 := make(chan float32, size)

	// Start all workers in separate goroutines, chained together via channels.
	go inputReader([3]chan<- string{text1, text2, text3})
	go mapper(text1, map1)
	go mapper(text2, map2)
	go mapper(text3, map3)
	go shuffler([]<-chan map[string]int{map1, map2, map3}, [2]chan<- int{reduce1, reduce2})
	go reducer(reduce1, avg1)
	go reducer(reduce2, avg2)

	// The outputWriter runs in the main thread.
	outputWriter([]<-chan float32{avg1, avg2})
}
