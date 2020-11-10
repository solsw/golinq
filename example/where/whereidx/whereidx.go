package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	en := typed.SliceintToEnumerable(numbers)
	query := en.WhereIdxMust(func(e common.Elem, i int) bool { return e.(int) <= i*10 })
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
