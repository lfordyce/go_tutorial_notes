package interview

// time complexity for search is O (log n)
func binarySearch(a []int, target int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > target:
		result, searchCount = binarySearch(a[:mid], target)
	case a[mid] < target:
		result, searchCount = binarySearch(a[mid:], target)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == target
		result = mid
	}
	searchCount++
	return
}

func BinarySearch(a []int, x int) int {
	r := -1 // not found
	start, end := 0, len(a)-1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}
