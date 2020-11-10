package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Repeat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat#examples

func main() {
	strings := enumerable.RepeatMust("I like programming.", 15)
	for strings.MoveNext() {
		fmt.Println(strings.Current())
	}
}
