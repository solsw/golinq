package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	average := fruits.AverageIntMust(func(e common.Elem) int { return len(e.(string)) })
	fmt.Printf("The average string length is %g.\n", average)
}
