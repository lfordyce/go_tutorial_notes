package main

import (
	"fmt"
	"github.com/lfordyce/generalNotes/concurrency"
	"github.com/lfordyce/generalNotes/interview"
	"github.com/lfordyce/generalNotes/sorting"
)

func main() {

	str := "racecar" // len = 7
	str2 := "aabbaa" // len 6
	fmt.Println("Is Palindrome: ", interview.IsPalindrome(str))
	fmt.Println("Is Palindrome: ", interview.IsPalindrome(str2))

	slice := sorting.GenerateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	fmt.Println("\n --- sorted --- \n\n", sorting.MergeSort(slice))

	values := []int32{1, 2, 3, 4, 5}
	fmt.Println("--- Left rotation: ",  interview.RotLeft(values, 4))

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

	fmt.Println("Increasing sort array: ",sorting.SelectionSort(data, increasing))
	fmt.Println("Decreasing sort array: ",sorting.SelectionSort(data, decreasing))

	concurrency.Init()
}
