package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see second example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	ints := enumerable.NewElems(4, 8, 8, 3, 9, 0, 7, 8, 2)
	// Count the even numbers in the array, using a seed value of 0.
	numEven := ints.AggregateSeedMust(0,
		func(total interface{}, next common.Elem) interface{} {
			if next.(int)%2 == 0 {
				return total.(int) + 1
			}
			return total.(int)
		},
	)
	fmt.Printf("The number of even integers is: %d.\n", numEven)
}
