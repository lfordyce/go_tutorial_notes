package hacker

import "fmt"

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {


	for i := 0; i < len(arr); i++ {

		//fmt.Printf("row: %d\n", arr[i])

		for j := 0; j < len(arr[i]); j++ {



			fmt.Printf("a[%d][%d] = %d\n", i, j, arr[i][j])

		}
	}
	return 0
}
