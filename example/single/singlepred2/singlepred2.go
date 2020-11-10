package main

import (
	"fmt"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/errors"
)

// see example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits := enumerable.NewElems("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit2, err := fruits.SinglePred(func(e common.Elem) bool { return len(e.(string)) > 5 })
	if err == errors.MultiMatch {
		fmt.Println("The collection does not contain exactly one element whose length is greater than 5.")
	} else {
		fmt.Println(fruit2)
	}
}
