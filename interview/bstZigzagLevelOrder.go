package interview

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	search(root, 0, &res)
	return res
}

func search(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}
	search(root.Left, depth+1, res)
	search(root.Right, depth+1, res)
}
