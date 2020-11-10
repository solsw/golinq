package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 19 – Join
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-19-join/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join

func joinEqPrim(outer, inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) *Enumerable {
	if keq == nil {
		keq = reflect.DeepEqual
	}
	ilk := toLookupPrim(inner, iksel, nil, keq, false)
	t := Empty()
	var oel, iel common.Elem
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for {
				if t.MoveNext() {
					iel = t.Current()
					return true
				}
				if !outer.MoveNext() {
					return false
				}
				oel = outer.Current()
				t = ilk.Item(oksel(oel))
			}
		},
		Crrnt: func() common.Elem { return rsel(oel, iel) },
		Rst:   func() { t = Empty(); outer.Reset() },
	}}
}

// JoinEq correlates the elements of two Enums based on matching keys.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
// 'outer' and 'inner' must NOT be based on the same Enumerable, otherwise use JoinSelfEq instead.
//
// (The similar to keys equality comparison functionality may be achieved using appropriate key selectors.
// See CustomComparer test for usage of case insensitive string keys.)
func (outer *Enumerable) JoinEq(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) (*Enumerable, error) {
	if oksel == nil || iksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	return joinEqPrim(outer, inner, oksel, iksel, rsel, keq), nil
}

// JoinEqMust is like JoinEq but panics in case of error.
func (outer *Enumerable) JoinEqMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) *Enumerable {
	r, err := outer.JoinEq(inner, oksel, iksel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// JoinSelfEq correlates the elements of two Enums based on matching keys.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
// 'outer' and 'inner' may be based on the same Enumerable.
func (outer *Enumerable) JoinSelfEq(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) (*Enumerable, error) {
	if oksel == nil || iksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	isl := inner.Slice()
	outer.Reset()
	return joinEqPrim(outer, NewElems(isl...), oksel, iksel, rsel, keq), nil
}

// JoinSelfEqMust is like JoinSelfEq but panics in case of error.
func (outer *Enumerable) JoinSelfEqMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) *Enumerable {
	r, err := outer.JoinSelfEq(inner, oksel, iksel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// Join correlates the elements of two Enums based on matching keys.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector.
// reflect.DeepEqual is used as keys Equality.
// 'outer' and 'inner' must NOT be based on the same Enumerable, otherwise use JoinSelf instead.
func (outer *Enumerable) Join(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	return outer.JoinEq(inner, oksel, iksel, rsel, reflect.DeepEqual)
}

// JoinMust is like Join but panics in case of error.
func (outer *Enumerable) JoinMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := outer.Join(inner, oksel, iksel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}

// JoinSelf correlates the elements of two Enums based on matching keys.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector.
// reflect.DeepEqual is used as keys Equality.
// 'outer' and 'inner' may be based on the same Enumerable.
func (outer *Enumerable) JoinSelf(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	return outer.JoinSelfEq(inner, oksel, iksel, rsel, reflect.DeepEqual)
}

// JoinSelfMust is like JoinSelf but panics in case of error.
func (outer *Enumerable) JoinSelfMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := outer.JoinSelf(inner, oksel, iksel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}
