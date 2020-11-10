package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Intersect help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

func main() {
	id1 := enumerable.NewElems(44, 26, 92, 30, 71, 38)
	id2 := enumerable.NewElems(39, 59, 83, 47, 26, 4, 30)
	both := id1.Intersect(id2)
	for both.MoveNext() {
		fmt.Println(both.Current())
	}
}
