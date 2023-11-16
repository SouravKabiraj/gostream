package gostream

import "context"

// Stream represents a streaming channel of elements of type T.
type Stream[T any] struct {
	ctx context.Context
	obj <-chan T
}

// New creates a new Stream from a slice of elements 'in'.
// It returns a pointer to the created Stream.
// To create a stream of int follow the example...
// Example: gostream.New[int](context.Background(),[]int{1,2,3})
func New[T any](ctx context.Context, in []T) *Stream[T] {
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
	return &Stream[T]{ctx, out}
}

// Collect retrieves all elements from the Stream until the context is canceled.
// It returns a slice containing the collected elements.
// Example: gostream.New[int](context.Background(),[]int{1,2,3}).Collect(context.Background())
func (s *Stream[T]) Collect() []T {
	out := make([]T, 0)
	for element := range s.obj {
		select {
		case <-s.ctx.Done():
			return out
		default:
			out = append(out, element)
		}
	}
	return out
}
