package main

import (
	"fmt"
	"strconv"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	numbers := enumerable.NewElems("10007", "37", "299846234235")
	average := numbers.AverageIntMust(func(e common.Elem) int {
		r, _ := strconv.Atoi(e.(string))
		return r
	})
	fmt.Printf("The average is %.f.\n", average)
}
