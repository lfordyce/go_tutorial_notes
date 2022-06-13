package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := new(ListNode)
	current := head
	for _, val := range array {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head.Next
}

func ListNodeToArray(head *ListNode) []int {
	limit := 100

	times := 0
	var res []int
	for head != nil {
		times++
		if times > limit {
			break
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var list []*ListNode
	temp := head
	for {
		list = append(list, temp)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	rotateCount := k % len(list)
	if rotateCount == 0 {
		return head
	}
	pivotIndex := len(list) - rotateCount - 1
	pivot := list[pivotIndex]
	newHead := pivot.Next
	pivot.Next = nil
	temp.Next = head
	return newHead
}

func rotateRightAlt(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	l := 0
	cur := newHead
	for cur.Next != nil {
		l++
		cur = cur.Next
	}
	rotateCount := k % l
	if rotateCount == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	//for i := l - k%l; i > 0; i-- {
	//	cur = cur.Next
	//}
	for i := l - rotateCount; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}
