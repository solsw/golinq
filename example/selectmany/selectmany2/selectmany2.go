package main

import (
	"fmt"
	"strconv"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see SelectManyEx2 example from Enumerable.SelectMany help
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
		petOwner{name: "Hines, Patrick", pets: []string{"Dusty"}},
	}
	en := enumerable.NewElems(SlicepetOwnerToSlice(petOwners)...)
	query := en.SelectManyIdxMust(
		func(e common.Elem, i int) *enumerable.Enumerable {
			return enumerable.NewElems(typed.SlicestringToSlice(e.(petOwner).pets)...).
				SelectMust(func(e common.Elem) common.Elem { return strconv.Itoa(i) + e.(string) })
		})
	ss, _ := typed.EnumerableToSlicestring(query)
	for _, s := range ss {
		fmt.Println(s)
	}
}
