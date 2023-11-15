package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	ctx := context.Background()
	ct := Collect[Int](ctx,
		ForEachContd[Int](ctx,
			func(i Int) { fmt.Println("this is > 10: ", i) },
			Filter[Int](ctx,
				func(i Int) bool { return i.number > 10 },
				ForEachContd[Int](ctx,
					func(t Int) { fmt.Println(t) },
					Stream[Int](ctx, []Int{{12}, {21}, {3}, {40}}),
				),
			),
		),
	)
	fmt.Println(ct)
}

type Int struct {
	number int
}
