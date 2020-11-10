package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 2 - "Where"
// https://codeblog.jonskeet.uk/2010/09/03/reimplementing-linq-to-objects-part-2-quot-where-quot/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

// see example/selectmany/selectmanyex3

// Where filters Enumerable's elements based on a predicate.
func (en *Enumerable) Where(pred common.Predicate) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for en.MoveNext() {
					c = en.Current()
					if pred(c) {
						return true
					}
				}
				return false
			},
			Crrnt: func() common.Elem { return c },
			Rst:   en.Reset,
		}},
		nil
}

// WhereMust is like Where but panics in case of error.
func (en *Enumerable) WhereMust(pred common.Predicate) *Enumerable {
	r, err := en.Where(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// WhereIdx filters Enumerable's elements based on a predicate.
// Each element's index is used in the logic of the predicate function.
func (en *Enumerable) WhereIdx(pred common.PredicateIdx) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	i := -1 // position before the first element
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for en.MoveNext() {
					c = en.Current()
					i++
					if pred(c, i) {
						return true
					}
				}
				return false
			},
			Crrnt: func() common.Elem { return c },
			Rst:   func() { i = -1; en.Reset() },
		}},
		nil
}

// WhereIdxMust is like WhereIdx but panics in case of error.
func (en *Enumerable) WhereIdxMust(pred common.PredicateIdx) *Enumerable {
	r, err := en.WhereIdx(pred)
	if err != nil {
		panic(err)
	}
	return r
}
