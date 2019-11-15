package concurrency

type Comp struct {
	value interface{}
	ok    bool
}

type Future chan Comp

func future(f func() (interface{}, bool)) Future {
	future := make(chan Comp)

	go func() {
		v, ok := f()
		c := Comp{
			value: v,
			ok:    ok,
		}
		for {
			future <- c
		}
	}()
	return future
}

type Promise struct {
	lock chan int
	ft   Future
	full bool
}

func promise() Promise {
	return Promise{
		lock: make(chan int, 1),
		ft:   make(chan Comp),
		full: false,
	}
}

func (pr Promise) future() Future {
	return pr.ft
}
