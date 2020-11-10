package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Skip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip#examples

func main() {
	grades := enumerable.NewElems(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := grades.
		OrderByDescendingMust(typed.Lessint).
		Enumerable().
		Skip(3)
	fmt.Println("All grades except the top three are:")
	for lowerGrades.MoveNext() {
		fmt.Println(lowerGrades.Current())
	}
}
