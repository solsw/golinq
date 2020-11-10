package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see SelectManyEx3 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

type petOwner struct {
	name string
	pets []string
}

type ownerAndPet struct {
	owner petOwner
	pet   string
}

type ownerNameAndPetName struct {
	owner string
	pet   string
}

func main() {
	petOwners := []petOwner{
		petOwner{name: "Higa", pets: []string{"Scruffy", "Sam"}},
		petOwner{name: "Ashkenazi", pets: []string{"Walker", "Sugar"}},
		petOwner{name: "Price", pets: []string{"Scratches", "Diesel"}},
		petOwner{name: "Hines", pets: []string{"Dusty"}},
	}
	en := enumerable.NewElems(SlicepetOwnerToSlice(petOwners)...)
	query := en.
		SelectManyCollMust(
			func(e common.Elem) *enumerable.Enumerable {
				return enumerable.NewElems(typed.SlicestringToSlice(e.(petOwner).pets)...)
			},
			func(e1, e2 common.Elem) common.Elem {
				return ownerAndPet{e1.(petOwner), e2.(string)}
			},
		).
		WhereMust(func(e common.Elem) bool {
			return strings.HasPrefix(e.(ownerAndPet).pet, "S")
		}).
		SelectMust(func(e common.Elem) common.Elem {
			return ownerNameAndPetName{e.(ownerAndPet).owner.name, e.(ownerAndPet).pet}
		})
	vv, _ := EnumerableToSliceownerNameAndPetName(query)
	for _, v := range vv {
		fmt.Printf("%+v\n", v)
	}
}
