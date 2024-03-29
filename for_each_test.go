package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	ctx := context.Background()
	numbers := []Int{{1}, {2}, {3}, {5}, {6}, {7}, {8}}

	New[Int](ctx, numbers).
		ForEach(func(i Int) { fmt.Println("--->", i.number*10) })

	collectedData := New[Int](ctx, numbers).
		ForEachContd(func(i Int) { fmt.Println(i) }).
		Filter(func(i Int) bool { return i.number%2 == 0 }).
		ForEachContd(func(i Int) { fmt.Println(i) }).
		Collect()

	fmt.Println(collectedData)
}

type Int struct {
	number int
}
