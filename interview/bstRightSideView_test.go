package interview

import (
	"testing"
)

func TestBstRightSideView(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   []int
	}{
		{"empty case", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"base case", []int{1, 2, 3, NULL, 5, NULL, 4}, []int{1, 3, 4}},
		{"second case", []int{3, 9, 20, NULL, NULL, 15, 7}, []int{3, 20, 7}},
		{"stumped", []int{1, 2, 3, 4}, []int{1, 3, 4}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			root := TreeNodeFromInts(tst.input)
			if actual := rightSideViewAlt(root); !IntArrayEqual(actual, tst.exp) {
				t.Errorf("expected: %v != actual: %v", tst.exp, actual)
			}
		})
	}
}
