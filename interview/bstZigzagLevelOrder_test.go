package interview

import (
	"fmt"
	"testing"
)

func TestBinaryTreeZigZagLevelOrderTraversal(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   [][]int
	}{
		{"base case", []int{3, 9, 20, NULL, NULL, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			root := TreeNodeFromInts(tst.input)
			actual := zigzagLevelOrder(root)
			fmt.Println(actual)
		})
	}
}

func twoDimensionalArrayEq(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	return true
}
