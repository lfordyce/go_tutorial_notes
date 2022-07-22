package main

import (
	"bytes"
	"fmt"
	"github.com/lfordyce/generalNotes/collections/trie"
	"github.com/lfordyce/generalNotes/concurrency"
	"github.com/lfordyce/generalNotes/interview"
	"github.com/lfordyce/generalNotes/sorting"
)

func staircase(n int) {

	var buffer bytes.Buffer
	stuff := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		stuff = append(stuff, i)
		fmt.Println(stuff)

		//buffer.Write([]byte("#"))
		buffer.WriteString("#")
		fmt.Println(buffer.String())
	}

}

type T struct {
	name string
}

func (t *T) SayHi() {
	fmt.Printf("Hi my name is %s\n", t.name)
}

func main() {

	t := &T{"Batman"}
	f := (*T).SayHi
	f(t)

	str := "racecar" // len = 7
	str2 := "aabbaa" // len 6
	fmt.Println("Is Palindrome: ", interview.IsPalindrome(str))
	fmt.Println("Is Palindrome: ", interview.IsPalindrome(str2))

	slice := sorting.GenerateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	fmt.Println("\n --- sorted --- \n\n", sorting.MergeSort(slice))

	values := []int32{1, 2, 3, 4, 5}
	fmt.Println("--- Left rotation: ", interview.RotLeft(values, 4))

	//s1 := []string{"hello", "hi", "world", "foo"}
	//s2 := []string{"hola", "hey", "bonjour", "foo", "hi"}

	s3 := interview.TwoStrings("hello", "hi")
	fmt.Println("\n -- Intersection: ", s3)

	increasing := func(a, b int) bool {
		return a <= b
	}
	decreasing := func(a, b int) bool {
		return a >= b
	}
	data := []int{31, 41, 59, 26, 41, 58}

	fmt.Println("Increasing sort array: ", sorting.SelectionSort(data, increasing))
	fmt.Println("Decreasing sort array: ", sorting.SelectionSort(data, decreasing))

	concurrency.Init()
	concurrency.HandleAsyncCalls()

	m := trie.NewTreeMap()

	cases := []struct {
		key   string
		value interface{}
	}{
		{"fish", 0},
		{"cat", 1},
		{"dog", 2},
		{"cats", 3},
		{"caterpillar", 4},
		{"cattle", 5},
		{"apple", 6},
		{"battle", 7},
		{"statistics", 8},
	}

	for _, c := range cases {
		m.Insert(c.key, c.value)
	}

	for _, c := range cases {
		if !m.Exists(c.key) {
			err := fmt.Errorf("map does not contain word: (%v)\n", c.key)
			fmt.Println(err.Error())
		}
	}

	for _, prefix := range []string{
		"app",
		"cat",
		"bat",
		"stat",
	} {
		if !m.Prefix(prefix) {
			err := fmt.Errorf("map does not contain prefix of: (%v)\n", prefix)
			fmt.Println(err.Error())
		}
	}

	for _, c := range cases {
		if val := m.Get(c.key); val != c.value {
			err := fmt.Errorf("expected key (%s) to have value (%v), got (%v)\n", c.key, c.value, val)
			fmt.Println(err.Error())
		}
	}
}
