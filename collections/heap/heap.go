package heap

import (
	"sort"
	"sync"
)

type Heap struct {
	b []int
	c *sync.Cond
}

func NewHeap() *Heap {
	return &Heap{
		c: sync.NewCond(new(sync.Mutex)),
	}
}

func (h *Heap) Pop() int {
	h.c.L.Lock()
	defer h.c.L.Unlock()
	for len(h.b) == 0 {
		h.c.Wait()
	}
	// There is definitely something in there
	x := h.b[len(h.b)-1]
	h.b = h.b[:len(h.b)-1]
	return x
}

func (h *Heap) Push(x int) {
	defer h.c.Signal()
	h.c.L.Lock()
	defer h.c.L.Unlock()

	// Add and sort to maintain priority (not really how the heap works)
	h.b = append(h.b, x)
	sort.Ints(h.b)
}
