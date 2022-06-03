package hacker

import "fmt"

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	sumArray := make([]int32, 0)

	for i := 0; i < len(arr)-2; i++ {

		for j := 0; j < len(arr)-2; j++ {

			fmt.Printf(" Top Row [%d], [%d], [%d]\n", arr[i][j], arr[i][j+1], arr[i][j+2])
			fmt.Printf(" Middle, [%d]\n", arr[i+1][j+1])
			fmt.Printf(" Bottom Row, [%d], [%d], [%d]\n", arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			sumArray = append(sumArray, sum)

		}
	}

	var max int32
	for i, e := range sumArray {
		if i == 0 || e > max {
			max = e
		}
	}

	fmt.Printf("MAX %d\n", max)
	return max
}
