package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see LongCountEx2 example from Enumerable.LongCount help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

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
	en := SlicepetToEnumerable(pets)
	const age = 3
	count := en.CountPredMust(func(e common.Elem) bool { return e.(pet).age > age })
	fmt.Printf("There are %d animals over age %d.\n", count, age)
}
