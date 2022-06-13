package interview

import "testing"

func TestBinaryTreeMaxPathSum(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   int
	}{
		{"simple", []int{1, 2, 3}, 6},
		{"base case", []int{-10, 9, 20, NULL, NULL, 15, 7}, 42},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			root := TreeNodeFromInts(tst.input)
			if actual := maxPathSum(root); actual != tst.exp {
				t.Errorf("expected: %v != actual: %v", tst.exp, actual)
			}
		})
	}
}
