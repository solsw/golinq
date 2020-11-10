package main

import (
	"fmt"

	"github.com/solsw/golinq/typed"
)

// see last example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

func main() {
	numbers := typed.SliceintToEnumerable([]int{}).DefaultIfEmpty()
	for numbers.MoveNext() {
		i, _ := typed.ElemToint(numbers.Current())
		fmt.Println(i)
	}
}
