package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	fruits := enumerable.NewElems("apple", "mango", "orange", "passionfruit", "grape")
	// Determine whether any string in the array is longer than "banana".
	longestName := fruits.AggregateSeedSelMust("banana",
		func(longest interface{}, next common.Elem) interface{} {
			if len(next.(string)) > len(longest.(string)) {
				return next
			}
			return longest
		},
		// Return the final result as an upper case string.
		func(fruit interface{}) interface{} { return strings.ToUpper(fruit.(string)) },
	)
	fmt.Printf("The fruit with the longest name is %s.\n", longestName)
}
