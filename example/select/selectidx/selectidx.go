package main

import (
	"fmt"

	"github.com/solsw/gohelpers/stringhelper"
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Select help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

type indexstr struct {
	index int
	str   string
}

func main() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	en := typed.SlicestringToEnumerable(fruits)
	query := en.SelectIdxMust(func(e common.Elem, i int) common.Elem {
		s, _ := stringhelper.SubstrBeg(e.(string), i)
		return indexstr{index: i, str: s}
	})
	for query.MoveNext() {
		fmt.Printf("%+v\n", query.Current())
	}
}
