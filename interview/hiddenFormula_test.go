package interview

import (
	"fmt"
	"testing"
)

func TestHiddenFormula(t *testing.T) {
	cases := [...]struct {
		desc   string
		input  func(int, int) int
		target int
	}{
		{"base case", func(x int, y int) int {
			return x + y
		}, 5},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			solution := findSolution(tst.input, tst.target)
			fmt.Println(solution)
		})
	}
}
