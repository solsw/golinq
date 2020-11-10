package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.First help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first

func main() {
	numbers := enumerable.NewElems(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19)
	first := numbers.FirstMust()
	fmt.Println(first)
}
