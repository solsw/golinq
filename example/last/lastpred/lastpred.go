package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see second example from Enumerable.Last help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last

func main() {
	numbers := enumerable.NewElems(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19)
	last := numbers.LastPredMust(func(e common.Elem) bool { return e.(int) > 80 })
	fmt.Println(last)
}
