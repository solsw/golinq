package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	grades := enumerable.NewElems(78, 92, 100, 37, 81)
	average := grades.AverageIntMust(func(e common.Elem) int { return e.(int) })
	fmt.Printf("The average grade is %g.\n", average)
}
