package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestStream(t *testing.T) {
	ctx := context.Background()
	fmt.Println(Collect[int](ctx, Stream[int](ctx, AsList[int](1, 2, 3))))
}
