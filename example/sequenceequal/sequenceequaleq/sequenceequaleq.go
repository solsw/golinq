package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see last example from Enumerable.SequenceEqual help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

type product struct {
	name string
	code int
}

func main() {
	storeA := enumerable.NewElems(
		product{name: "Apple", code: 9},
		product{name: "orange", code: 4})
	storeB := enumerable.NewElems(
		product{name: "applE", code: 9},
		product{name: "orange", code: 4})
	equalAB := storeA.SequenceEqualEq(storeB, func(e1, e2 common.Elem) bool {
		p1 := e1.(product)
		p2 := e2.(product)
		return p1.code == p2.code && strings.ToLower(p1.name) == strings.ToLower(p2.name)
	})
	fmt.Printf("Equal? %t\n", equalAB)
}
