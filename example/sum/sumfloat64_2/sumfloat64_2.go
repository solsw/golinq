package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see SumEx1 example from Enumerable.Sum help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

type packag struct {
	company string
	weight  float64
}

func main() {
	packages := enumerable.NewElems(
		packag{company: "Coho Vineyard", weight: 25.2},
		packag{company: "Lucerne Publishing", weight: 18.7},
		packag{company: "Wingtip Toys", weight: 6.0},
		packag{company: "Adventure Works", weight: 33.8},
	)
	totalWeight := packages.SumFloat64Must(func(e common.Elem) float64 { return e.(packag).weight })
	fmt.Printf("The total weight of the packages is: %g\n", totalWeight)
}
