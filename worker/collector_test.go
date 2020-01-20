package worker

import "testing"

func TestCollector(t *testing.T) {
	collector := StartDispatcher(5)

	for i, job := range CreateJobs(100) {
		collector.Work <- Work{
			Id:  i,
			Job: job,
		}
	}
}
