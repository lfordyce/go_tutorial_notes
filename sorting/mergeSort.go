package sorting

import (
	"math/rand"
	"time"
)

const NADA int = -1

func DeepCopy(vals []int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return tmp
}

// GenerateSlice Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	mid := (len(slice)) / 2

	// slice expressions: input[low:high]
	// numbers := [5]int{1, 2, 3, 4, 5}
	// numbers[1:3] -> {2, 3}
	// numbers[1:] -> {2, 3, 4, 5}
	// numbers[:4] -> {1, 2, 3, 4}
	// numbers[2:4] -> {3, 4}
	// slice := int[]{9, 3, 6, 8, 13, 5, 6}
	// slice[mid:] -> [8, 13, 5, 6]
	// slice[:mid] -> [9, 3, 6]
	low := slice[mid:]
	high := slice[:mid]

	//sort := MergeSort(low)

	//return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
	return Merge(MergeSort(low), MergeSort(high))
}

// Merge Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
