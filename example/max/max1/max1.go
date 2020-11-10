package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

func main() {
	longs := enumerable.NewElems(4294967296, 466855135, 81125)
	max := longs.MaxMust(typed.Lessint)
	fmt.Printf("The largest number is %d.\n", max)
}
