package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.ThenBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby

func main() {
	fruits := enumerable.NewElems("grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry")
	// Sort the strings first by their length and then alphabetically by passing the identity selector function.
	query := fruits.
		OrderBySelMust(typed.Lessint, func(fruit common.Elem) common.Elem { return len(fruit.(string)) }).
		ThenByMust(typed.Lessstring).
		Enumerable()
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
