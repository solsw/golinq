package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

// AverageInt computes the average of int values
// that are obtained by invoking a transform function on each element of Enumerable.
func (en *Enumerable) AverageInt(sel func(common.Elem) int) (float64, error) {
	if sel == nil {
		return 0, errors.NilSel
	}
	s, count := en.sumIntCount(sel)
	if count == 0 {
		return 0, errors.EmptyEnum
	}
	return float64(s) / float64(count), nil
}

// AverageIntMust is like AverageInt but panics in case of error.
func (en *Enumerable) AverageIntMust(sel func(common.Elem) int) float64 {
	r, err := en.AverageInt(sel)
	if err != nil {
		panic(err)
	}
	return r
}

// AverageFloat64 computes the average of float64 values
// that are obtained by invoking a transform function on each element of Enumerable.
func (en *Enumerable) AverageFloat64(sel func(common.Elem) float64) (float64, error) {
	if sel == nil {
		return 0.0, errors.NilSel
	}
	s, count := en.sumFloat64Count(sel)
	if count == 0 {
		return 0.0, errors.EmptyEnum
	}
	return s / float64(count), nil
}

// AverageFloat64Must is like AverageFloat64 but panics in case of error.
func (en *Enumerable) AverageFloat64Must(sel func(common.Elem) float64) float64 {
	r, err := en.AverageFloat64(sel)
	if err != nil {
		panic(err)
	}
	return r
}
