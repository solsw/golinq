package chain

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

type (
	// Link projects one enumerable.Enumerable into another.
	Link = func(*enumerable.Enumerable) (*enumerable.Enumerable, error)
)

// Chain sequentially applies 'links' to 'en' and returns the resulting enumerable.Enumerable.
// If any of the 'links' returns an error, the error is immediately returned.
func Chain(en *enumerable.Enumerable, links ...Link) (*enumerable.Enumerable, error) {
	if en == nil || len(links) == 0 {
		return en, nil
	}
	r, err := links[0](en)
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(links); i++ {
		r, err = links[i](r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// WhereLink produces WhereEr-based Link.
func WhereLink(pred common.Predicate) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Where(pred) }
}

// WhereIdxLink produces WhereIdxEr-based Link.
func WhereIdxLink(pred common.PredicateIdx) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.WhereIdx(pred) }
}

// SelectLink produces SelectEr-based Link.
func SelectLink(sel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Select(sel) }
}

// SelectIdxLink produces SelectIdxEr-based Link.
func SelectIdxLink(sel func(common.Elem, int) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.SelectIdx(sel) }
}

// RangeLink produces RangeEr-based Link.
func RangeLink(start, count int) Link {
	return func(*enumerable.Enumerable) (*enumerable.Enumerable, error) { return enumerable.Range(start, count) }
}

// EmptyLink produces Empty-based Link.
func EmptyLink() Link {
	return func(*enumerable.Enumerable) (*enumerable.Enumerable, error) { return enumerable.Empty(), nil }
}

// RepeatLink produces RepeatEr-based Link.
func RepeatLink(el common.Elem, count int) Link {
	return func(*enumerable.Enumerable) (*enumerable.Enumerable, error) { return enumerable.Repeat(el, count) }
}

// ConcatLink produces Concat-based Link.
func ConcatLink(en2 *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Concat(en2), nil }
}

// SelectManyLink produces SelectManyEr-based Link.
func SelectManyLink(sel func(common.Elem) *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.SelectMany(sel) }
}

// SelectManyIdxLink produces SelectManyIdxEr-based Link.
func SelectManyIdxLink(sel func(common.Elem, int) *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.SelectManyIdx(sel) }
}

// SelectManyCollLink produces SelectManyCollEr-based Link.
func SelectManyCollLink(sel1 func(common.Elem) *enumerable.Enumerable,
	sel2 func(common.Elem, common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.SelectManyColl(sel1, sel2)
	}
}

// SelectManyIdxCollLink produces SelectManyIdxCollEr-based Link.
func SelectManyIdxCollLink(sel1 func(common.Elem, int) *enumerable.Enumerable,
	sel2 func(common.Elem, common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.SelectManyIdxColl(sel1, sel2)
	}
}

// DefaultIfEmptyLink produces DefaultIfEmpty-based Link.
func DefaultIfEmptyLink() Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.DefaultIfEmpty(), nil }
}

// DefaultIfEmptyDefLink produces DefaultIfEmptyDef-based Link.
func DefaultIfEmptyDefLink(def common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.DefaultIfEmptyDef(def), nil }
}

// DistinctLink produces Distinct-based Link.
func DistinctLink() Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Distinct(), nil }
}

// DistinctEqLink produces DistinctEq-based Link.
func DistinctEqLink(eq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.DistinctEq(eq), nil }
}

// DistinctCmpLink produces DistinctCmpEr-based Link.
func DistinctCmpLink(cmp common.Comparison) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.DistinctCmp(cmp) }
}

// DistinctLsLink produces DistinctLsEr-based Link.
func DistinctLsLink(ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.DistinctLs(ls) }
}

// UnionLink produces Union-based Link.
func UnionLink(en2 *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Union(en2), nil }
}

// UnionEqLink produces UnionEq-based Link.
func UnionEqLink(en2 *enumerable.Enumerable, eq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.UnionEq(en2, eq), nil }
}

