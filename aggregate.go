package gostream

import (
	"context"
)

func Aggregate[I any, O any](ctx context.Context, action func(*O, I), in <-chan I) *O {
	var aggregatedResult O
	for {
		select {
		case <-ctx.Done():
			return &aggregatedResult
		case item, ok := <-in:
			if !ok {
				return &aggregatedResult
			}
			action(&aggregatedResult, item)
		}
	}
}
