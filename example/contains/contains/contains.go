package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.Contains help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit := "mango"
	hasMango := fruits.Contains(fruit)
	var what string
	if hasMango {
		what = "does"
	} else {
		what = "does not"
	}
	fmt.Printf("The array %s contain '%s'.\n", what, fruit)
}
