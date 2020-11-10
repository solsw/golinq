package main

import (
	"fmt"

	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Count help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count

func main() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	en := typed.SlicestringToEnumerable(fruits)
	numberOfFruits := en.Count()
	fmt.Printf("There are %d fruits in the collection.\n", numberOfFruits)
}
