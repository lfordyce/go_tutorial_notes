package interview

type point struct {
	x int
	y int
}

type gridInfo struct {
	m             int
	n             int
	grid          [][]int
	originalColor int
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	srcColor := grid[row][col]
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// Return value: 0 means this point does not belong to this connected component
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return 0
		}
		// Because the color attribute of the point is modified later, the visited judgment must be made first.
		// If it has been traversed, it must belong to this connected component.
		if vis[i][j] {
			return 1
		}
		// It has not been traversed, and the color is different from the click point, so it does not belong to this connected component
		if grid[i][j] != srcColor {
			return 0
		}

		vis[i][j] = true
		res := dfs(i-1, j) + dfs(i+1, j) + dfs(i, j+1) + dfs(i, j-1)
		if res < 4 {
			grid[i][j] = color
		}
		return 1
	}
	dfs(row, col)
	return grid
}

//func dfsGridColor() {}
