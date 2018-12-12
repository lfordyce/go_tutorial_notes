package interview

import (
	"fmt"
	"math"
)

// Initial assignments:
// array = [1, 2, 3, 4]
// length_of_array = array.length = 4
// no_of_left_rotation = k = 2
// new_arr = Arra.new(length_of_array)
// new_arr: [nil, nil, nil, nil]
//
// NOTE:
// length_of_array.times do |i|
// is equivalent to
// for(i = 0; i < length_of_array; i++)
//
// Algorithm to calculate new index and update new array for each index (i):
// new_index = (i + no_of_left_rotation) % length_of_array
// new_arr[i] = array[new_index]
//
// LOOP1:
// i = 0
// new_index = (0 + 2) % 4 = 2
// new_arr[i = 0] = array[new_index = 2] = 3
// new_arr: [3, nil, nil, nil]
//
// LOOP2:
// i = 1
// new_index = (1 + 2) % 4 = 3
// new_arr[i = 1] = array[new_index = 3] = 4
// new_arr: [3, 4, nil, nil]
//
// LOOP3:
// i = 2
// new_index = (2 + 2) % 4 = 0
// new_arr[i = 2] = array[new_index = 0] = 1
// new_arr: [3, 4, 1, nil]
//
// LOOP4:
// i = 3
// new_index = (3 + 2) % 4 = 1
// new_arr[i = 3] = array[new_index = 1] = 2
// new_arr: [3, 4, 1, 2]
//
// After final loop our new roated array is [3, 4, 1, 2]
// You can return the output:
// new_arr.join(' ') => 3 4 1 2
func RotLeft(a []int32, d int32) []int32 {

	length := len(a)

	tmp := make([]int32, len(a))

	for i := 0; i < length; i++ {
		index := int32(i)
		rotation := (index + d ) % int32(length)
		tmp[i] = a[rotation]
	}
	return tmp
}

// Complete the minimumBribes function below.
func MinimumBribes(q []int32) {

	bribes := 0

	for i := len(q) -1; i >= 0; i-- {
		item := int(q[i])
		if item - (i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j:= int(math.Max(0, float64(item - 2))); j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	fmt.Println(bribes)
}

func MinimumSwaps(arr []int32) int32 {

	swaps := 0
	for i := 0; i < len(arr); i ++ {

		if i + 1 != int(arr[i]) {
			t := i
			for int(arr[t]) != i + 1 {
				t++
			}
			temp := arr[t]
			arr[t] = arr[i]
			arr[i] = temp
			swaps++
		}
	}
	return int32(swaps)
}