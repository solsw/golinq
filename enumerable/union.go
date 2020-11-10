package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 15 – Union
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-15-union/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

// Union produces the set union of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use UnionSelf instead.
func (en *Enumerable) Union(en2 *Enumerable) *Enumerable {
	return en.UnionEq(en2, reflect.DeepEqual)
}

// UnionSelf produces the set union of two Enumerables using reflect.DeepEqual as Equality.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) UnionSelf(en2 *Enumerable) *Enumerable {
	return en.UnionSelfEq(en2, reflect.DeepEqual)
}

// UnionEq produces the set union of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use UnionSelfEq instead.
func (en *Enumerable) UnionEq(en2 *Enumerable, eq common.Equality) *Enumerable {
	return en.Concat(en2).DistinctEq(eq)
}

// UnionSelfEq produces the set union of two Enumerables using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) UnionSelfEq(en2 *Enumerable, eq common.Equality) *Enumerable {
	sl2 := en2.Slice()
	en.Reset()
	return en.UnionEq(NewElems(sl2...), eq)
}

// UnionCmp produces the set union of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use UnionSelfCmp instead.
func (en *Enumerable) UnionCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	return distinctCmpPrim(en.Concat(en2), cmp), nil
}

// UnionCmpMust is like UnionCmp but panics in case of error.
func (en *Enumerable) UnionCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.UnionCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionSelfCmp produces the set union of two Enumerables using Comparison.
// (See help for DistinctCmp method.)
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) UnionSelfCmp(en2 *Enumerable, cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	sl2 := en2.Slice()
	en.Reset()
	return distinctCmpPrim(en.Concat(NewElems(sl2...)), cmp), nil
}

// UnionSelfCmpMust is like UnionSelfCmp but panics in case of error.
func (en *Enumerable) UnionSelfCmpMust(en2 *Enumerable, cmp common.Comparison) *Enumerable {
	r, err := en.UnionSelfCmp(en2, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionLs produces the set union of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use UnionSelfLs instead.
func (en *Enumerable) UnionLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	return distinctCmpPrim(en.Concat(en2), common.LessToComparison(ls)), nil
}

// UnionLsMust is like UnionLs but panics in case of error.
func (en *Enumerable) UnionLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.UnionLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionSelfLs produces the set union of two Enumerables using Less.
// (See help for DistinctCmp method.)
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) UnionSelfLs(en2 *Enumerable, ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	sl2 := en2.Slice()
	en.Reset()
	return distinctCmpPrim(en.Concat(NewElems(sl2...)), common.LessToComparison(ls)), nil
}

// UnionSelfLsMust is like UnionSelfLs but panics in case of error.
func (en *Enumerable) UnionSelfLsMust(en2 *Enumerable, ls common.Less) *Enumerable {
	r, err := en.UnionSelfLs(en2, ls)
	if err != nil {
		panic(err)
	}
	return r
}
