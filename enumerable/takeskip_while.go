package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 23 – Take/Skip/TakeWhile/SkipWhile
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-part-23-take-skip-takewhile-skipwhile/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.take
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

// Take returns a specified number of contiguous elements from the start of a sequence.
func (en *Enumerable) Take(count int) *Enumerable {
	i := 0
	var c common.Elem
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			if en.MoveNext() && i < count {
				c = en.Current()
				i++
				return true
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { i = 0; en.Reset() },
	}}
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
func (en *Enumerable) Skip(count int) *Enumerable {
	i := 1
	var c common.Elem
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for en.MoveNext() {
				c = en.Current()
				if i > count {
					return true
				}
				i++
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { i = 1; en.Reset() },
	}}
}

// TakeWhile returns elements from a sequence as long as a specified condition is true.
func (en *Enumerable) TakeWhile(pred common.Predicate) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				if en.MoveNext() {
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

// TakeWhileMust is like TakeWhile but panics in case of error.
func (en *Enumerable) TakeWhileMust(pred common.Predicate) *Enumerable {
	r, err := en.TakeWhile(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// TakeWhileIdx returns elements from a sequence as long as a specified condition is true.
// The element's index is used in the logic of the predicate function.
func (en *Enumerable) TakeWhileIdx(pred common.PredicateIdx) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	i := -1 // position before the first element
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				if en.MoveNext() {
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

// TakeWhileIdxMust is like TakeWhileIdx but panics in case of error.
func (en *Enumerable) TakeWhileIdxMust(pred common.PredicateIdx) *Enumerable {
	r, err := en.TakeWhileIdx(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
func (en *Enumerable) SkipWhile(pred common.Predicate) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	allRemaining := false
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for en.MoveNext() {
					c = en.Current()
					if allRemaining || !pred(c) {
						allRemaining = true
						return true
					}
				}
				return false
			},
			Crrnt: func() common.Elem { return c },
			Rst:   func() { allRemaining = false; en.Reset() },
		}},
		nil
}

// SkipWhileMust is like SkipWhile but panics in case of error.
func (en *Enumerable) SkipWhileMust(pred common.Predicate) *Enumerable {
	r, err := en.SkipWhile(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipWhileIdx bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// The element's index is used in the logic of the predicate function.
func (en *Enumerable) SkipWhileIdx(pred common.PredicateIdx) (*Enumerable, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	var c common.Elem
	i := -1 // position before the first element
	allRemaining := false
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for en.MoveNext() {
					c = en.Current()
					i++
					if allRemaining || !pred(c, i) {
						allRemaining = true
						return true
					}
				}
				return false
			},
			Crrnt: func() common.Elem { return c },
			Rst:   func() { i = -1; allRemaining = false; en.Reset() },
		}},
		nil
}

// SkipWhileIdxMust is like SkipWhileIdx but panics in case of error.
func (en *Enumerable) SkipWhileIdxMust(pred common.PredicateIdx) *Enumerable {
	r, err := en.SkipWhileIdx(pred)
	if err != nil {
		panic(err)
	}
	return r
}
