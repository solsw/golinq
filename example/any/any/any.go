package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

func main() {
	numbers := enumerable.NewElems(1, 2)
	hasElements := numbers.Any()
	var what string
	if hasElements {
		what = "is not"
	} else {
		what = "is"
	}
	fmt.Printf("The list %s empty.\n", what)
}
