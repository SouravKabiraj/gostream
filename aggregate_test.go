package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestExampleAggregate(t *testing.T) {
	ctx := context.Background()
	filteredList := New[int](ctx, []int{1, 2, 9, 29, 67, 78}).
		Filter(func(i int) bool { return i < 10 })

	fmt.Println(AggregateFn[Int, int](filteredList, func(a *Int, b int) { a.number = a.number + b }))

	slice := []Int{{1}, {2}, {9}, {29}, {67}, {78}}
	sum := New[Int](ctx, slice).
		Aggregate(func(a *Int, b Int) { a.number = a.number + b.number })
	fmt.Println(sum) /// 186
}
