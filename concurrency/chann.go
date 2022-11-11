package concurrency

// MakeChan is a generic implementation that returns a sender and
// a receiver of an arbitrary sized channel of a arbitrary type.
//
// If the given size is positive, the returned channel is a regular
// fix-sized buffered channel.
// If the given size is zero, the returned channel is a unbuffered
// channel.
// If the given size is -1, the returned channel is a infinite sized
// buffered channel.
func MakeChan[T any](size int) (chan<- T, <-chan T) {
	switch {
	case size == 0:
		ch := make(chan T)
		return ch, ch
	case size > 0:
		ch := make(chan T, size)
		return ch, ch
	case size != -1:
		panic("unsupported buffer size")
	default:
		// size == 1
	}

	in, out := make(chan T), make(chan T)

	go func() {
		var q []T
		for {
			e, ok := <-in
			if !ok {
				close(out)
				return
			}
			q = append(q, e)
			for len(q) > 0 {
				select {
				case out <- q[0]:
					q = q[1:]
				case e, ok := <-in:
					if ok {
						q = append(q, e)
						break
					}
					for _, e := range q {
						out <- e
					}
					close(out)
					return
				}
			}
		}
	}()
	return in, out
}
