package interview

func islandPerimeter(grid [][]int) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					counter++
				}
				if i-1 < 0 || grid[i-1][j] == 0 {
					counter++
				}
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
					counter++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					counter++
				}
			}
		}
	}
	return counter
}

// Time O(NM) where n is the number of rows and m is the number of columns in the grid
// Space O(NM)
func islandPerimeterDfs(grid [][]int) int {
	var ans int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfsPerimeter(grid, i, j, &ans)
			}
		}
	}
	return ans
}

func dfsPerimeter(grid [][]int, row, col int, perimeter *int) {
	if row < 0 || col < 0 || row > len(grid)-1 || col > len(grid[0])-1 || grid[row][col] == -1 || grid[row][col] == 0 {
		return
	}
	grid[row][col] = -1
	// travel to the left col - 1
	if (col-1 >= 0 && grid[row][col-1] == 0) || (col-1 < 0) {
		*perimeter += 1
	}

	// travel to the right col + 1
	if (col+1 <= len(grid[0])-1 && grid[row][col+1] == 0) || (col+1 > len(grid[0])-1) {
		*perimeter += 1
	}

	// travel to the bottom row + 1
	if (row+1 <= len(grid)-1 && grid[row+1][col] == 0) || (row+1 > len(grid)-1) {
		*perimeter += 1
	}

	// travel to the top row - 1
	if (row-1 >= 0 && grid[row-1][col] == 0) || (row-1 < 0) {
		*perimeter += 1
	}

	dfsPerimeter(grid, row+1, col, perimeter)
	dfsPerimeter(grid, row-1, col, perimeter)
	dfsPerimeter(grid, row, col+1, perimeter)
	dfsPerimeter(grid, row, col-1, perimeter)
}
