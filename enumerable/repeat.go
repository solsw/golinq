package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 6 – Repeat
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-6-repeat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat

// Repeat generates Enumerable that contains one repeated value 'el'.
func Repeat(el common.Elem, count int) (*Enumerable, error) {
	if count < 0 {
		return nil, errors.NegCount
	}
	i := 0
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				if i >= count {
					return false
				}
				i++
				return true
			},
			Crrnt: func() common.Elem { return el },
			Rst:   func() { i = 0 },
		}},
		nil
}

// RepeatMust is like Repeat but panics in case of error.
func RepeatMust(el common.Elem, count int) *Enumerable {
	r, err := Repeat(el, count)
	if err != nil {
		panic(err)
	}
	return r
}
