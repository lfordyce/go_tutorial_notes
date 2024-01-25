package generics

// Plex multiplexes hi and lo into a single channel, prioritizing events coming into hi.
// The returned channel is closed after both hi and lo are closed and all elements have been written out.
func Plex[E any](hi, lo <-chan E) <-chan E {
	ch := make(chan E)
	go func(ch chan<- E) {
		defer close(ch)
		var (
			cache  E
			cached bool
		)
		for hi != nil || lo != nil {
			if cached {
				select {
				case e, ok := <-hi:
					if ok {
						ch <- e
					} else {
						hi = nil
					}
					continue
				case ch <- cache:
					cached = false
				}
			}
			select {
			case e, ok := <-hi:
				if !ok {
					hi = nil
					continue
				}
				ch <- e
			case e, ok := <-lo:
				if !ok {
					lo = nil
					continue
				}
				cache, cached = e, true
			}
		}
	}(ch)
	return ch
}
