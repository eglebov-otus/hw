package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrErrorsInvalidWorkersCount = errors.New("you should use at lease 1 worker")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	if N <= 0 {
		return ErrErrorsInvalidWorkersCount
	}

	taskCh := make(chan Task, len(tasks))
	errCh := make(chan error)
	stopCh := make(chan struct{}, 1)
	errCounter := &errCounter{
		mux: &sync.Mutex{},
	}

	wgCons := &sync.WaitGroup{}
	wgErr := &sync.WaitGroup{}

	wgCons.Add(N)
	wgErr.Add(1)

	go produce(tasks, taskCh)

	go handleErrors(wgErr, errCounter, M, errCh, stopCh)

	for i := 0; i < N; i++ {
		go consume(wgCons, stopCh, taskCh, errCh)
	}

	wgCons.Wait()
	close(errCh)
	wgErr.Wait()
	close(stopCh)

	if errCounter.Get() >= M {
		return ErrErrorsLimitExceeded
	}

	return nil
}

type errCounter struct {
	mux *sync.Mutex
	cnt int
}

func (c *errCounter) Incr() {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cnt++
}

func (c *errCounter) Get() int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.cnt
}

func produce(tasks []Task, taskCh chan<- Task) {
	for _, task := range tasks {
		taskCh <- task
	}

	close(taskCh)
}

func consume(wg *sync.WaitGroup, stopCh <-chan struct{}, taskCh <-chan Task, errCh chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			return
		default:
		}

		select {
		case <-stopCh:
			return
		case task, ok := <-taskCh:
			if !ok {
				return
			}

			errCh <- task()
		}
	}
}

func handleErrors(wg *sync.WaitGroup, errCounter *errCounter, M int, errCh <-chan error, stopCh chan<- struct{}) {
	defer wg.Done()

	for err := range errCh {
		if err != nil {
			errCounter.Incr()

			if errCounter.Get() >= M {
				stopCh <- struct{}{}
			}
		}
	}
}
