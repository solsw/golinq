package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see ThenByDescendingEx1 example from Enumerable.ThenByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

func main() {
	fruits := enumerable.NewElems("apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE")
	// Sort the strings first ascending by their length and then descending using a custom case insensitive comparer.
	query := fruits.
		OrderBySelMust(typed.Lessint, func(fruit common.Elem) common.Elem { return len(fruit.(string)) }).
		ThenByDescendingMust(func(e1, e2 common.Elem) bool {
			return strings.ToUpper(e1.(string)) < strings.ToUpper(e2.(string))
		}).
		Enumerable()
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
