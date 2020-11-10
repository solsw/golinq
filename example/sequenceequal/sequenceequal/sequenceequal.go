package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see SequenceEqualEx1 example from Enumerable.SequenceEqual help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

type pet struct {
	name string
	age  int
}

func main() {
	pet1 := pet{name: "Turbo", age: 2}
	pet2 := pet{name: "Peanut", age: 8}
	pets1 := enumerable.NewElems(pet1, pet2)
	pets2 := enumerable.NewElems(pet1, pet2)
	equal := pets1.SequenceEqual(pets2)
	var what string
	if equal {
		what = "are"
	} else {
		what = "are not"
	}
	fmt.Printf("The lists %s equal.\n", what)
}
