package interview

import "testing"

func TestMinHeightTrees(t *testing.T) {
	cases := [...]struct {
		desc  string
		edges [][]int
		n     int
		exp   []int
	}{
		{"example 1", [][]int{{1, 0}, {1, 2}, {1, 3}}, 4, []int{1}},
		{"example 2", [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, 6, []int{3, 4}},
		{"example 3", [][]int{}, 1, []int{0}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := findMinHeightTreesAlt(tst.n, tst.edges); !IntArrayEqual(actual, tst.exp) {
				t.Errorf("expected %v != actual %v", tst.exp, actual)
			}
		})
	}
}
