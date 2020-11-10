package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min

func main() {
	doubles := enumerable.NewElems(1.5e+104, 9e+103, -2e+103)
	min := doubles.MinMust(typed.Lessfloat64)
	fmt.Printf("The smallest number is %G.\n", min)
}
