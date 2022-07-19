package interview

import (
	"fmt"
	"strings"
	"testing"
)

var jumbleRaw string = `TYNAUUHTWGCODE
UCGDPHGEFYZXUD
YOSEAPESIWEBEE
PMQVLPLTLEKFGP
TPPEKAPIEZEOVL
MIBLVUCNCZMHSO
ALJOAKPGHAMNZY
AEWPROGIKPTGSS
RRSEYPYHLQBIRF
SQWRYOXJVZPGOJ
EJVMODULENELEN
OENVIRONMENTXL
GWCQPZMWAJKYRR
RUGGXZMOQYCMWL`

var wordlist []string = []string{
	"WEB",
	"TESTING",
	"DEVELOPER",
	"APPLICATION",
	"ENVIRONMENT",
	"COMPILER",
	"DEPLOY",
	"FILE",
	"MODULE",
	"CODE",
	"SEO",
	"VAR",
}

func TestWordJumble(t *testing.T) {
	fmt.Printf("starting word search\n\n")

	fmt.Println("word jumble:")
	jumble := [14][14]string{}
	for row, line := range strings.Split(jumbleRaw, "\n") {
		for col, char := range strings.Split(line, "") {
			jumble[row][col] = char
		}
	}

	fmt.Println(jumble)

	fmt.Printf("\nwords to find:")
	fmt.Println(wordlist)

	fmt.Printf("starting search\n\n")

	// your code here...
	horizontal := "HORIZONTAL"
	vertical := "VERTICAL"
	diagonal := "DIAGONAL"

	// anonymous recursive function for dfs lookup
	var dfsLookup func(jumble [14][14]string, visited [][]bool, word string, index, x, y int, direction *string) bool

	dfsLookup = func(jumble [14][14]string, visited [][]bool, word string, index, x, y int, direction *string) bool {
		if x < 0 || x >= len(jumble) || y < 0 || y >= len(jumble[0]) || visited[x][y] {
			return false
		}
		if index == len(word)-1 {
			if jumble[x][y] == string(word[index]) {
				fmt.Printf("direction (%s) ", *direction)
				return true
			}
		}
		if jumble[x][y] == string(word[index]) {
			visited[x][y] = true

			if dfsLookup(jumble, visited, word, index+1, x+1, y, &vertical) ||
				dfsLookup(jumble, visited, word, index+1, x, y+1, &horizontal) ||
				dfsLookup(jumble, visited, word, index+1, x+1, y+1, &diagonal) {
				return true
			}
			visited[x][y] = false
		}
		return false
	}

	for _, word := range wordlist {
		visited := make([][]bool, len(jumble))
		for i := 0; i < len(visited); i++ {
			visited[i] = make([]bool, len(jumble[0]))
		}
		for i, v := range jumble {
			for j := range v {
				var direction string
				if dfsLookup(jumble, visited, word, 0, i, j, &direction) {
					// coordinates are represented as 0 index
					fmt.Printf("found: (%s) starting at: (%d, %d) \n", word, i, j)
				}
			}
		}
	}
	fmt.Println("complete!")
}

func isInBoard(jumble [14][14]string, x, y int) bool {
	return x >= 0 && x < len(jumble) && y >= 0 && y < len(jumble[0])
}

//
//func searchWordAlt(jumble [14][14]string, visited [][]bool, word string, index, x, y int, direction *string) bool {
//	if x < 0 || x >= len(jumble) || y < 0 || y >= len(jumble[0]) || visited[x][y] {
//		return false
//	}
//	dirs := [][]int{
//		{1, 1},
//		{0, 1},
//		{1, 0},
//	}
//
//	if index == len(word)-1 {
//		if jumble[x][y] == string(word[index]) {
//			fmt.Printf("DIRECTION (%s) ", *direction)
//			return true
//		}
//	}
//	if jumble[x][y] == string(word[index]) {
//		visited[x][y] = true
//
//		//
//		//if isInBoard(jumble, x, y) && !visited[x][y] && searchWordAlt(jumble, visited, word, index+1, x, y) {
//		//	return true
//		//}
//
//		for i := 0; i < 3; i++ {
//			nx := x + dirs[i][0]
//			ny := y + dirs[i][1]
//			if searchWordAlt(jumble, visited, word, index+1, nx, ny) {
//				return true
//			}
//		}
//		visited[x][y] = false
//	}
//	return false
//}

func searchWord(jumble [14][14]string, visited [][]bool, word string, index, x, y int) bool {
	dirs := [][]int{
		{1, 1},
		{0, 1},
		{1, 0},
	}
	if index == len(word)-1 {
		return jumble[x][y] == string(word[index])
	}
	if jumble[x][y] == string(word[index]) {
		visited[x][y] = true
		for i := 0; i < 3; i++ {
			nx := x + dirs[i][0]
			ny := y + dirs[i][1]
			if isInBoard(jumble, nx, ny) && !visited[nx][ny] && searchWord(jumble, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

//func patternSearch(jumble [][]string, word string) {
//	visited := make([][]bool, len(jumble))
//	for i := 0; i < len(visited); i++ {
//		visited[i] = make([]bool, len(jumble[0]))
//	}
//	for i, v := range jumble {
//		for j := range v {
//			if searchWord(jumble, visited, word, 0, i, j) {
//				fmt.Printf("FOUND: (%s) STARTING AT (%d, %d)\n", word, i, j)
//			}
//		}
//	}
//}

//var dirs = [][]int{
//	{1, 1},
//	{0, 1},
//	{1, 0},
//}
//var dirs2 = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
