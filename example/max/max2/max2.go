package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see MaxEx3 example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

type pet struct {
	name string
	age  int
}

func main() {
	pets := enumerable.NewElems(
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
	)
	max := pets.MaxMust(func(e1, e2 common.Elem) bool {
		p1 := e1.(pet)
		p2 := e2.(pet)
		return typed.Lessint(p1.age+len(p1.name), p2.age+len(p2.name))
	})
	fmt.Printf("The 'maximum' animal is %s.\n", max.(pet).name)
}
