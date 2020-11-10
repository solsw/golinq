package main

import (
	"fmt"
)

// see DefaultIfEmptyEx2 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type pet struct {
	name string
	age  int
}

func main() {
	defaultPet := pet{name: "Default Pet", age: 0}
	pets1 := []pet{
		pet{name: "Barley", age: 8},
		pet{name: "Boots", age: 4},
		pet{name: "Whiskers", age: 1},
	}
	en1 := SlicepetToEnumerable(pets1).DefaultIfEmptyDef(defaultPet)
	for en1.MoveNext() {
		fmt.Printf("Name: %s\n", en1.Current().(pet).name)
	}
	pets2 := []pet{}
	en2 := SlicepetToEnumerable(pets2).DefaultIfEmptyDef(defaultPet)
	for en2.MoveNext() {
		fmt.Printf("\nName: %s\n", en2.Current().(pet).name)
	}
}
