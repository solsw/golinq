package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Zip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

func main() {
	numbers := enumerable.NewElems(1, 2, 3, 4)
	words := enumerable.NewElems("one", "two", "three")
	numbersAndWords := numbers.ZipMust(words,
		func(first, second common.Elem) common.Elem { return fmt.Sprintf("%d %s", first, second) })
	for numbersAndWords.MoveNext() {
		fmt.Println(numbersAndWords.Current())
	}
}
