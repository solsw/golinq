package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see JoinEx1 example from Enumerable.Join help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join

type person struct {
	name string
}

type pet struct {
	name  string
	owner person
}

type ownerAndPet struct {
	ownername string
	pet       string
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

	query := people.JoinMust(pets,
		common.Identity,
		func(e common.Elem) common.Elem { return e.(pet).owner },
		func(pr, pt common.Elem) common.Elem {
			return ownerAndPet{ownername: pr.(person).name, pet: pt.(pet).name}
		},
	)
	for query.MoveNext() {
		op := query.Current().(ownerAndPet)
		fmt.Printf("%s - %s\n", op.ownername, op.pet)
	}
}
