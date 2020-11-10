package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Take help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.take#examples

func main() {
	grades := enumerable.NewElems(59, 82, 70, 56, 92, 98, 85)
	topThreeGrades := grades.
		OrderByDescendingMust(typed.Lessint).
		Enumerable().
		Take(3)
	fmt.Println("The top three grades are:")
	for topThreeGrades.MoveNext() {
		fmt.Println(topThreeGrades.Current())
	}
}
