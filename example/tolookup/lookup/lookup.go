package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see LookupExample example from Lookup Class help
// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2#examples

type packag struct {
	company        string
	weight         float64
	trackingNumber int64
}

func main() {
	// Create a list of Packages to put into a Lookup data structure.
	packags := []packag{
		packag{company: "Coho Vineyard", weight: 25.2, trackingNumber: 89453312},
		packag{company: "Lucerne Publishing", weight: 18.7, trackingNumber: 89112755},
		packag{company: "Wingtip Toys", weight: 6.0, trackingNumber: 299456122},
		packag{company: "Contoso Pharmaceuticals", weight: 9.3, trackingNumber: 670053128},
		packag{company: "Wide World Importers", weight: 33.8, trackingNumber: 4665518773},
	}
	// Create a Lookup to organize the packages.
	// Use the first character of Company as the key value.
	// Select Company appended to TrackingNumber for each element value in the Lookup.
	lookup := SlicepackagToEnumerable(packags).ToLookupSelMust(
		func(e common.Elem) common.Elem {
			return []rune(e.(packag).company)[0]
		},
		func(e common.Elem) common.Elem {
			p := e.(packag)
			return fmt.Sprintf("%s %d", p.company, p.trackingNumber)
		},
	)
	// Iterate through each Grouping in the Lookup and output the contents.
	for _, eg := range lookup.Slice() {
		g := eg.(enumerable.Grouping)
		// Print the key value of the Grouping.
		fmt.Println(string(g.Key().(rune)))
		// Iterate through each value in the Grouping and print its value.
		for _, es := range g.Slice() {
			fmt.Printf("    %s\n", es.(string))
		}
	}
	// Get the number of key-collection pairs in the Lookup.
	count := lookup.Count()
	fmt.Printf("\n%d\n", count)
	// Select a collection of Packages by indexing directly into the Lookup.
	cgroup, _ := typed.EnumerableToSlicestring(lookup.Item('C'))
	// Output the results.
	fmt.Println("\nPackages that have a key of 'C':")
	for _, str := range cgroup {
		fmt.Println(str)
	}
	hasG := lookup.Contains('G')
	fmt.Printf("\n%t\n", hasG)
}
