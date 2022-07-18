package interview

import "strconv"

func dfsBoard(board [][]byte, r int, c int) {

	if r >= len(board) || c >= len(board[0]) || r < 0 || c < 0 {
		return
	}

	if board[r][c] != 'E' {
		return
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	totalMines := 0
	for _, dir := range dirs {
		nr := r + dir[0]
		nc := c + dir[1]
		if nr < len(board) && nc < len(board[0]) && nr >= 0 && nc >= 0 {
			if board[nr][nc] == 'M' {
				totalMines++
			}
		}
	}

	if totalMines != 0 {
		board[r][c] = byte(totalMines) + '0'
		return
	}

	board[r][c] = 'B'
	for _, dir := range dirs {
		nr := r + dir[0]
		nc := c + dir[1]
		if nr < len(board) && nc < len(board[0]) && nr >= 0 && nc >= 0 {
			dfsBoard(board, nr, nc)
		}
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	row := click[0]
	col := click[1]
	if board[row][col] == 'M' {
		board[row][col] = 'X'
		return board
	}
	dfsBoard(board, row, col)
	return board
}

func updateBoard1(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	}
	rows, cols := len(board)-1, len(board[0])-1
	isValid := func(r, c int) bool {
		if r < 0 || r > rows || c < 0 || c > cols {
			return false
		}
		return true
	}

	// mark with the appropriate new value. return if we are making a new wall
	mark := func(r, c int) bool {
		// count the surrounding cells
		count := 0
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if isValid(r+i, c+j) && board[r+i][c+j] == 'M' {
					count++
				}
			}
		}
		if count > 0 {
			// alternate: board[r][c] = byte(count + '0')
			board[r][c] = strconv.Itoa(count)[0]
			return true
		} else {
			board[r][c] = 'B'
			return false
		}
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// check if we are in the board
		if !isValid(r, c) {
			return
		}

		// do not continue if we hit a wall
		if mark(r, c) {
			return
		}

		// traverse all the directions.
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if j == 0 && i == 0 {
					continue
				}
				dfs(r+i, c+j)
			}
		}
	}
	dfs(r, c)
	return board
}
