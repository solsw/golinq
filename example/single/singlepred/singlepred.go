package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit1 := fruits.SinglePredMust(func(e common.Elem) bool { return len(e.(string)) > 10 })
	fmt.Println(fruit1)
}
