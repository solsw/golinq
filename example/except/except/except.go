package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see first example from Enumerable.Except help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

func main() {
	numbers1 := enumerable.NewElems(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	numbers2 := enumerable.NewElems(2.2)
	onlyInFirstSet := numbers1.Except(numbers2)
	for onlyInFirstSet.MoveNext() {
		fmt.Println(onlyInFirstSet.Current())
	}
}
