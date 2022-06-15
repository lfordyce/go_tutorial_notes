package interview

import (
	"testing"
)

func TestConnectionsReorder(t *testing.T) {
	cases := [...]struct {
		desc       string
		n          int
		connection [][]int
		exp        int
	}{
		{"base case", 6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 3},
		{"example 2", 5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, 2},
		{"example 3", 3, [][]int{{1, 0}, {2, 0}}, 0},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := minReorder(tst.n, tst.connection); actual != tst.exp {
				t.Errorf("actual %v != expected %v", actual, tst.exp)
			}
		})
	}
}
