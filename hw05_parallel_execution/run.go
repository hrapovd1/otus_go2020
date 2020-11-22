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
	var wtg sync.WaitGroup
	wtg.Add(n)
	taskCh := make(chan Task, len(tasks))
	var errCount int32

	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	for i := 0; i < n; i++ {
		go func() {
			defer wtg.Done()
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
				if errCount >= int32(m) {
					break
				}
			}
		}()
	}

	wtg.Wait()

	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
