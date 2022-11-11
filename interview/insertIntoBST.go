package interview

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}

func insert(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}
	if n.Val < val {
		n.Right = insert(n.Right, val)
	} else {
		n.Left = insert(n.Left, val)
	}
	return n
}
