package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see AnyEx2 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

type pet struct {
	name string
	age  int
}

type person struct {
	lastName string
	pets     []pet
}

func main() {
	people := []person{
		person{
			lastName: "Haas",
			pets: []pet{
				pet{name: "Barley", age: 10},
				pet{name: "Boots", age: 14},
				pet{name: "Whiskers", age: 6},
			},
		},
		person{
			lastName: "Fakhouri",
			pets: []pet{
				pet{name: "Snowball", age: 1},
			},
		},
		person{
			lastName: "Antebi",
			pets:     []pet{},
		},
		person{
			lastName: "Philips",
			pets: []pet{
				pet{name: "Sweetie", age: 2},
				pet{name: "Rover", age: 13},
			},
		},
	}
	// Determine which people have a non-empty Pet array.
	names := SlicepersonToEnumerable(people).
		WhereMust(func(e common.Elem) bool {
			return SlicepetToEnumerable(e.(person).pets).Any()
		}).
		SelectMust(func(e common.Elem) common.Elem { return e.(person).lastName })
	for names.MoveNext() {
		fmt.Println(names.Current())
	}
}
