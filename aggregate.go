package gostream

// AggregateFn method aggregate all the items of type 'I' and return a object of type 'O'.
// it aggregates elements from the input Stream using the provided
// action function and returns the result. The context 'ctx' is used to
// control the execution and can be used to cancel the operation prematurely.
// The 'action' function takes a pointer to the aggregated result 'O' (a different type of Object) and an
// element of type 'I' (Type of input object), and modifies the result accordingly.
// Disclamer: 'O' type can't be a premitive datatype
// Example: to add all the integer
// var sumInt int = AggregateFn[Int, int](ctx, func(a *Int, b int) { a.number = a.number + b }, New[int](ctx, []int{1,2,3})).number
func AggregateFn[O any, I any](action func(*O, I), inStream *Stream[I]) *O {
	var aggregatedResult O
	for {
		select {
		case <-inStream.ctx.Done():
			return &aggregatedResult
		case item, ok := <-inStream.obj:
			if !ok {
				return &aggregatedResult
			}
			action(&aggregatedResult, item)
		}
	}
}

// Aggregate method aggregate all the items of type T and return a object of type T.
// it applies the provided action function to each element in the Stream
// and aggregates the results. The context 'ctx' is used to control the
// execution and can be used to cancel the operation prematurely. The 'action'
// function takes a pointer to the aggregated result 'T' and an element of type
// 'T', and modifies the result accordingly.
//
// The function returns a pointer to the final aggregated result. If the
// context is canceled or if the input Stream is closed, the function
// prematurely returns the current aggregated result.
/*
Example: Add a list of imteger
slice := []Int{{1}, {2}, {9}, {29}, {67}, {78}}
sum := New[Int](ctx, slice).Aggregate(ctx, func(a *Int, b Int) { a.number = a.number + b.number })
fmt.Println(sum.number) /// 186
*/
func (s *Stream[T]) Aggregate(action func(*T, T)) *T {
	var aggregatedResult T
	for {
		select {
		case <-s.ctx.Done():
			return &aggregatedResult
		case item, ok := <-s.obj:
			if !ok {
				return &aggregatedResult
			}
			action(&aggregatedResult, item)
		}
	}
}
