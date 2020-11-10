package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Reverse help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse#examples

func main() {
	apple := enumerable.NewElems("a", "p", "p", "l", "e")
	reversed := apple.Reverse()
	for reversed.MoveNext() {
		fmt.Printf("%s ", reversed.Current())
	}
	fmt.Println()
}
