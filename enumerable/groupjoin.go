package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 22 – GroupJoin
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-22-groupjoin/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin

// see example/groupjoin

func groupJoinEqPrim(outer, inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) *Enumerable {
	if keq == nil {
		keq = reflect.DeepEqual
	}
	ilk := toLookupPrim(inner, iksel, nil, keq, false)
	return selectPrim(outer,
		func(el common.Elem) common.Elem { return rsel(el, ilk.Item(oksel(el))) })
}

// GroupJoinEq correlates the elements of two Enums based on key equality and groups the results.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
// 'outer' and 'inner' must NOT be based on the same Enumerable, otherwise use GroupJoinSelfEq instead.
func (outer *Enumerable) GroupJoinEq(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) (*Enumerable, error) {
	if oksel == nil || iksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	return groupJoinEqPrim(outer, inner, oksel, iksel, rsel, keq), nil
}

// GroupJoinEqMust is like GroupJoinEq but panics in case of error.
func (outer *Enumerable) GroupJoinEqMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) *Enumerable {
	r, err := outer.GroupJoinEq(inner, oksel, iksel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoinSelfEq correlates the elements of two Enums based on key equality and groups the results.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
// 'outer' and 'inner' may be based on the same Enumerable.
func (outer *Enumerable) GroupJoinSelfEq(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) (*Enumerable, error) {
	if oksel == nil || iksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	isl := inner.Slice()
	outer.Reset()
	return groupJoinEqPrim(outer, NewElems(isl...), oksel, iksel, rsel, keq), nil
}

// GroupJoinSelfEqMust is like GroupJoinSelfEq but panics in case of error.
func (outer *Enumerable) GroupJoinSelfEqMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) *Enumerable {
	r, err := outer.GroupJoinSelfEq(inner, oksel, iksel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoin correlates the elements of two Enums based on equality of keys and groups the results.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector.
// reflect.DeepEqual is used as keys Equality.
// 'outer' and 'inner' must NOT be based on the same Enumerable, otherwise use GroupJoinSelf instead.
func (outer *Enumerable) GroupJoin(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) (*Enumerable, error) {
	return outer.GroupJoinEq(inner, oksel, iksel, rsel, reflect.DeepEqual)
}

// GroupJoinMust is like GroupJoin but panics in case of error.
func (outer *Enumerable) GroupJoinMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) *Enumerable {
	r, err := outer.GroupJoin(inner, oksel, iksel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoinSelf correlates the elements of two Enums based on equality of keys and groups the results.
// 'oksel', 'iksel' - outer and inner key selectors. 'rsel' - result selector.
// reflect.DeepEqual is used as keys Equality.
// 'outer' and 'inner' may be based on the same Enumerable.
func (outer *Enumerable) GroupJoinSelf(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) (*Enumerable, error) {
	return outer.GroupJoinSelfEq(inner, oksel, iksel, rsel, reflect.DeepEqual)
}

// GroupJoinSelfMust is like GroupJoinSelf but panics in case of error.
func (outer *Enumerable) GroupJoinSelfMust(inner *Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) *Enumerable {
	r, err := outer.GroupJoinSelf(inner, oksel, iksel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}
