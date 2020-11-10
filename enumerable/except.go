package enumerable

import (
	"reflect"
	"sort"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 17 – Except
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-17-except/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

// Except produces the set difference of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ExceptSelf instead.
func (en *Enumerable) Except(en2 *Enumerable) *Enumerable {
	return en.ExceptEq(en2, reflect.DeepEqual)
}

// ExceptSelf produces the set difference of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) ExceptSelf(en2 *Enumerable) *Enumerable {
	return en.ExceptSelfEq(en2, reflect.DeepEqual)
}

// ExceptEq produces the set difference of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ExceptSelfEq instead.
func (en *Enumerable) ExceptEq(en2 *Enumerable, eq common.Equality) *Enumerable {
	if eq == nil {
		eq = reflect.DeepEqual
	}
	d := en.DistinctEq(eq)
	var c common.Elem
	ds2 := en2.DistinctEq(eq).Slice()
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for d.MoveNext() {
				c = d.Current()
				if !elInElelEq(c, ds2, eq) {
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { d.Reset() },
	}}
}

// ExceptSelfEq produces the set difference of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) ExceptSelfEq(en2 *Enumerable, eq common.Equality) *Enumerable {
	sl2 := en2.Slice()
	en.Reset()
	return en.ExceptEq(NewElems(sl2...), eq)
}

func exceptCmpPrim(en1, en2 *Enumerable, cmp common.Comparison) *Enumerable {
	d := en1.DistinctCmpMust(cmp)
	var c common.Elem
	ds2 := en2.DistinctCmpMust(cmp).Slice()
	sort.Slice(ds2, func(i, j int) bool { return cmp(ds2[i], ds2[j]) < 0 })
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for d.MoveNext() {
				c = d.Current()
				if !elInElelCmp(c, ds2, cmp) {
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { d.Reset() },
	}}
}

// ExceptCmp produces the set difference of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ExceptSelfCmp instead.
func (en *Enumerable) ExceptCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	return exceptCmpPrim(en, en2, cmp), nil
}

// ExceptCmpMust is like ExceptCmp but panics in case of error.
func (en *Enumerable) ExceptCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.ExceptCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptSelfCmp produces the set difference of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) ExceptSelfCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	sl2 := en2.Slice()
	en.Reset()
	return exceptCmpPrim(en, NewElems(sl2...), cmp), nil
}

// ExceptSelfCmpMust is like ExceptSelfCmp but panics in case of error.
func (en *Enumerable) ExceptSelfCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.ExceptSelfCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptLs produces the set difference of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ExceptSelfLs instead.
func (en *Enumerable) ExceptLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	return exceptCmpPrim(en, en2, common.LessToComparison(ls)), nil
}

// ExceptLsMust is like ExceptLs but panics in case of error.
func (en *Enumerable) ExceptLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.ExceptLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptSelfLs produces the set difference of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) ExceptSelfLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	sl2 := en2.Slice()
	en.Reset()
	return exceptCmpPrim(en, NewElems(sl2...), common.LessToComparison(ls)), nil
}

// ExceptSelfLsMust is like ExceptSelfLs but panics in case of error.
func (en *Enumerable) ExceptSelfLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.ExceptSelfLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}
