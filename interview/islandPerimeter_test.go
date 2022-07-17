package interview

import "testing"

func TestIslandPerimeter(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]int
		exp   int
	}{
		{"example 1", [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		}, 16},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := islandPerimeter(tst.input); actual != tst.exp {
				t.Errorf("actual (%d) != expected (%d)", actual, tst.exp)
			}
		})
	}
}
