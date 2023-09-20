package main

// The builder design pattern. It is a creational pattern that is used to construct a complex object step by step.
// It separates the construction of an object from its representation so that the same construction process can
// create different representations.
// When you are building a complex object that has multiple parts or configuration options, creating and
// configuring the object directly within client code or via a constructor can become unwieldy and hard to read.
// The Builder pattern addresses this issue by providing a way to build the complex object step by step, with each step
// clearly encapsulated by a method in the builder.

import (
	"fmt"

	"github.com/liviudnicoara/go-design-patterns/2.Builder/infra"
)

func main() {
	builder := infra.NewHTTPRequestBuilder().
		WithMethod("POST").
		WithURL("http://test.com").
		WithHeader("Content-Type", "application/json").
		WithBody(`{"test": "value"}`)

	req, err := builder.Build()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *req)
}
