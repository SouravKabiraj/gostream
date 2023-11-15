package gostream

import "context"

func Stream[T any](ctx context.Context, in []T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, element := range in {
			select {
			case <-ctx.Done():
				return
			case out <- element:
			}
		}
	}()
	return out
}

func AsList[T any](items ...T) []T {
	return items
}

func Collect[T any](ctx context.Context, in <-chan T) []T {
	out := make([]T, 0)
	for element := range in {
		select {
		case <-ctx.Done():
			return out
		default:
			out = append(out, element)
		}
	}
	return out
}
