package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.ElementAt help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat#examples

func main() {
	names := []string{"Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	rand.Seed(time.Now().UnixNano())
	name := typed.SlicestringToEnumerable(names).ElementAtMust(rand.Intn(len(names)))
	fmt.Printf("The name chosen at random is '%s'.\n", name)
}
