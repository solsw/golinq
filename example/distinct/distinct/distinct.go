package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.Distinct help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

func main() {
	ages := enumerable.NewElems(21, 46, 46, 55, 17, 21, 55, 55)
	distinctAges := ages.Distinct()
	fmt.Println("Distinct ages:")
	for distinctAges.MoveNext() {
		fmt.Println(distinctAges.Current())
	}
}
