package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see MinEx3 example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min

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
	min := pets.MinMust(func(e1, e2 common.Elem) bool { return typed.Lessint(e1.(pet).age, e2.(pet).age) })
	fmt.Printf("The 'minimum' animal is %s.\n", min.(pet).name)
}
