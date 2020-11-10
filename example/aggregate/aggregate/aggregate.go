package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/typed"
)

// see last example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	sentence := "the quick brown fox jumps over the lazy dog"
	// Split the string into individual words.
	words := strings.Fields(sentence)
	// Prepend each word to the beginning of the new sentence to reverse the word order.
	reversed := typed.SlicestringToEnumerable(words).AggregateMust(
		func(workingSentence interface{}, next common.Elem) interface{} {
			return next.(string) + " " + workingSentence.(string)
		},
	)
	fmt.Println(reversed)
}
