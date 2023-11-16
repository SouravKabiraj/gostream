package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestStream(t *testing.T) {
	ctx := context.Background()
	out := New[Int](ctx, []Int{{1}, {2}, {3}, {4}}).Collect()
	fmt.Println(out)

	val1 := New[Int](ctx, []Int{{50}, {60}, {70}, {80}}).Aggregate(func(sum *Int, a Int) { sum.number = sum.number + a.number }).number
	val2 := New[Int](ctx, []Int{{1}, {2}, {3}, {4}}).Aggregate(func(sum *Int, a Int) { sum.number = sum.number + a.number }).number
	val3 := New[Int](ctx, []Int{{10}, {20}, {30}, {40}}).Aggregate(func(sum *Int, a Int) { sum.number = sum.number + a.number }).number
	fmt.Println(val1, val2, val3)
}
