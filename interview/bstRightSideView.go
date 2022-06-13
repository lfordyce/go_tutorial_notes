package interview

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rightSide := make([]int, 0)
	dfsRightSide(root, 1, &rightSide)
	return rightSide
}

func dfsRightSide(root *TreeNode, level int, side *[]int) {
	if root == nil {
		return
	}
	if len(*side) < level {
		*side = append(*side, root.Val)
	}
	dfsRightSide(root.Right, level+1, side)
	dfsRightSide(root.Left, level+1, side)
}

func rightSideViewAlt(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}
	return res
}
