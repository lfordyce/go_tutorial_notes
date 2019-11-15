package concurrency

// yieldFn reports true if an iteration should continue.
// It is called on values of a collection
type yieldFn func(interface{}) (stopIterating bool)

// mapperFn calls yieldFn for each member of a collection
type mapperFn func(yieldFn)

// iteratorFn returns the next item in an iteration or the zero value.
// The second return value is true when iteration is complete.
type iteratorFn func() (value interface{}, done bool)

// cancelFn should be called to clean up the goroutine that would otherwise leak.
type cancelFn func()

// mapperToIterator returns an iteratorFn version of a mappingFn. The second
// return value must be called at the end of iteration, or the underlying
// goroutine will leak.
func mapperToIterator(m mapperFn) (iteratorFn, cancelFn) {
	generatedValues := make(chan interface{}, 1)
	stopCh := make(chan interface{}, 1)

	go func() {
		m(func(obj interface{}) (stopIterating bool) {
			select {
			case <-stopCh:
				return false
			case generatedValues <- obj:
				return true
			}
		})
		close(generatedValues)
	}()
	iter := func() (value interface{}, notDone bool) {
		value, notDone = <-generatedValues
		return
	}
	return iter, func() {
		stopCh <- nil
	}
}
