package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 7 – Count and LongCount
// https://codeblog.jonskeet.uk/2010/12/26/reimplementing-linq-to-objects-part-7-count-and-longcount/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

func countPrim(en *Enumerable, pred common.Predicate) int {
	r := 0
	for en.MoveNext() {
		if pred == nil || pred(en.Current()) {
			r++
		}
	}
	return r
}

// Count returns the number of elements in the Enumerable.
func (en *Enumerable) Count() int {
	if c, ok := (en.enmr).(common.Counter); ok {
		return c.Count()
	}
	return countPrim(en, nil)
}

// CountPred returns a number that represents how many elements in the Enumerable satisfy a condition.
func (en *Enumerable) CountPred(pred common.Predicate) (int, error) {
	if pred == nil {
		return 0, errors.NilPred
	}
	return countPrim(en, pred), nil
}

// CountPredMust is like CountPred but panics in case of error.
func (en *Enumerable) CountPredMust(pred common.Predicate) int {
	r, err := en.CountPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}
