package interview

// TreeNode is tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func TreeNodeFromInts(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int, 0)
	return findTargetDfs(root, k, m)
}

func findTargetDfs(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return ok
	}
	m[root.Val]++
	return findTargetDfs(root.Left, k, m) || findTargetDfs(root.Right, k, m)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, idx}
		}
		m[num] = idx
	}
	return nil
}

// twoSumSorted: given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find
// two number such they add up to a specific target number.
func twoSumSorted(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
