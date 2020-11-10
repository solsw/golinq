package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits1 := enumerable.NewElems("orange")
	fruit1 := fruits1.SingleMust()
	fmt.Println(fruit1)
}
