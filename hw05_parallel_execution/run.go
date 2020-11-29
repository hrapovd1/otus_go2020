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
	var wtg sync.WaitGroup // Wait group для ожидания окончания работы всех goroutines
	wtg.Add(n)
	taskCh := make(chan Task, len(tasks))
	var errCount int32
	var checkError bool
	if m >= 0 { // Если m положительное, то проверяем ошибки
		checkError = true
	} else { // Если m отрицательное, то игнорируем ошибки
		checkError = false
	}

	// Цикл отправки задач в goroutines
	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	for i := 0; i < n; i++ {
		go func() {
			defer wtg.Done()
			// Цикл чтения задач в goroutine
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
				if errCount >= int32(m) {
					if checkError {
						break
					}
				}
			}
		}()
	}

	wtg.Wait()

	if checkError && errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
