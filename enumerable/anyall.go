package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 10 – Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all

// Any determines whether Enumerable contains any elements.
// 'en' may need Reset() for further use.
func (en *Enumerable) Any() bool {
	if c, ok := (en.enmr).(common.Counter); ok {
		return c.Count() > 0
	}
	return en.MoveNext()
}

// AnyPred determines whether any element of Enumerable satisfies a condition.
// 'en' may need Reset() for further use.
func (en *Enumerable) AnyPred(pred common.Predicate) (bool, error) {
	if pred == nil {
		return false, errors.NilPred
	}
	for en.MoveNext() {
		if pred(en.Current()) {
			return true, nil
		}
	}
	return false, nil
}

// AnyPredMust is like AnyPred but panics in case of error.
// 'en' may need Reset() for further use.
func (en *Enumerable) AnyPredMust(pred common.Predicate) bool {
	r, err := en.AnyPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// All determines whether all elements of Enumerable satisfy a condition.
// 'en' may need Reset() for further use.
func (en *Enumerable) All(pred common.Predicate) (bool, error) {
	if pred == nil {
		return false, errors.NilPred
	}
	for en.MoveNext() {
		if !pred(en.Current()) {
			return false, nil
		}
	}
	return true, nil
}

// AllMust is like All but panics in case of error.
// 'en' may need Reset() for further use.
func (en *Enumerable) AllMust(pred common.Predicate) bool {
	r, err := en.All(pred)
	if err != nil {
		panic(err)
	}
	return r
}
