package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	ctx := context.Background()
	numbers := []Int{{1}, {2}, {3}, {5}, {6}, {7}, {8}}

	collectedData := New[Int](ctx, numbers).
		Filter(func(i Int) bool { return i.number%2 == 0 }).
		Collect()

	fmt.Println(collectedData)
}
