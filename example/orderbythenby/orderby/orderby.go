package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see OrderByEx1 example from Enumerable.OrderBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby

type pet struct {
	name string
	age  int
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
	}
	query := SlicepetToEnumerable(pets).
		OrderByMust(func(e1, e2 common.Elem) bool { return e1.(pet).age < e2.(pet).age }).
		Enumerable()
	for query.MoveNext() {
		p := query.Current().(pet)
		fmt.Printf("%s - %d\n", p.name, p.age)
	}
}
