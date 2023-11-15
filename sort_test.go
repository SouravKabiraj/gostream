package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	ctx := context.Background()
	collect := Collect[int](ctx,
		Sort[int](ctx,
			func(a int, b int) int { return a - b },
			Stream[int](ctx,
				[]int{2, 5, 1, 4, 6, 9, 3, -1},
			),
		),
	)
	fmt.Println(collect)
}
