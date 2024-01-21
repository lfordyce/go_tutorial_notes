package hacker

import "sort"

func triangleOrNot(a []int32, b []int32, c []int32) (ret []string) {
	if isPossibleTriangle(a) {
		ret = append(ret, "Yes")
	} else {
		ret = append(ret, "No")
	}

	if isPossibleTriangle(b) {
		ret = append(ret, "Yes")
	} else {
		ret = append(ret, "No")
	}

	if isPossibleTriangle(c) {
		ret = append(ret, "Yes")
	} else {
		ret = append(ret, "No")
	}

	return
}

func checkValid(a, b, c int32) bool {
	if a+b <= c || a+c <= b || b+c <= a {
		return true
	}
	return false
}

func checkTriangle(input []int32) bool {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	if input[0]+input[1] <= input[2] {
		return true
	}
	return false
}

func isPossibleTriangle(arr []int32) bool {
	// If number of elements are less than 3, then no
	// triangle is possible
	if len(arr) < 3 {
		return false
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// loop for all 3 consecutive triplets
	for i := 0; i < len(arr)-2; i++ {
		// If triplet satisfies triangle condition, break
		if arr[i]+arr[i+1] <= arr[i+2] {
			return true
		}
	}
	return false
}
