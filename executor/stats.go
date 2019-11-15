package executor

import "time"

type ExecutorStats struct {
	NumTask       uint64
	NumSuccess    uint64
	NumErr        uint64
	Started       time.Time
	Finished      time.Time
	LastProcessed time.Time
}
