package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.Union help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

func main() {
	ints1 := enumerable.NewElems(5, 3, 9, 7, 5, 9, 3, 7)
	ints2 := enumerable.NewElems(8, 3, 6, 4, 4, 9, 1, 0)
	union := ints1.Union(ints2)
	for union.MoveNext() {
		fmt.Printf("%d ", union.Current())
	}
	fmt.Println()
}
