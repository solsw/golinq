package main

import (
	"fmt"
)

// see DefaultIfEmptyEx1 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type pet struct {
	name string
	age  int
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
	}
	en := SlicepetToEnumerable(pets).DefaultIfEmpty()
	for en.MoveNext() {
		fmt.Printf("Name: %s\n", en.Current().(pet).name)
	}
}
