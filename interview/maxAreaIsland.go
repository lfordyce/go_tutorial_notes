package interview

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}
			area := areaOfIsland(grid, i, j)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func areaOfIsland(grid [][]int, x int, y int) int {
	if !isInGrid(grid, x, y) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	total := 1
	for i := 0; i < 4; i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]
		total += areaOfIsland(grid, dx, dy)
	}
	return total
}

func isInGrid(grid [][]int, x int, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func maxAreaOfIslandAlt(grid [][]int) int {
	var max int
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				continue
			}
			i := searchIsland(grid, x, y)
			if max < i {
				max = i
			}
		}
	}
	return max
}

func searchIsland(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	return 1 + searchIsland(grid, x-1, y) + searchIsland(grid, x+1, y) + searchIsland(grid, x, y-1) + searchIsland(grid, x, y+1)
}
