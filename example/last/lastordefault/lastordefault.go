package main

import (
	"fmt"

	"github.com/solsw/golinq/typed"
)

// see first two examples from Enumerable.LastOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

func main() {
	fruits := typed.SlicestringToEnumerable([]string{})
	last, _ := typed.ElemTostring(fruits.LastOrDefaultMust())
	if last == "" {
		fmt.Println("<string is empty>")
	} else {
		fmt.Println(last)
	}

	daysOfMonth := []int{}
	// Setting the default value to 1 after the query.
	lastDay1, _ := typed.ElemToint(typed.SliceintToEnumerable(daysOfMonth).LastOrDefaultMust())
	if lastDay1 == 0 {
		lastDay1 = 1
	}
	fmt.Printf("The value of the lastDay1 variable is %d\n", lastDay1)
	// Setting the default value to 1 by using DefaultIfEmpty() in the query.
	lastDay2, _ := typed.ElemToint(typed.SliceintToEnumerable(daysOfMonth).DefaultIfEmptyDef(1).LastMust())
	fmt.Printf("The value of the lastDay2 variable is %d\n", lastDay2)
}
