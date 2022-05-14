package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	if n <= 0 || m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var errCount int32
	wg := &sync.WaitGroup{}
	wg.Add(n)
	tasksChain := make(chan Task, len(tasks)+1)

	completeChan(tasksChain, tasks)

	for i := 0; i < n; i++ {
		go rutine(wg, tasksChain, &errCount, m)
	}

	wg.Wait()

	if int(errCount) >= m {
		err = ErrErrorsLimitExceeded
	}

	return err
}

func rutine(wg *sync.WaitGroup, tasks chan Task, errCount *int32, m int) {
	defer wg.Done()

	for task := range tasks {
		err := task()
		if err != nil {
			atomic.AddInt32(errCount, 1)
		}

		if int(atomic.LoadInt32(errCount)) >= m {
			return
		}
	}
}

func completeChan(ch chan Task, tasks []Task) {
	for i := 0; i < len(tasks); i++ {
		ch <- tasks[i]
	}
	close(ch)
}
