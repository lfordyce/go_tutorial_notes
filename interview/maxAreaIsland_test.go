package interview

import "testing"

func TestMaxAreaIsland(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]int
		exp   int
	}{
		{desc: "base case", input: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}, exp: 6},
		{desc: "example 2", input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
		}, exp: 0},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := maxAreaOfIslandAlt(tst.input); tst.exp != actual {
				t.Errorf("expected (%d) != actual (%d)", tst.exp, actual)
			}
		})
	}
}
