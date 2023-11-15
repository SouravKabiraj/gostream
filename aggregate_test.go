package gostream

import (
	"context"
	"fmt"
	"testing"
)

func TestExampleAggregate(t *testing.T) {
	ctx := context.Background()
	stream := Stream[Person](ctx, []Person{{3}, {1}, {1}, {6}})
	totalAgePersion := Aggregate[Person, Person](ctx, func(output *Person, person Person) {
		output.age = output.age + person.age
	}, stream)

	fmt.Println(totalAgePersion)
}

type Person struct {
	age int
}
