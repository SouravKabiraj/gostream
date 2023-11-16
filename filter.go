package gostream

// Filter filter out items from Stream of T objects
// It creates a new Stream by applying a filtering function to each element
// in the input Stream. The context 'ctx' is used to control the execution and
// can be used to cancel the operation prematurely. The 'filterFn' function
// takes an element of type 'T' and returns a boolean indicating whether the
// element should be included in the filtered Stream.
//
// The function returns a pointer to the new filtered Stream. If the context is
// canceled, the filtering operation is terminated, and the filtered Stream may
// contain only a subset of the elements from the original Stream.
func (s *Stream[T]) Filter(filterFn func(T) bool) *Stream[T] {
	out := make(chan T)
	go func() {
		defer close(out)
		for element := range s.obj {
			if filterFn(element) {
				select {
				case <-s.ctx.Done():
					return
				case out <- element:
				}
			}
		}
	}()
	return &Stream[T]{s.ctx, out}
}
