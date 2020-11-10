package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func main() {
	amounts := enumerable.NewElems(5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500)
	query := amounts.
		SkipWhileIdxMust(func(amount common.Elem, index int) bool { return amount.(int) > index*1000 })
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
