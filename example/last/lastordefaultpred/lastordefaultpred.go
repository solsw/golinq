package main

import (
	"fmt"
	"math"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see last example from Enumerable.LastOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

func main() {
	numbers := enumerable.NewElems(49.6, 52.3, 51.0, 49.4, 50.2, 48.3)
	last50 := numbers.LastOrDefaultPredMust(func(e common.Elem) bool { return math.Round(e.(float64)) == 50.0 })
	fmt.Printf("The last number that rounds to 50 is %g\n", last50)
	numbers.Reset()
	last40, _ := typed.ElemTofloat64(numbers.LastOrDefaultPredMust(func(e common.Elem) bool { return math.Round(e.(float64)) == 40.0 }))
	var what string
	if last40 == 0.0 {
		what = "<DOES NOT EXIST>"
	} else {
		what = fmt.Sprintf("%g", last40)
	}
	fmt.Printf("The last number that rounds to 40 is %s.\n", what)
}
