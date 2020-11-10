package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see ToDictionaryEx1 example from Enumerable.ToDictionary help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

type packag struct {
	company        string
	weight         float64
	trackingNumber int64
}

func main() {
	packags := []packag{
		packag{company: "Coho Vineyard", weight: 25.2, trackingNumber: 89453312},
		packag{company: "Lucerne Publishing", weight: 18.7, trackingNumber: 89112755},
		packag{company: "Wingtip Toys", weight: 6.0, trackingNumber: 299456122},
		packag{company: "Adventure Works", weight: 33.8, trackingNumber: 4665518773},
	}
	// Create a Dictionary of Package objects, using TrackingNumber as the key.
	dictionary := SlicepackagToEnumerable(packags).ToDictionaryMust(
		func(p common.Elem) common.Elem {
			return p.(packag).trackingNumber
		})
	den := dictionary.Enumerable()
	for den.MoveNext() {
		kvp := den.Current().(enumerable.KeyValue)
		p := kvp.Value().(packag)
		fmt.Printf("Key %d: %s, %g pounds\n", kvp.Key(), p.company, p.weight)
	}
}
