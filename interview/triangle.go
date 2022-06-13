package interview

/**
minimumTotal: Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally,
if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
Bottom up solution with dynamic programming:
For 'top-down' DP, starting from the node on the very top, we recursively find the minimum path sum of each node.
When a path sum is calculated, we store it in an array (memoization);
the next time we need to calculate the path sum of the same node,
just retrieve it from the array.
However, you will need a cache that is at least the same size as the input triangle itself to store the pathsum,
which takes O(N^2) space.
With some clever thinking, it might be possible to release some of the memory that will never be used after a particular point,
but the order of the nodes being processed is not straightforwardly seen in a recursive solution,
so deciding which part of the cache to discard can be a hard job.

'Bottom-up' DP, on the other hand, is very straightforward: we start from the nodes on the bottom row;
the min pathsums for these nodes are the values of the nodes themselves.
From there, the min pathsum at the ith node on the kth row would be the
lesser of the pathsums of its two children plus the value of itself, i.e.:
*/
func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
