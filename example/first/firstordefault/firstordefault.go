package main

import (
	"fmt"

	"github.com/solsw/golinq/typed"
)

// see second example from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

func main() {
	numbers := typed.SliceintToEnumerable([]int{})
	first, _ := typed.ElemToint(numbers.FirstOrDefault())
	fmt.Println(first)

	months := []int{}
	// Setting the default value to 1 after the query.
	firstMonth1, _ := typed.ElemToint(typed.SliceintToEnumerable(months).FirstOrDefault())
	if firstMonth1 == 0 {
		firstMonth1 = 1
	}
	fmt.Printf("The value of the firstMonth1 variable is %d\n", firstMonth1)
	// Setting the default value to 1 by using DefaultIfEmpty() in the query.
	firstMonth2, _ := typed.ElemToint(typed.SliceintToEnumerable(months).DefaultIfEmptyDef(1).FirstMust())
	fmt.Printf("The value of the firstMonth2 variable is %d\n", firstMonth2)
}
