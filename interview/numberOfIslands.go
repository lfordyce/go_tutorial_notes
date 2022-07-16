package interview

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfsGrid(&grid, i, j)
			}
		}
	}
	return ans
}

func dfsGrid(grid *[][]byte, x int, y int) {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) || (*grid)[x][y] == '0' {
		return
	}
	(*grid)[x][y] = '0'
	dfsGrid(grid, x+1, y)
	dfsGrid(grid, x-1, y)
	dfsGrid(grid, x, y+1)
	dfsGrid(grid, x, y-1)
}

func numIslandsAlt(grid [][]byte) int {
	numOfIslands := 0
	// Initialize 2D matrix.
	visited := make([][]bool, len(grid))
	for row := range visited {
		visited[row] = make([]bool, len(grid[0]))
	}

	// For each cell - we perform dfs.
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			// Skip performing dfs from water.
			if grid[x][y] == '0' {
				continue
			}
			// We mark the board as visited as we dfs on it.
			if visited[x][y] {
				continue
			}
			// DFS will mark all connected land from this cell as visited.
			islandDFS(grid, visited, x, y)
			numOfIslands++
		}
	}

	return numOfIslands
}

func islandDFS(grid [][]byte, visited [][]bool, x, y int) {
	// Boundary check.
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}
	// Return if we've already visited it.
	if visited[x][y] {
		return
	}
	// Return if we hit water.
	if grid[x][y] == '0' {
		return
	}

	// Mark current cell as visited.
	visited[x][y] = true

	// A neighbor can be traversed to (top, bottom, right, left).
	for _, direction := range getDirections() {
		dx, dy := direction[0], direction[1]
		islandDFS(grid, visited, x+dx, y+dy)
	}
}

func getDirections() [][]int {
	return [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
}
