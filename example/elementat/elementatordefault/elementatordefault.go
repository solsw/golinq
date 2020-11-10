package main

import (
	"fmt"

	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.ElementAtOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault#examples

func main() {
	names := enumerable.NewElems("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	index := 20
	name, _ := typed.ElemTostring(names.ElementAtOrDefault(index))
	var what string
	if name == "" {
		what = "<no name at this index>"
	} else {
		what = name
	}
	fmt.Printf("The name chosen at index %d is '%s'.\n", index, what)
}
