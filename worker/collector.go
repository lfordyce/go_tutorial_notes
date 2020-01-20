package worker

import "log"

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work chan Work
	End  chan struct{}
}

func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan Work)
	end := make(chan struct{})
	collector := Collector{
		Work: input,
		End:  end,
	}

	for i < workerCount {
		i++
		log.Println("starting worker: ", i)
		worker := Worker{
			ID:            i,
			WorkerChannel: WorkerChannel,
			Channel:       make(chan Work),
			End:           make(chan struct{}),
		}
		worker.Start()
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				worker := <-WorkerChannel
				worker <- work
			}
		}
	}()
	return collector
}
