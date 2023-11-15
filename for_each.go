package gostream

import "context"

func ForEach[T any](ctx context.Context, action func(T), in <-chan T) {
	for {
		select {
		case <-ctx.Done():
			return
		case item, ok := <-in:
			if !ok {
				return
			}
			action(item)
		}
	}
}

func ForEachContd[T any](ctx context.Context, action func(T), in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case item, ok := <-in:
				if !ok {
					return
				}
				action(item)
				select {
				case <-ctx.Done():
					return
				case out <- item:
				}
			}
		}
	}()

	return out
}
