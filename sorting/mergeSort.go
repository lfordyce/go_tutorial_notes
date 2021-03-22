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

	low := slice[:mid]
	high := slice[mid:]

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


func MergeSortAlt(src []int64) {
	if len(src) < 2 {
		return
	}

	mid := len(src) / 2
	// Create new input slices since MergeSort doesn't return a new array,
	// but overwrites the input.
	left := make([]int64, mid)
	right := make([]int64, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	MergeSortAlt(left)
	MergeSortAlt(right)
	mergeAlt(src, left, right)
}

func mergeAlt(result, left, right []int64) {
	//	Slice: The size specifies the length. The capacity of the slice is
	//	equal to its length. A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length. For example, make([]int, 0, 10) allocates an underlying array
	//	of size 10 and returns a slice of length 0 and capacity 10 that is
	//	backed by this underlying array.

	//ret := make([]int, 0, len(left) + len(right))

	var l, r, i int // default to 0

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}

// This is the only exported function which will take a slice as an input
// and return a different slice with the original integers sorted
func Sort(input []int) []int {
	// making a new slice and copying the contents of the input
	// into it
	// this slice is going to be subsequently mutated so that
	// we get the desired order of elements
	sorted := make([]int, len(input))
	copy(sorted, input)

	// calling the private "subSort" function which can sort a
	// sub-slice, by taking two more arguments: the start and
	// end index
	subSort(sorted, 0, len(input)-1)
	return sorted
}

// this function is going to call itself recursively with the left
// and the right halves of the input slice (the divide)
// then call the "merge" function (the conquest)
func subSort(sorted []int, leftStart int, rightEnd int) {
	// stop the recursion if there is nothing to divide
	if leftStart >= rightEnd {
		return
	}
	// calculating the middle element so that we can divide our
	// slice
	middle := (leftStart + rightEnd) / 2
	// calling itself recursively with both halves
	subSort(sorted, leftStart, middle)
	subSort(sorted, middle+1, rightEnd)
	// merging the two sorted halves
	merge(sorted, leftStart, rightEnd)
}

func merge(sorted []int, leftStart int, rightEnd int) {
	// creating a temporary slice, as we can't easily do it in place
	temp := make([]int, len(sorted))
	copy(temp, sorted)

	// the end of the left sub-slice will end in the middle
	leftEnd := (leftStart + rightEnd) / 2
	// the start of the right slice will be right after the left ends
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart

	// this is the loop where the actual sorting happens
	// we iterate until either the left or the right sub-slice
	// runs out of elements
	for left <= leftEnd && right <= rightEnd {
		// here we start adding elements to the temporary
		// slice we created above
		// we choose the smaller element every time
		if sorted[left] < sorted[right] {
			temp[index] = sorted[left]
			left++
		} else {
			temp[index] = sorted[right]
			right++
		}
		index++
	}

	// here we append to the temporary slice the remaining elements
	// that were not picked in the loop above
	// first we do it for the left and the for the right sub-slice
	// one of them will not contain any remaining elements
	// so will make no changes
	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
	// finally we store the sorted numbers from the temporary slice
	// into our sorted slice
	copy(sorted, temp)
}
