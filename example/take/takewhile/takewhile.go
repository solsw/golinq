package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.TakeWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	query := fruits.TakeWhileMust(func(fruit common.Elem) bool {
		return typed.Cmpstring("orange", fruit.(string)) != 0
	})
	for query.MoveNext() {
		fmt.Println(query.Current())
	}
}
