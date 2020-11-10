package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see MinEx4 example from Enumerable.Min help
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
	min := pets.MinSelMust(typed.Lessint, func(e common.Elem) common.Elem { return e.(pet).age })
	fmt.Printf("The youngest animal is age %d.\n", min)
}