// UnionCmpLink produces UnionCmpEr-based Link.
func UnionCmpLink(en2 *enumerable.Enumerable, cmp common.Comparison) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.UnionCmp(en2, cmp) }
}

// UnionLsLink produces UnionLsEr-based Link.
func UnionLsLink(en2 *enumerable.Enumerable, ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.UnionLs(en2, ls) }
}

// IntersectLink produces Intersect-based Link.
func IntersectLink(en2 *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Intersect(en2), nil }
}

// IntersectEqLink produces IntersectEq-based Link.
func IntersectEqLink(en2 *enumerable.Enumerable, eq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.IntersectEq(en2, eq), nil }
}

// IntersectCmpLink produces IntersectCmpEr-based Link.
func IntersectCmpLink(en2 *enumerable.Enumerable, cmp common.Comparison) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.IntersectCmp(en2, cmp) }
}

// IntersectLsLink produces IntersectLsEr-based Link.
func IntersectLsLink(en2 *enumerable.Enumerable, ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.IntersectLs(en2, ls) }
}

// ExceptLink produces Except-based Link.
func ExceptLink(en2 *enumerable.Enumerable) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Except(en2), nil }
}

// ExceptEqLink produces ExceptEq-based Link.
func ExceptEqLink(en2 *enumerable.Enumerable, eq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.ExceptEq(en2, eq), nil }
}

// ExceptCmpLink produces ExceptCmpEr-based Link.
func ExceptCmpLink(en2 *enumerable.Enumerable, cmp common.Comparison) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.ExceptCmp(en2, cmp) }
}

// ExceptLsLink produces ExceptLsEr-based Link.
func ExceptLsLink(en2 *enumerable.Enumerable, ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.ExceptLs(en2, ls) }
}

// ToLookupSelEqLink produces ToLookupSelEqEr-based Link.
// The resulting Link returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupSelEqLink(ksel, esel func(common.Elem) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		l, err := en.ToLookupSelEq(ksel, esel, keq)
		if err != nil {
			return nil, err
		}
		return l.Enumerable(), nil
	}
}

// ToLookupSelLink produces ToLookupSelEr-based Link.
// The resulting Link returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupSelLink(ksel, esel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		l, err := en.ToLookupSel(ksel, esel)
		if err != nil {
			return nil, err
		}
		return l.Enumerable(), nil
	}
}

// ToLookupEqLink produces ToLookupEqEr-based Link.
// The resulting Link returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupEqLink(ksel func(common.Elem) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		l, err := en.ToLookupEq(ksel, keq)
		if err != nil {
			return nil, err
		}
		return l.Enumerable(), nil
	}
}

// ToLookupLink produces ToLookupEr-based Link.
// The resulting Link returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupLink(ksel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		l, err := en.ToLookup(ksel)
		if err != nil {
			return nil, err
		}
		return l.Enumerable(), nil
	}
}

// JoinEqLink produces JoinEqEr-based Link.
func JoinEqLink(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.JoinEq(inner, oksel, iksel, rsel, keq)
	}
}

// JoinLink produces JoinEr-based Link.
func JoinLink(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.Join(inner, oksel, iksel, rsel)
	}
}

// GroupBySelResLink produces GroupBySelResEr-based Link.
func GroupBySelResLink(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupBySelRes(ksel, esel, rsel)
	}
}

// GroupBySelResEqLink produces GroupBySelResEqEr-based Link.
func GroupBySelResEqLink(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupBySelResEq(ksel, esel, rsel, keq)
	}
}

// GroupBySelLink produces GroupBySelEr-based Link.
func GroupBySelLink(ksel, esel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.GroupBySel(ksel, esel) }
}

// GroupBySelEqLink produces GroupBySelEqEr-based Link.
func GroupBySelEqLink(ksel, esel func(common.Elem) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupBySelEq(ksel, esel, keq)
	}
}

// GroupByResLink produces GroupByResEr-based Link.
func GroupByResLink(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.GroupByRes(ksel, rsel) }
}

