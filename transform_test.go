package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	ctx := context.Background()
	list := []Int{{1}, {2}, {3}, {4}}

	filteredStream := New[Int](ctx, list).
		Filter(func(i Int) bool { return i.number < 4 })

	filteredFloatList := TransformFn[Int, Float](func(i Int) Float { return Float{float64(i.number) * 10.1} }, filteredStream).
		ForEachContd(func(float Float) { fmt.Println(float) }).Collect()

	fmt.Println(filteredFloatList)
}

type Float struct {
	number float64
}
