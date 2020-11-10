package enumerable

import (
	"math"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 4 – Range
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-4-range/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range

// Range generates a sequence of int numbers within a specified range.
func Range(start, count int) (*Enumerable, error) {
	if count < 0 {
		return nil, errors.NegCount
	}
	if int64(start)+int64(count)-int64(1) > math.MaxInt32 {
		return nil, errors.WrongStrtCnt
	}
	i := 0
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				if i < count {
					i++
					return true
				}
				return false
			},
			Crrnt: func() common.Elem { return start + i - 1 },
			Rst:   func() { i = 0 },
		}},
		nil
}

// RangeMust is like Range but panics in case of error.
func RangeMust(start, count int) *Enumerable {
	r, err := Range(start, count)
	if err != nil {
		panic(err)
	}
	return r
}
