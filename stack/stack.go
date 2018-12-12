package stack

//SafeStack is the structure that contains the channel used to
//communicate with the stack.
type SafeStack struct {
	op chan func(*stack)
}

//Len will get return the number of items in the stack.
func (s *SafeStack) Len() (i int64) {
	lChan := make(chan int64)
	s.op <- func(curr *stack) {
		lChan <- curr.size
	}
	return <-lChan
}

//Pop will perform a pop on the stack, removing the first item
//and returning it's value.
func (s *SafeStack) Pop() (v interface{}) {
	vChan := make(chan interface{})
	s.op <- func(curr *stack) {
		if curr.size == 0 {
			vChan <- nil
			return
		}
		val := curr.top.value
		curr.top = curr.top.next
		curr.size--
		vChan <- val
		return
	}
	return <-vChan
}

//Push will push the value v onto the stack.
func (s *SafeStack) Push(v interface{}) {
	s.op <- func(curr *stack) {
		curr.top = &elem{v, curr.top}
		curr.size++
	}
}

//Destroy closes the primary channel thus stopping
//the running go-routine.
func (s *SafeStack) Destroy() {
	close(s.op)
}

//New creates a new Safe Stack, this also starts the go-routine
//so once this is called, you need to clean up after yourself
//by using the Destroy method.
func New() (s *SafeStack) {
	s = &SafeStack{make(chan func(*stack))}
	go s.loop()
	return
}

//stack is the basic container for the stack
type stack struct {
	top  *elem
	size int64
}

//elem is the element structure
type elem struct {
	value interface{}
	next  *elem
}

//loop creates the guarded data structure and listens for
//methods on the op channel. loop terminates when the op
//channel is closed.
func (s *SafeStack) loop() {
	st := &stack{}
	for op := range s.op {
		op(st)
	}
}

// ------------------

// type Element struct {
// 	item interface{}
// 	next *Element
// }
// type Stack struct {
// 	// lock  *sync.Mutex
// 	lock sync.Mutex
// 	data []int
// 	// head  *Element
// 	// depth uint64
// }
// func NewStack() *Stack {
// 	return &Stack{lock: sync.Mutex{}, data: make([]int, 0)}
// }
// func (stk *Stack) Push(v int) {
// 	stk.lock.Lock()
// 	defer stk.lock.Lock()

// 	stk.data = append(stk.data, v)
// }
// func (stk *Stack) Pop() (int, error) {
// 	stk.lock.Lock()
// 	defer stk.lock.Unlock()

// 	l := len(stk.data)
// 	if l == 0 {
// 		return 0, errors.New("Empty Stack")
// 	}

// 	res := stk.data[l-1]
// 	stk.data = stk.data[:l-1]
// 	return res, nil
// }

// func New() *Stack {
// 	stack := new(Stack)
// 	// var stack *Stack = new(Stack)

// 	stack.lock = &sync.Mutex{}
// 	stack.depth = 0
// 	return stack
// }

// func (stk *Stack) Push(data interface{}) {

// 	stk.lock.Lock()

// 	element := new(Element)
// 	element.head = data

// 	stack.head = &Element{item: data, next: stack.head}
// 	stack.depth++
// }
