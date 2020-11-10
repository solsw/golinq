package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see GroupJoinEx1 example from Enumerable.GroupJoin help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin

type person struct {
	name string
}

type pet struct {
	name  string
	owner person
}

type ownerAndPets struct {
	ownername string
	pets      *enumerable.Enumerable
}

func main() {
	magnus := person{name: "Hedlund, Magnus"}
	terry := person{name: "Adams, Terry"}
	charlotte := person{name: "Weiss, Charlotte"}

	barley := pet{name: "Barley", owner: terry}
	boots := pet{name: "Boots", owner: terry}
	whiskers := pet{name: "Whiskers", owner: charlotte}
	daisy := pet{name: "Daisy", owner: magnus}

	people := enumerable.NewElems(magnus, terry, charlotte)
	pets := enumerable.NewElems(barley, boots, whiskers, daisy)

	query := people.GroupJoinMust(pets,
		common.Identity,
		func(e common.Elem) common.Elem {
			return e.(pet).owner
		},
		func(prsn common.Elem, pts *enumerable.Enumerable) common.Elem {
			return ownerAndPets{
				ownername: prsn.(person).name,
				pets:      pts.SelectMust(func(pt common.Elem) common.Elem { return pt.(pet).name })}
		})
	for query.MoveNext() {
		ops := query.Current().(ownerAndPets)
		fmt.Printf("%s:\n", ops.ownername)
		for ops.pets.MoveNext() {
			fmt.Printf("  %s\n", ops.pets.Current())
		}
	}
}
