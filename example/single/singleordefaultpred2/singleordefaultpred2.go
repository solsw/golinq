package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit2, _ := typed.ElemTostring(fruits.SingleOrDefaultPredMust(func(e common.Elem) bool { return len(e.(string)) > 15 }))
	var what string
	if fruit2 == "" {
		what = "No such string!"
	} else {
		what = fruit2
	}
	fmt.Println(what)
}
