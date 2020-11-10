package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see MaxEx4 example from Enumerable.Max help
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
	max := pets.MaxSelMust(typed.Lessint,
		func(e common.Elem) common.Elem { return e.(pet).age + len(e.(pet).name) })
	fmt.Printf("The maximum pet age plus name length is %d.\n", max)
}
