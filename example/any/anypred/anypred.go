package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see AnyEx3 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

type pet struct {
	name       string
	age        int
	vaccinated bool
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 8, vaccinated: true},
		pet{name: "Boots", age: 4, vaccinated: false},
		pet{name: "Whiskers", age: 1, vaccinated: false},
	}
	// Determine whether any pets over age 1 are also unvaccinated.
	unvaccinated := SlicepetToEnumerable(pets).
		AnyPredMust(func(e common.Elem) bool {
			p := e.(pet)
			return p.age > 1 && p.vaccinated == false
		})
	var what string
	if unvaccinated {
		what = "are"
	} else {
		what = "are not any"
	}
	fmt.Printf("There %s unvaccinated animals over age one.\n", what)
}
