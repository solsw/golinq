package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see second example from Enumerable.Concat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples

type pet struct {
	name string
	age  int
}

func main() {
	cats := []pet{
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
	}
	enCats := SlicepetToEnumerable(cats)
	dogs := []pet{
		pet{name: "Bounder", age: 3},
		pet{name: "Snoopy", age: 14},
		pet{name: "Fido", age: 9},
	}
	enDogs := SlicepetToEnumerable(dogs)
	query := typed.SliceEnumerableToEnumerable([]*enumerable.Enumerable{
		enCats.SelectMust(func(e common.Elem) common.Elem { return e.(pet).name }),
		enDogs.SelectMust(func(e common.Elem) common.Elem { return e.(pet).name }),
	}).
		SelectManyMust(func(e common.Elem) *enumerable.Enumerable {
			return e.(*enumerable.Enumerable)
		})
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
