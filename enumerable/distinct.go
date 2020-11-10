package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 14 – Distinct
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-14-distinct/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

// Distinct returns distinct elements from Enumerable using reflect.DeepEqual as Equality.
func (en *Enumerable) Distinct() *Enumerable {
	return en.DistinctEq(reflect.DeepEqual)
}

// DistinctEq returns distinct elements from Enumerable using Equality.
// If 'eq' is nil reflect.DeepEqual is used.
func (en *Enumerable) DistinctEq(eq common.Equality) *Enumerable {
	if eq == nil {
		eq = reflect.DeepEqual
	}
	var c common.Elem
	seen := make([]common.Elem, 0)
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for en.MoveNext() {
				c = en.Current()
				if !elInElelEq(c, seen, eq) {
					seen = append(seen, c)
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { seen = make([]common.Elem, 0); en.Reset() },
	}}
}

func distinctCmpPrim(en *Enumerable, cmp common.Comparison) *Enumerable {
	var c common.Elem
	seen := make([]common.Elem, 0)
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			for en.MoveNext() {
				c = en.Current()
				i := elIdxInElelCmp(c, seen, cmp)
				if i == len(seen) || cmp(c, seen[i]) != 0 {
					elIntoElelAtIdx(c, &seen, i)
					return true
				}
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { seen = make([]common.Elem, 0); en.Reset() },
	}}
}

// DistinctCmp returns distinct elements from Enumerable using Comparison.
//
// Sorted slice of already seen elements is internally built.
// Sorted slice allows to use fast binary search to determine whether the element was seen or not.
// This may give performance gain when processing large Enumerables (though this is subject for benchmarking).
func (en *Enumerable) DistinctCmp(cmp common.Comparison) (*Enumerable, error) {
	if cmp == nil {
		return nil, errors.NilCmp
	}
	return distinctCmpPrim(en, cmp), nil
}

// DistinctCmpMust is like DistinctCmp but panics in case of error.
func (en *Enumerable) DistinctCmpMust(cmp common.Comparison) *Enumerable {
	r, err := en.DistinctCmp(cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctLs returns distinct elements from Enumerable using Less.
// (See help for DistinctCmp method.)
func (en *Enumerable) DistinctLs(ls common.Less) (*Enumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	return en.DistinctCmp(common.LessToComparison(ls))
}

// DistinctLsMust is like DistinctLs but panics in case of error.
func (en *Enumerable) DistinctLsMust(ls common.Less) *Enumerable {
	r, err := en.DistinctLs(ls)
	if err != nil {
		panic(err)
	}
	return r
}
