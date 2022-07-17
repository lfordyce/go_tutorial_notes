package interview

import (
	"testing"
)

func TestColoredBorder(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]int
		row   int
		col   int
		color int
		exp   [][]int
	}{
		{"example 1", [][]int{
			{1, 1},
			{1, 2},
		}, 0, 0, 3, [][]int{
			{3, 3},
			{3, 2},
		}},

		{"example 2", [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, 1, 1, 2, [][]int{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if out := colorBorder(tst.input, tst.row, tst.col, tst.color); !matrixEqual(out, tst.exp) {
				t.Errorf("actual (%v) != expected (%v)", out, tst.exp)
			}
		})
	}
}

func matrixEqual(gridA [][]int, gridB [][]int) bool {
	for i := 0; i < len(gridA); i++ {
		for j := 0; j < len(gridA[0]); j++ {
			if gridA[i][j] != gridB[i][j] {
				return false
			}
		}
	}
	return true
}
