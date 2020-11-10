package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

// returned count is used by AverageInt
func (en *Enumerable) sumIntCount(sel func(common.Elem) int) (int, int) {
	r := 0
	count := 0
	for en.MoveNext() {
		r += sel(en.Current())
		count++
	}
	return r, count
}

// SumInt computes sum of the sequence of int values
// that are obtained by invoking a transform function on each element of Enumerable.
func (en *Enumerable) SumInt(sel func(common.Elem) int) (int, error) {
	if sel == nil {
		return 0, errors.NilSel
	}
	r, _ := en.sumIntCount(sel)
	return r, nil
}

// SumIntMust is like SumInt but panics in case of error.
func (en *Enumerable) SumIntMust(sel func(common.Elem) int) int {
	r, err := en.SumInt(sel)
	if err != nil {
		panic(err)
	}
	return r
}

// returned count is used by AverageFloat64
func (en *Enumerable) sumFloat64Count(sel func(common.Elem) float64) (float64, int) {
	r := 0.0
	count := 0
	for en.MoveNext() {
		r += sel(en.Current())
		count++
	}
	return r, count
}

// SumFloat64 computes the sum of the sequence of float64 values
// that are obtained by invoking a transform function on each element of Enumerable.
func (en *Enumerable) SumFloat64(sel func(common.Elem) float64) (float64, error) {
	if sel == nil {
		return 0, errors.NilSel
	}
	r, _ := en.sumFloat64Count(sel)
	return r, nil
}

// SumFloat64Must is like SumFloat64 but panics in case of error.
func (en *Enumerable) SumFloat64Must(sel func(common.Elem) float64) float64 {
	r, err := en.SumFloat64(sel)
	if err != nil {
		panic(err)
	}
	return r
}
