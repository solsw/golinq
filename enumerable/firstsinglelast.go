package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 11 – First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

// First returns the first element of Enumerable.
func (en *Enumerable) First() (common.Elem, error) {
	if c, cok := (en.enmr).(common.Counter); cok {
		if c.Count() == 0 {
			return nil, errors.EmptyEnum
		}
		// if en.enmr is Itemer but not Counter, it's not safe to call Item(0) because en.enmr may be empty
		if i, iok := (en.enmr).(common.Itemer); iok {
			return i.Item(0), nil
		}
	}
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	return en.Current(), nil
}

// FirstMust is like First but panics in case of error.
func (en *Enumerable) FirstMust() common.Elem {
	r, err := en.First()
	if err != nil {
		panic(err)
	}
	return r
}

// FirstPred returns the first element in Enumerable that satisfies a specified condition.
func (en *Enumerable) FirstPred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	r := en.Current()
	if pred(r) {
		return r, nil
	}
	for en.MoveNext() {
		r = en.Current()
		if pred(r) {
			return r, nil
		}
	}
	return nil, errors.NoMatch
}

// FirstPredMust is like FirstPred but panics in case of error.
func (en *Enumerable) FirstPredMust(pred common.Predicate) common.Elem {
	r, err := en.FirstPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstOrDefault returns the first element of Enumerable or nil if Enumerable contains no elements.
func (en *Enumerable) FirstOrDefault() common.Elem {
	if !en.MoveNext() {
		return nil
	}
	return en.Current()
}

// FirstOrDefaultPred returns the first element of Enumerable that satisfies a condition or nil if no such element is found.
func (en *Enumerable) FirstOrDefaultPred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	for en.MoveNext() {
		r := en.Current()
		if pred(r) {
			return r, nil
		}
	}
	return nil, nil
}

// FirstOrDefaultPredMust is like FirstOrDefaultPred but panics in case of error.
func (en *Enumerable) FirstOrDefaultPredMust(pred common.Predicate) common.Elem {
	r, err := en.FirstOrDefaultPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// Single returns the only element of Enumerable and returns an error if there is not exactly one element in Enumerable.
func (en *Enumerable) Single() (common.Elem, error) {
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	if en.MoveNext() {
		return nil, errors.MultiElems
	}
	return en.Current(), nil
}

// SingleMust is like Single but panics in case of error.
func (en *Enumerable) SingleMust() common.Elem {
	r, err := en.Single()
	if err != nil {
		panic(err)
	}
	return r
}

// SinglePred returns the only element of Enumerable that satisfies a specified condition,
// and returns an error if there is no such element or more than one such element exists.
func (en *Enumerable) SinglePred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	empty := true
	found := false
	var r common.Elem
	for en.MoveNext() {
		empty = false
		c := en.Current()
		if pred(c) {
			if found {
				return nil, errors.MultiMatch
			}
			found = true
			r = c
		}
	}
	if empty {
		return nil, errors.EmptyEnum
	}
	if !found {
		return nil, errors.NoMatch
	}
	return r, nil
}

// SinglePredMust is like SinglePred but panics in case of error.
func (en *Enumerable) SinglePredMust(pred common.Predicate) common.Elem {
	r, err := en.SinglePred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefault returns the only element of Enumerable or a default value if Enumerable is empty.
// This method returns an error if there is more than one element in Enumerable.
func (en *Enumerable) SingleOrDefault() (common.Elem, error) {
	if !en.MoveNext() {
		return nil, nil
	}
	if en.MoveNext() {
		return nil, errors.MultiElems
	}
	return en.Current(), nil
}

// SingleOrDefaultMust is like SingleOrDefault but panics in case of error.
func (en *Enumerable) SingleOrDefaultMust() common.Elem {
	r, err := en.SingleOrDefault()
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefaultPred returns the only element of Enumerable
// that satisfies a specified condition or a default value if no such element exists.
// This method returns an error if more than one element satisfies the condition.
func (en *Enumerable) SingleOrDefaultPred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	found := false
	var r common.Elem
	for en.MoveNext() {
		c := en.Current()
		if pred(c) {
			if found {
				return nil, errors.MultiMatch
			}
			found = true
			r = c
		}
	}
	if !found {
		return nil, nil
	}
	return r, nil
}

// SingleOrDefaultPredMust is like SingleOrDefaultPred but panics in case of error.
func (en *Enumerable) SingleOrDefaultPredMust(pred common.Predicate) common.Elem {
	r, err := en.SingleOrDefaultPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// Last returns the last element of Enumerable.
func (en *Enumerable) Last() (common.Elem, error) {
	if c, cok := (en.enmr).(common.Counter); cok {
		if c.Count() == 0 {
			return nil, errors.EmptyEnum
		}
		if i, iok := (en.enmr).(common.Itemer); iok {
			return i.Item(c.Count() - 1), nil
		}
	}
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	r := en.Current()
	for en.MoveNext() {
		r = en.Current()
	}
	return r, nil
}

// LastMust is like Last but panics in case of error.
func (en *Enumerable) LastMust() common.Elem {
	r, err := en.Last()
	if err != nil {
		panic(err)
	}
	return r
}

// LastPred returns the last element of Enumerable that satisfies a specified condition.
func (en *Enumerable) LastPred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	found := false
	var r common.Elem
	c := en.Current()
	if pred(c) {
		found = true
		r = c
	}
	for en.MoveNext() {
		c = en.Current()
		if pred(c) {
			found = true
			r = c
		}
	}
	if !found {
		return nil, errors.NoMatch
	}
	return r, nil
}

// LastPredMust is like LastPred but panics in case of error.
func (en *Enumerable) LastPredMust(pred common.Predicate) common.Elem {
	r, err := en.LastPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefault returns the last element of Enumerable or a default value if Enumerable contains no elements.
func (en *Enumerable) LastOrDefault() (common.Elem, error) {
	if c, cok := (en.enmr).(common.Counter); cok {
		if c.Count() == 0 {
			return nil, nil
		}
		if i, iok := (en.enmr).(common.Itemer); iok {
			return i.Item(c.Count() - 1), nil
		}
	}
	if !en.MoveNext() {
		return nil, nil
	}
	r := en.Current()
	for en.MoveNext() {
		r = en.Current()
	}
	return r, nil
}

// LastOrDefaultMust is like LastOrDefault but panics in case of error.
func (en *Enumerable) LastOrDefaultMust() common.Elem {
	r, err := en.LastOrDefault()
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefaultPred returns the last element of Enumerable
// that satisfies a condition or a default value if no such element is found.
func (en *Enumerable) LastOrDefaultPred(pred common.Predicate) (common.Elem, error) {
	if pred == nil {
		return nil, errors.NilPred
	}
	empty := true
	var r common.Elem = nil
	for en.MoveNext() {
		empty = false
		c := en.Current()
		if pred(c) {
			r = c
		}
	}
	if empty {
		return nil, nil
	}
	return r, nil
}

// LastOrDefaultPredMust is like LastOrDefaultPred but panics in case of error.
func (en *Enumerable) LastOrDefaultPredMust(pred common.Predicate) common.Elem {
	r, err := en.LastOrDefaultPred(pred)
	if err != nil {
		panic(err)
	}
	return r
}
