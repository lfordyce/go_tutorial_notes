package interview

func maxDistToClosest(seats []int) int {
	i, j := 0, len(seats)-1

	for i < j {

	}

	for i := 0; i != j; i = (i + 1) % j {

	}

	return 0
}

/**
Input: height = [1, 8, 6 , 2, 5, 4, 8, 3, 7]
Output: 49
Explanation: The above vertical lines are represented by array
[1, 8, 6 , 2, 5, 4, 8, 3, 7]. In this case, the max area of water
*/
func maxArea(height []int) int {
	maxArea, result, i, j := 0, 0, 0, len(height)-1

	for i < j {
		if height[i] <= height[j] {
			result = height[i] * (j - i)
			i++
		} else {
			result = height[j] * (j - i)
			j--
		}

		if result > maxArea {
			maxArea = result
		}
	}
	return maxArea
}
