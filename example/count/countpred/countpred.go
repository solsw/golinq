package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see CountEx2 example from Enumerable.Count help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count

type pet struct {
	name       string
	vaccinated bool
}

func main() {
	pets := []pet{
		pet{name: "Barley", vaccinated: true},
		pet{name: "Boots", vaccinated: false},
		pet{name: "Whiskers", vaccinated: false},
	}
	en := SlicepetToEnumerable(pets)
	numberUnvaccinated := en.CountPredMust(func(e common.Elem) bool { return !e.(pet).vaccinated })
	fmt.Printf("There are %d unvaccinated animals.\n", numberUnvaccinated)
}
