package interview

import "testing"

func TestFloodFill(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]int
		row   int
		col   int
		color int
		exp   [][]int
	}{
		{"example 1", [][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		}, 1, 1, 2, [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if out := floodFill(tst.input, tst.row, tst.col, tst.color); !matrixEqual(out, tst.exp) {
				t.Errorf("actual (%v) != expected (%v)", out, tst.exp)
			}
		})
	}
}
