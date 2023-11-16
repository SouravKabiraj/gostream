package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	ctx := context.Background()
	numbers := []Int{{10}, {2}, {30}, {5}, {6}, {70}, {8}}

	sortedList := New[Int](ctx, numbers).
		Sort(func(a Int, b Int) int { return a.number - b.number }).
		Collect()

	fmt.Println(sortedList)
}
