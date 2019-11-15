package concurrency

import "sync"

func Parallelize(functions ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(functions))

	defer wg.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer wg.Done()
			copy()
		}(function)
	}
}
