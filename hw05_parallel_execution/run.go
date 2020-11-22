package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	wtg := sync.WaitGroup{}
	taskPool := make([]Task, 0)
	var errCount int32
	for _, t := range tasks {
		if len(taskPool) < n+1 {
			taskPool = append(taskPool, t)
		} else {
			for _, wt := range taskPool {
				var err error
				wtg.Add(1)
				go func(t Task) {
					if err = t(); err != nil {
						atomic.AddInt32(&errCount, 1)
					}
					defer wtg.Done()
				}(wt)
				if err != nil {
					break
				}
			}
			taskPool = make([]Task, 0)
			wtg.Wait()
			if errCount >= int32(m) {
				return ErrErrorsLimitExceeded
			}
		}
	}
	return nil
}
