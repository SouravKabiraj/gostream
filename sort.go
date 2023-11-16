package gostream

import (
	"sync"
)

func (s *Stream[T]) Sort(sortAlgo func(T, T) int) *Stream[T] {
	out := make(chan T)
	var wg sync.WaitGroup
	data := make([]T, 0)

	go func() {
		defer close(out)

		for {
			select {
			case <-s.ctx.Done():
				return
			case val, ok := <-s.obj:
				if !ok {
					sortSlice(data, sortAlgo)
					for _, item := range data {
						select {
						case <-s.ctx.Done():
							return
						case out <- item:
						}
					}
					return
				}
				data = append(data, val)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sortSlice(data, sortAlgo)
	}()

	go func() {
		wg.Wait()
	}()

	return &Stream[T]{s.ctx, out}
}

func sortSlice[T any](data []T, sortAlgo func(T, T) int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortAlgo(data[j], data[j+1]) > 0 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
