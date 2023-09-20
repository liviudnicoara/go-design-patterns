package main

import (
	"fmt"
	"os"

	"github.com/liviudnicoara/go-design-patterns/9.Prototype/page"
)

// The prototype design pattern involves creating objects based on a template of an existing object through cloning.
// It is especially useful when the construction of a new object is inefficient or complex.
// In the prototype pattern, an instance of actual object (which is created by expensive database operation) is cloned
// to create a new object.
// This pattern provides a mechanism to copy the original object to a new object and then modify it according to our needs.

func main() {

	prototype := &page.PaymentPage{
		Provider:    "MasterCard",
		Token:       "eyJhbGciOiJIUzI1NiJ...",
		Certificate: *os.NewFile(1, "certificate"),
	}

	clone := prototype.Clone()

	fmt.Printf("Prototype. Type: %T. Pointer: %v. Value: %v\n", prototype, &prototype, prototype)
	fmt.Printf("Clone. Type: %T. Pointer: %v. Value: %v\n", clone, &clone, clone)
}
