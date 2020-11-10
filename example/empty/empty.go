package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/typed"
)

// see example from Enumerable.Empty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.empty#examples

func main() {
	names1 := []string{"Hartono, Tommy"}
	en1 := typed.SlicestringToEnumerable(names1)
	names2 := []string{"Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	en2 := typed.SlicestringToEnumerable(names2)
	names3 := []string{"Solanki, Ajay", "Hoeing, Helge", "Andersen, Henriette Thaulow", "Potra, Cristina", "Iallo, Lucio"}
	en3 := typed.SlicestringToEnumerable(names3)
	namesList := typed.SliceEnumerableToEnumerable([]*enumerable.Enumerable{en1, en2, en3})

	allNames := namesList.AggregateSeedMust(enumerable.Empty(),
		func(current interface{}, next common.Elem) interface{} {
			enNext := next.(*enumerable.Enumerable)
			ss, _ := typed.EnumerableToSlicestring(enNext)
			if len(ss) > 3 {
				return current.(*enumerable.Enumerable).Union(enNext)
			}
			return current
		}).(*enumerable.Enumerable)

	allNames.Reset()
	for allNames.MoveNext() {
		fmt.Println(allNames.Current())
	}
}
