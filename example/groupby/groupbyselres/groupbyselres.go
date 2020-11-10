package main

import (
	"fmt"
	"math"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see GroupByEx4 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

type pet struct {
	name string
	age  float64
}

type result struct {
	key      float64
	count    int
	min, max float64
}

func main() {
	pets := []pet{
		pet{name: "Barley", age: 8.3},
		pet{name: "Boots", age: 4.9},
		pet{name: "Whiskers", age: 1.5},
		pet{name: "Daisy", age: 4.3},
	}
	query := SlicepetToEnumerable(pets).GroupBySelResMust(
		func(e common.Elem) common.Elem { return math.Floor(e.(pet).age) },
		func(e common.Elem) common.Elem { return e.(pet).age },
		func(baseAge common.Elem, ages *enumerable.Enumerable) common.Elem {
			c := ages.Count()
			ages.Reset()
			mn := ages.MinMust(typed.Lessfloat64)
			ages.Reset()
			mx := ages.MaxMust(typed.Lessfloat64)
			return result{key: baseAge.(float64), count: c, min: mn.(float64), max: mx.(float64)}
		})
	for query.MoveNext() {
		r := query.Current().(result)
		fmt.Printf("\nAge group: %g\n", r.key)
		fmt.Printf("Number of pets in this age group: %d\n", r.count)
		fmt.Printf("Minimum age: %g\n", r.min)
		fmt.Printf("Maximum age: %g\n", r.max)
	}
}
