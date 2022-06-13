package interview

import "errors"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("cannot insert a value into a nil tree")
	}

	switch {
	case n.Val == 0:
		n.Val = value
		return nil
	case value == n.Val:
		return nil
	case value < n.Val:
		if n.Left == nil {
			n.Left = &Node{Val: value}
			return nil
		}
		return n.Left.Insert(value)
	case value > n.Val:
		if n.Right == nil {
			n.Right = &Node{Val: value}
			return nil
		}
		return n.Right.Insert(value)
	}
	return nil
}

func (n *Node) Traverse(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.Traverse(f)
	f(n)
	n.Right.Traverse(f)
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	return root
}