// GroupByResEqLink produces GroupByResEqEr-based Link.
func GroupByResEqLink(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupByResEq(ksel, rsel, keq)
	}
}

// GroupByLink produces GroupByEr-based Link.
func GroupByLink(ksel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.GroupBy(ksel) }
}

// GroupByEqLink produces GroupByEqEr-based Link.
func GroupByEqLink(ksel func(common.Elem) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.GroupByEq(ksel, keq) }
}

// GroupJoinEqLink produces GroupJoinEqEr-based Link.
func GroupJoinEqLink(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupJoinEq(inner, oksel, iksel, rsel, keq)
	}
}

// GroupJoinLink produces GroupJoinEr-based Link.
func GroupJoinLink(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		return en.GroupJoin(inner, oksel, iksel, rsel)
	}
}

// TakeLink produces Take-based Link.
func TakeLink(count int) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Take(count), nil }
}

// SkipLink produces Skip-based Link.
func SkipLink(count int) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Skip(count), nil }
}

// TakeWhileLink produces TakeWhileEr-based Link.
func TakeWhileLink(pred common.Predicate) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.TakeWhile(pred) }
}

// TakeWhileIdxLink produces TakeWhileIdxEr-based Link.
func TakeWhileIdxLink(pred common.PredicateIdx) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.TakeWhileIdx(pred) }
}

// SkipWhileLink produces SkipWhileEr-based Link.
func SkipWhileLink(pred common.Predicate) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.SkipWhile(pred) }
}

// SkipWhileIdxLink produces SkipWhileIdxEr-based Link.
func SkipWhileIdxLink(pred common.PredicateIdx) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.SkipWhileIdx(pred) }
}

// ToDictionarySelLink produces ToDictionarySelEr-based Link.
// The resulting Link returns enumerable.Enumerable of KeyValues containing in the Dictionary.
func ToDictionarySelLink(ksel, esel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		d, err := en.ToDictionarySel(ksel, esel)
		if err != nil {
			return nil, err
		}
		return d.Enumerable(), nil
	}
}

// ToDictionaryLink produces ToDictionaryEr-based Link.
// The resulting Link returns enumerable.Enumerable of KeyValues containing in the Dictionary.
func ToDictionaryLink(ksel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		d, err := en.ToDictionary(ksel)
		if err != nil {
			return nil, err
		}
		return d.Enumerable(), nil
	}
}

// OrderByLink produces OrderByEr-based Link.
// The resulting Link returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderByEr.
func OrderByLink(ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		oe, err := en.OrderBy(ls)
		if err != nil {
			return nil, err
		}
		return oe.Enumerable(), nil
	}
}

// OrderBySelLink produces OrderBySelEr-based Link.
// The resulting Link returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderBySelEr.
func OrderBySelLink(ls common.Less, ksel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		oe, err := en.OrderBySel(ls, ksel)
		if err != nil {
			return nil, err
		}
		return oe.Enumerable(), nil
	}
}

// OrderByDescendingLink produces OrderByDescendingEr-based Link.
// The resulting Link returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderByDescendingEr.
func OrderByDescendingLink(ls common.Less) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		oe, err := en.OrderByDescending(ls)
		if err != nil {
			return nil, err
		}
		return oe.Enumerable(), nil
	}
}

// OrderByDescendingSelLink produces OrderByDescendingSelEr-based Link.
// The resulting Link returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderByDescendingSelEr.
func OrderByDescendingSelLink(ls common.Less, ksel func(common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) {
		oe, err := en.OrderByDescendingSel(ls, ksel)
		if err != nil {
			return nil, err
		}
		return oe.Enumerable(), nil
	}
}

// ZipLink produces ZipEr-based Link.
func ZipLink(en2 *enumerable.Enumerable, sel func(common.Elem, common.Elem) common.Elem) Link {
	return func(en *enumerable.Enumerable) (*enumerable.Enumerable, error) { return en.Zip(en2, sel) }
}
