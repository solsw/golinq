package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see SelectManyEx1 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

type petOwner struct {
	name string
	pets []string
}

func main() {
	petOwners := []petOwner{
		petOwner{name: "Higa, Sidney", pets: []string{"Scruffy", "Sam"}},
		petOwner{name: "Ashkenazi, Ronen", pets: []string{"Walker", "Sugar"}},
		petOwner{name: "Price, Vernette", pets: []string{"Scratches", "Diesel"}},
	}

	en1 := enumerable.NewElems(SlicepetOwnerToSlice(petOwners)...)
	query1 := en1.SelectManyMust(
		func(e common.Elem) *enumerable.Enumerable {
			return enumerable.NewElems(typed.SlicestringToSlice(e.(petOwner).pets)...)
		})
	fmt.Println("Using SelectMany():")
	ss1, _ := typed.EnumerableToSlicestring(query1)
	for _, s := range ss1 {
		fmt.Println(s)
	}

	en2 := enumerable.NewElems(SlicepetOwnerToSlice(petOwners)...)
	query2 := en2.SelectMust(func(e common.Elem) common.Elem {
		return enumerable.NewElems(typed.SlicestringToSlice(e.(petOwner).pets)...)
	})
	fmt.Println("\nUsing Select():")
	for query2.MoveNext() {
		ss2, _ := typed.EnumerableToSlicestring(query2.Current().(*enumerable.Enumerable))
		for _, s := range ss2 {
			fmt.Println(s)
		}
		fmt.Println()
	}
}
