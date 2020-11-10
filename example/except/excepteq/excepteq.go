package main

import (
	"fmt"
	"strings"

	"github.com/solsw/golinq/common"
)

// see last two examples from Enumerable.Except help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

type product struct {
	name string
	code int
}

func main() {
	fruits1 := SliceproductToEnumerable([]product{
		product{name: "apple", code: 9},
		product{name: "orange", code: 4},
		product{name: "lemon", code: 12},
	})
	fruits2 := SliceproductToEnumerable([]product{
		product{name: "APPLE", code: 9},
	})
	//Get all the elements from the first array except for the elements from the second array.
	except := fruits1.ExceptEq(fruits2,
		func(e1, e2 common.Elem) bool {
			p1 := e1.(product)
			p2 := e2.(product)
			return p1.code == p2.code && strings.ToUpper(p1.name) == strings.ToUpper(p2.name)
		})
	for except.MoveNext() {
		p := except.Current().(product)
		fmt.Printf("%s %d\n", p.name, p.code)
	}
}
