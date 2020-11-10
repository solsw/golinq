package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see LookupExample example from IOrderedEnumerable Interface help
// https://docs.microsoft.com/dotnet/api/system.linq.iorderedenumerable-1#examples

func main() {
	fruits := enumerable.NewElems("apricot", "orange", "banana", "mango", "apple", "grape", "strawberry")
	// Sort the strings first by their length and then alphabetically by passing the identity selector function.
	sortedFruits1 := fruits.
		OrderBySelMust(typed.Lessint, func(fruit common.Elem) common.Elem { return len(fruit.(string)) }).
		ThenByMust(typed.Lessstring).
		Enumerable()
	for sortedFruits1.MoveNext() {
		fmt.Println(sortedFruits1.Current())
	}
}
