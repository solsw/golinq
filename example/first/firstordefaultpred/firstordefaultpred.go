package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see first example from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

func main() {
	names := enumerable.NewElems("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	firstLongName := names.FirstOrDefaultPredMust(func(e common.Elem) bool { return len(e.(string)) > 20 })
	fmt.Printf("The first long name is '%s'.\n", firstLongName)
	names.Reset()
	firstVeryLongName, _ := typed.ElemTostring(
		names.FirstOrDefaultPredMust(func(e common.Elem) bool { return len(e.(string)) > 30 }))
	var what string
	if firstVeryLongName == "" {
		what = "not a"
	} else {
		what = "a"
	}
	fmt.Printf("There is %s name longer than 30 characters.\n", what)
}
