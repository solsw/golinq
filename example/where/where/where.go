package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	en := typed.SlicestringToEnumerable(fruits)
	query := en.WhereMust(func(e common.Elem) bool { return len(e.(string)) < 6 })
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
