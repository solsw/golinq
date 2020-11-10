package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see second example from Enumerable.Contains help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

type product struct {
	name string
	code int
}

var productEq = func(e1, e2 common.Elem) bool {
	p1 := e1.(product)
	p2 := e2.(product)
	return p1.code == p2.code && strings.ToLower(p1.name) == strings.ToLower(p2.name)
}

func main() {
	fruits := enumerable.NewElems(
		product{name: "apple", code: 9},
		product{name: "orange", code: 4},
		product{name: "lemon", code: 12},
	)
	apple := product{name: "Apple", code: 9}
	kiwi := product{name: "kiwi", code: 8}
	hasApple := fruits.ContainsEq(apple, productEq)
	hasKiwi := fruits.ContainsEq(kiwi, productEq)
	fmt.Printf("Apple? %t\n", hasApple)
	fmt.Printf("Kiwi? %t\n", hasKiwi)
}
