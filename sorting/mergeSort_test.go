package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {

	slice := generateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	sort := MergeSort(slice)
	fmt.Println("\n --- sorted --- \n\n",sort)

	for i := 0; i < len(sort) - 1; i++ {
		if sort[i] > sort[i + 1] {
			t.Error("Merge sort failed")
		}
	}
}

func TestMergeSortOffLength(t *testing.T) {
	numbers := [7]int{9, 3, 6, 8, 13, 5, 6}
	slice := numbers[:]
	sort := MergeSort(slice)

	for i := 0; i < len(sort) - 1; i++ {
		if sort[i] > sort[i + 1] {
			t.Error("Merge sort failed")
		}
	}
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}