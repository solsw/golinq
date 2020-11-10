package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/errors"
)

// see example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits2 := enumerable.NewElems("orange", "apple")
	fruit2, err := fruits2.Single()
	if err == errors.MultiElems {
		fmt.Println("The collection does not contain exactly one element.")
	} else {
		fmt.Println(fruit2)
	}
}
