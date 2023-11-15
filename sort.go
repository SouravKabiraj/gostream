package gostream

import (
	"context"
	"sync"
)

func Sort[T any](ctx context.Context, sortAlgo func(T, T) int, in <-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	data := make([]T, 0)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					SortSlice(data, sortAlgo)
					for _, item := range data {
						select {
						case <-ctx.Done():
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
		SortSlice(data, sortAlgo)
	}()

	go func() {
		wg.Wait()
	}()

	return out
}

func SortSlice[T any](data []T, sortAlgo func(T, T) int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortAlgo(data[j], data[j+1]) > 0 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
