package main

import (
	"fmt"

	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	pageNumbers := []int{}
	// Setting the default value to 1 after the query.
	pageNumber1, _ := typed.ElemToint(typed.SliceintToEnumerable(pageNumbers).SingleOrDefaultMust())
	if pageNumber1 == 0 {
		pageNumber1 = 1
	}
	fmt.Printf("The value of the pageNumber1 variable is %d\n", pageNumber1)
	// Setting the default value to 1 by using DefaultIfEmpty() in the query.
	pageNumber2, _ := typed.ElemToint(typed.SliceintToEnumerable(pageNumbers).DefaultIfEmptyDef(1).SingleMust())
	fmt.Printf("The value of the pageNumber2 variable is %d\n", pageNumber2)
}
