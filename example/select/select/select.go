package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Select help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

func main() {
	en := enumerable.RangeMust(1, 10)
	squares := en.SelectMust(func(e common.Elem) common.Elem { i := e.(int); return i * i })
	for squares.MoveNext() {
		fmt.Println(squares.Current())
	}
}
