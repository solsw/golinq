package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.TakeWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func main() {
	fruits := enumerable.NewElems("apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry")
	query := fruits.TakeWhileIdxMust(func(fruit common.Elem, index int) bool {
		return len(fruit.(string)) >= index
	})
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
