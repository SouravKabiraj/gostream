package gostream

func (s *Stream[T]) ForEach(action func(T)) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case item, ok := <-s.obj:
			if !ok {
				return
			}
			action(item)
		}
	}
}

func (s *Stream[T]) ForEachContd(action func(T)) *Stream[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-s.ctx.Done():
				return
			case item, ok := <-s.obj:
				if !ok {
					return
				}
				action(item)
				select {
				case <-s.ctx.Done():
					return
				case out <- item:
				}
			}
		}
	}()

	return &Stream[T]{s.ctx, out}
}
