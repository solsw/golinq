package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// see AllEx2 example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples

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
			pets: []pet{
				pet{name: "Belle", age: 8},
			},
		},
		person{
			lastName: "Philips",
			pets: []pet{
				pet{name: "Sweetie", age: 2},
				pet{name: "Rover", age: 13},
			},
		},
	}
	// Determine which people have pets that are all older than 5.
	names := SlicepersonToEnumerable(people).
		WhereMust(func(e common.Elem) bool {
			return SlicepetToEnumerable(e.(person).pets).
				AllMust(func(e common.Elem) bool { return e.(pet).age > 5 })
		}).
		SelectMust(func(e common.Elem) common.Elem { return e.(person).lastName })
	for names.MoveNext() {
		fmt.Println(names.Current())
	}
}
