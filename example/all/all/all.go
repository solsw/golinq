package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
)

// see AllEx example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples

type pet struct {
	name string
	age  int
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 10},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 6},
	}
	// Determine whether all pet names in the array start with 'B'.
	allStartWithB := SlicepetToEnumerable(pets).
		AllMust(func(e common.Elem) bool { return strings.HasPrefix(e.(pet).name, "B") })
	var what string
	if allStartWithB {
		what = "All"
	} else {
		what = "Not all"
	}
	fmt.Printf("%s pet names start with 'B'.\n", what)
}
