package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Тип для счетчика ошибок.
type ErrCnt struct {
	mu  sync.Mutex
	cnt *int32
}

func (ec *ErrCnt) Incrnt() {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	atomic.AddInt32(ec.cnt, 1)
}

func (ec *ErrCnt) Get() int32 {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	return *ec.cnt
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	var wtg sync.WaitGroup // Wait group для ожидания окончания работы всех goroutines
	wtg.Add(n)
	taskCh := make(chan Task, len(tasks))

	var count int32
	errCount := ErrCnt{mu: sync.Mutex{}, cnt: &count}
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
					errCount.Incrnt()
				}
				if errCount.Get() >= int32(m) {
					if checkError {
						break
					}
				}
			}
		}()
	}

	wtg.Wait()

	if checkError && errCount.Get() >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
