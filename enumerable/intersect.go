package enumerable

import (
	"reflect"
	"sort"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 16 – Intersect (and build fiddling)
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-16-intersect-and-build-fiddling/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

// Intersect produces the set intersection of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use IntersectSelf instead.
func (en *Enumerable) Intersect(en2 *Enumerable) *Enumerable {
	return en.IntersectEq(en2, reflect.DeepEqual)
}

// IntersectSelf produces the set intersection of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) IntersectSelf(en2 *Enumerable) *Enumerable {
	return en.IntersectSelfEq(en2, reflect.DeepEqual)
}

// IntersectEq produces the set intersection of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use IntersectSelfEq instead.
func (en *Enumerable) IntersectEq(en2 *Enumerable, eq common.Equality) *Enumerable {
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
				if elInElelEq(c, ds2, eq) {
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { en.Reset(); d = en.DistinctEq(eq) },
	}}
}

// IntersectSelfEq produces the set intersection of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) IntersectSelfEq(en2 *Enumerable, eq common.Equality) *Enumerable {
	sl2 := en2.Slice()
	en.Reset()
	return en.IntersectEq(NewElems(sl2...), eq)
}

func intersectCmpPrim(en, en2 *Enumerable, cmp common.Comparison) *Enumerable {
	d := en.DistinctCmpMust(cmp)
	var c common.Elem
	ds2 := en2.DistinctCmpMust(cmp).Slice()
	sort.Slice(ds2, func(i, j int) bool { return cmp(ds2[i], ds2[j]) < 0 })
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for d.MoveNext() {
				c = d.Current()
				if elInElelCmp(c, ds2, cmp) {
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { en.Reset(); d = en.DistinctCmpMust(cmp) },
	}}
}

// IntersectCmp produces the set intersection of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use IntersectSelfCmp instead.
func (en *Enumerable) IntersectCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	return intersectCmpPrim(en, en2, cmp), nil
}

// IntersectCmpMust is like IntersectCmp but panics in case of error.
func (en *Enumerable) IntersectCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.IntersectCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectSelfCmp produces the set intersection of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) IntersectSelfCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	sl2 := en2.Slice()
	en.Reset()
	return intersectCmpPrim(en, NewElems(sl2...), cmp), nil
}

// IntersectSelfCmpMust is like IntersectSelfCmp but panics in case of error.
func (en *Enumerable) IntersectSelfCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.IntersectSelfCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectLs produces the set intersection of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use IntersectSelfLs instead.
func (en *Enumerable) IntersectLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	return intersectCmpPrim(en, en2, common.LessToComparison(ls)), nil
}

// IntersectLsMust is like IntersectLs but panics in case of error.
func (en *Enumerable) IntersectLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.IntersectLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectSelfLs produces the set intersection of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en2' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'en'.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) IntersectSelfLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	sl2 := en2.Slice()
	en.Reset()
	return intersectCmpPrim(en, NewElems(sl2...), common.LessToComparison(ls)), nil
}

// IntersectSelfLsMust is like IntersectSelfLs but panics in case of error.
func (en *Enumerable) IntersectSelfLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.IntersectSelfLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}
