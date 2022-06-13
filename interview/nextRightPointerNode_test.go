package interview

import (
	"fmt"
	"testing"
)

func TestNextRightPointerNode(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   int
	}{
		{"simple", []int{1, 2, 3, 4, 5, 6, 7}, 6},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			tree := &Node{}
			for i := 0; i < len(tst.input); i++ {
				if err := tree.Insert(tst.input[i]); err != nil {
					t.Fatal("Error inserting value")
				}
			}
			tree.Traverse(func(node *Node) {
				fmt.Printf("%v", node.Val)
			})
			//root := TreeNodeFromInts(tst.input)
			//if actual := maxPathSum(root); actual != tst.exp {
			//	t.Errorf("expected: %v != actual: %v", tst.exp, actual)
			//}
		})
	}
}
