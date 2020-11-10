package main

import (
	"fmt"
	"math"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see OrderByDescendingEx1 example from Enumerable.OrderByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending

func main() {
	decimals := enumerable.NewElems(6.2, 8.3, 0.5, 1.3, 6.3, 9.7)
	query := decimals.
		OrderByDescendingSelMust(typed.Lessfloat64,
			func(e common.Elem) common.Elem {
				_, fr := math.Modf(e.(float64))
				return fr
			}).
		Enumerable()
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
