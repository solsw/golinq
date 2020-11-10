package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
)

// see last two examples from Enumerable.Distinct help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

type product struct {
	name string
	code int
}

func main() {
	products := []product{
		product{name: "apple", code: 9},
		product{name: "orange", code: 4},
		product{name: "Apple", code: 9},
		product{name: "lemon", code: 12},
	}
	//Exclude duplicates.
	noduplicates := SliceproductToEnumerable(products).DistinctEq(
		func(e1, e2 common.Elem) bool {
			p1 := e1.(product)
			p2 := e2.(product)
			return p1.code == p2.code && strings.ToUpper(p1.name) == strings.ToUpper(p2.name)
		})
	for noduplicates.MoveNext() {
		p := noduplicates.Current().(product)
		fmt.Printf("%s %d\n", p.name, p.code)
	}
}
