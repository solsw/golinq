package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits2 := enumerable.NewElems()
	fruit2, _ := typed.ElemTostring(fruits2.SingleOrDefaultMust())
	var what string
	if fruit2 == "" {
		what = "No such string!"
	} else {
		what = fruit2
	}
	fmt.Println(what)
}
