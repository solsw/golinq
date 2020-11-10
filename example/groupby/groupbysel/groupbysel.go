package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see GroupByEx1 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

type pet struct {
	name string
	age  int
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
		pet{name: "Daisy", age: 4},
	}
	// Group the pets using Age as the key value and selecting only the pet's Name for each value.
	query := SlicepetToEnumerable(pets).GroupBySelMust(
		func(e common.Elem) common.Elem { return e.(pet).age },
		func(e common.Elem) common.Elem { return e.(pet).name })
	for query.MoveNext() {
		g := query.Current().(enumerable.Grouping)
		fmt.Println(g.Key())
		for _, n := range g.Slice() {
			fmt.Printf("  %s\n", n)
		}
	}
}
