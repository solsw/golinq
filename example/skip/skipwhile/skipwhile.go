package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func main() {
	grades := enumerable.NewElems(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := grades.
		OrderByDescendingMust(typed.Lessint).
		Enumerable().
		SkipWhileMust(func(grade common.Elem) bool { return grade.(int) >= 80 })
	fmt.Println("All grades below 80:")
	for lowerGrades.MoveNext() {
		fmt.Println(lowerGrades.Current())
	}
}
