package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Sum help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

func main() {
	numbers := enumerable.NewElems(43.68, 1.25, 583.7, 6.5)
	sum := numbers.SumFloat64Must(func(e common.Elem) float64 { return e.(float64) })
	fmt.Printf("The sum of the numbers is %g.\n", sum)
}
