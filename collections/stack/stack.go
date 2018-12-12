package stack

import "sync"

//type (
//	Stack struct {
//		top *node
//		length int
//	}
//	node struct {
//		value interface{}
//		prev *node
//	}
//)
//
//func New() *Stack {
//	return &Stack{nil, 0}
//}
//
//func (stk *Stack) Len() int {
//	return stk.length
//}
//
//func (stk *Stack) Peek() interface{} {
//	if stk.length == 0 {
//		return nil
//	}
//	return stk.top.value
//}
//
//func (stk *Stack) Pop() interface{} {
//	if stk.length == 0 {
//		return nil
//	}
//	n := stk.top
//	stk.top = n.prev
//	stk.length--
//	return n.value
//}
//
//func (stk *Stack) Push(value interface{})  {
//
//	n := &node{value:value, prev: stk.top}
//	stk.top = n
//	stk.length++
//}

type (
	Stack struct {
		lock *sync.RWMutex
		head *node
		Size int
	}
	node struct {
		value interface{}
		next *node
	}
)

func New() *Stack  {
	stk := new(Stack)
	stk.lock = &sync.RWMutex{}

	return stk
}

func (stk *Stack) Push(value interface{}) {
	stk.lock.Lock()

	node := new(node)
	node.value = value
	temp := stk.head
	node.next = temp
	stk.head = node
	stk.Size++

	stk.lock.Unlock()
}

func (stk *Stack) Pop() interface{}  {
	if stk.head == nil {
		return nil
	}
	stk.lock.Lock()

	n := stk.head.value
	stk.head = stk.head.next
	stk.Size--

	stk.lock.Unlock()

	return n
}

func (stk *Stack) Peek() interface{} {
	if stk.Size == 0 {
		return nil
	}
	stk.lock.Lock()
	h := stk.head.value
	stk.lock.Unlock()

	return h
}
