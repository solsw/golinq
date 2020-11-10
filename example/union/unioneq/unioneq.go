package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
)

// see last two example from Enumerable.Union help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

type product struct {
	name string
	code int
}

func main() {
	store1 := SliceproductToEnumerable([]product{
		product{name: "APPLE", code: 9},
		product{name: "orange", code: 4},
	})
	store2 := SliceproductToEnumerable([]product{
		product{name: "apple", code: 9},
		product{name: "lemon", code: 12},
	})
	//Get the products from the both arrays excluding duplicates.
	union := store1.UnionEq(store2,
		func(e1, e2 common.Elem) bool {
			p1 := e1.(product)
			p2 := e2.(product)
			return p1.code == p2.code && strings.ToUpper(p1.name) == strings.ToUpper(p2.name)
		})
	for union.MoveNext() {
		p := union.Current().(product)
		fmt.Printf("%s %d\n", p.name, p.code)
	}
}
