package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits1 := enumerable.NewElems("orange")
	fruit1 := fruits1.SingleOrDefaultMust()
	fmt.Println(fruit1)
}
