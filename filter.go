package gostream

import "context"

func Filter[T any](ctx context.Context, predicate func(T) bool, in <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for element := range in {
			if predicate(element) {
				select {
				case <-ctx.Done():
					return
				case out <- element:
				}
			}
		}
	}()
	return out
}
