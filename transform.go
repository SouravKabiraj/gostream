package gostream

func TransformFn[I any, O any](inputStream *Stream[I], transformer func(I) O) (outputStream *Stream[O]) {
	out := make(chan O)
	go func() {
		defer close(out)
		for element := range inputStream.obj {
			select {
			case <-inputStream.ctx.Done():
				return
			case out <- transformer(element):
			}
		}
	}()
	return &Stream[O]{inputStream.ctx, out}
}

func (s *Stream[T]) Transform(transformer func(T) T) (outputStream *Stream[T]) {
	out := make(chan T)
	go func() {
		defer close(out)
		for element := range s.obj {
			select {
			case <-s.ctx.Done():
				return
			case out <- transformer(element):
			}
		}
	}()
	return &Stream[T]{s.ctx, out}
}
