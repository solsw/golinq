package chain

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

type (
	// LinkMust projects one enumerable.Enumerable into another.
	LinkMust = func(*enumerable.Enumerable) *enumerable.Enumerable
)

// ChainMust sequentially applies 'links' to 'en' and returns the resulting enumerable.Enumerable.
func ChainMust(en *enumerable.Enumerable, links ...LinkMust) *enumerable.Enumerable {
	if en == nil || len(links) == 0 {
		return en
	}
	r := links[0](en)
	for i := 1; i < len(links); i++ {
		r = links[i](r)
	}
	return r
}

// WhereLinkMust produces Where-based Link.
func WhereLinkMust(pred common.Predicate) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.WhereMust(pred) }
}

// WhereIdxLinkMust produces WhereIdx-based Link.
func WhereIdxLinkMust(pred common.PredicateIdx) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.WhereIdxMust(pred) }
}

// SelectLinkMust produces Select-based Link.
func SelectLinkMust(sel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectMust(sel) }
}

// SelectIdxLinkMust produces SelectIdx-based Link.
func SelectIdxLinkMust(sel func(common.Elem, int) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectIdxMust(sel) }
}

// RangeLinkMust produces Range-based Link.
func RangeLinkMust(start, count int) LinkMust {
	return func(*enumerable.Enumerable) *enumerable.Enumerable { return enumerable.RangeMust(start, count) }
}

// EmptyLinkMust produces Empty-based Link.
func EmptyLinkMust() LinkMust {
	return func(*enumerable.Enumerable) *enumerable.Enumerable { return enumerable.Empty() }
}

// RepeatLinkMust produces Repeat-based Link.
func RepeatLinkMust(el common.Elem, count int) LinkMust {
	return func(*enumerable.Enumerable) *enumerable.Enumerable { return enumerable.RepeatMust(el, count) }
}

// ConcatLinkMust produces Concat-based Link.
func ConcatLinkMust(en2 *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Concat(en2) }
}

// SelectManyLinkMust produces SelectMany-based Link.
func SelectManyLinkMust(sel func(common.Elem) *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectManyMust(sel) }
}

// SelectManyIdxLinkMust produces SelectManyIdx-based Link.
func SelectManyIdxLinkMust(sel func(common.Elem, int) *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectManyIdxMust(sel) }
}

// SelectManyCollLinkMust produces SelectManyColl-based Link.
func SelectManyCollLinkMust(sel1 func(common.Elem) *enumerable.Enumerable,
	sel2 func(common.Elem, common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectManyCollMust(sel1, sel2) }
}

// SelectManyIdxCollLinkMust produces SelectManyIdxColl-based Link.
func SelectManyIdxCollLinkMust(sel1 func(common.Elem, int) *enumerable.Enumerable,
	sel2 func(common.Elem, common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SelectManyIdxCollMust(sel1, sel2) }
}

// DefaultIfEmptyLinkMust produces DefaultIfEmpty-based Link.
func DefaultIfEmptyLinkMust() LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.DefaultIfEmpty() }
}

// DefaultIfEmptyDefLinkMust produces DefaultIfEmptyDef-based Link.
func DefaultIfEmptyDefLinkMust(def common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.DefaultIfEmptyDef(def) }
}

// DistinctLinkMust produces Distinct-based Link.
func DistinctLinkMust() LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Distinct() }
}

// DistinctEqLinkMust produces DistinctEq-based Link.
func DistinctEqLinkMust(eq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.DistinctEq(eq) }
}

// DistinctCmpLinkMust produces DistinctCmp-based Link.
func DistinctCmpLinkMust(cmp common.Comparison) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.DistinctCmpMust(cmp) }
}

// DistinctLsLinkMust produces DistinctLs-based Link.
func DistinctLsLinkMust(ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.DistinctLsMust(ls) }
}

// UnionLinkMust produces Union-based Link.
func UnionLinkMust(en2 *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Union(en2) }
}

// UnionEqLinkMust produces UnionEq-based Link.
func UnionEqLinkMust(en2 *enumerable.Enumerable, eq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.UnionEq(en2, eq) }
}

// UnionCmpLinkMust produces UnionCmp-based Link.
func UnionCmpLinkMust(en2 *enumerable.Enumerable, cmp common.Comparison) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.UnionCmpMust(en2, cmp) }
}

// UnionLsLinkMust produces UnionLs-based Link.
func UnionLsLinkMust(en2 *enumerable.Enumerable, ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.UnionLsMust(en2, ls) }
}

// IntersectLinkMust produces Intersect-based Link.
func IntersectLinkMust(en2 *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Intersect(en2) }
}

// IntersectEqLinkMust produces IntersectEq-based Link.
func IntersectEqLinkMust(en2 *enumerable.Enumerable, eq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.IntersectEq(en2, eq) }
}

// IntersectCmpLinkMust produces IntersectCmp-based Link.
func IntersectCmpLinkMust(en2 *enumerable.Enumerable, cmp common.Comparison) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.IntersectCmpMust(en2, cmp) }
}

// IntersectLsLinkMust produces IntersectLs-based Link.
func IntersectLsLinkMust(en2 *enumerable.Enumerable, ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.IntersectLsMust(en2, ls) }
}

// ExceptLinkMust produces Except-based Link.
func ExceptLinkMust(en2 *enumerable.Enumerable) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Except(en2) }
}

// ExceptEqLinkMust produces ExceptEq-based Link.
func ExceptEqLinkMust(en2 *enumerable.Enumerable, eq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.ExceptEq(en2, eq) }
}

// ExceptCmpLinkMust produces ExceptCmp-based Link.
func ExceptCmpLinkMust(en2 *enumerable.Enumerable, cmp common.Comparison) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.ExceptCmpMust(en2, cmp) }
}

// ExceptLsLinkMust produces ExceptLs-based Link.
func ExceptLsLinkMust(en2 *enumerable.Enumerable, ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.ExceptLsMust(en2, ls) }
}

// ToLookupSelEqLinkMust produces ToLookupSelEq-based Link.
// The resulting LinkMust returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupSelEqLinkMust(ksel, esel func(common.Elem) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToLookupSelEqMust(ksel, esel, keq).Enumerable()
	}
}

// ToLookupSelLinkMust produces ToLookupSel-based Link.
// The resulting LinkMust returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupSelLinkMust(ksel, esel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToLookupSelMust(ksel, esel).Enumerable()
	}
}

// ToLookupEqLinkMust produces ToLookupEq-based Link.
// The resulting LinkMust returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupEqLinkMust(ksel func(common.Elem) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToLookupEqMust(ksel, keq).Enumerable()
	}
}

// ToLookupLinkMust produces ToLookup-based Link.
// The resulting LinkMust returns enumerable.Enumerable of Groupings containing in the Lookup.
func ToLookupLinkMust(ksel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToLookupMust(ksel).Enumerable()
	}
}

// JoinEqLinkMust produces JoinEq-based Link.
func JoinEqLinkMust(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.JoinEqMust(inner, oksel, iksel, rsel, keq)
	}
}

// JoinLinkMust produces Join-based Link.
func JoinLinkMust(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.JoinMust(inner, oksel, iksel, rsel) }
}

// GroupBySelResLinkMust produces GroupBySelRes-based Link.
func GroupBySelResLinkMust(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupBySelResMust(ksel, esel, rsel) }
}

// GroupBySelResEqLinkMust produces GroupBySelResEq-based Link.
func GroupBySelResEqLinkMust(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.GroupBySelResEqMust(ksel, esel, rsel, keq)
	}
}

// GroupBySelLinkMust produces GroupBySel-based Link.
func GroupBySelLinkMust(ksel, esel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupBySelMust(ksel, esel) }
}

// GroupBySelEqLinkMust produces GroupBySelEq-based Link.
func GroupBySelEqLinkMust(ksel, esel func(common.Elem) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupBySelEqMust(ksel, esel, keq) }
}

// GroupByResLinkMust produces GroupByRes-based Link.
func GroupByResLinkMust(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupByResMust(ksel, rsel) }
}

// GroupByResEqLinkMust produces GroupByResEq-based Link.
func GroupByResEqLinkMust(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupByResEqMust(ksel, rsel, keq) }
}

// GroupByLinkMust produces GroupBy-based Link.
func GroupByLinkMust(ksel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupByMust(ksel) }
}

// GroupByEqLinkMust produces GroupByEq-based Link.
func GroupByEqLinkMust(ksel func(common.Elem) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.GroupByEqMust(ksel, keq) }
}

// GroupJoinEqLinkMust produces GroupJoinEq-based Link.
func GroupJoinEqLinkMust(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem, keq common.Equality) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.GroupJoinEqMust(inner, oksel, iksel, rsel, keq)
	}
}

// GroupJoinLinkMust produces GroupJoin-based Link.
func GroupJoinLinkMust(inner *enumerable.Enumerable, oksel, iksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *enumerable.Enumerable) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.GroupJoinMust(inner, oksel, iksel, rsel)
	}
}

// TakeLinkMust produces Take-based Link.
func TakeLinkMust(count int) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Take(count) }
}

// SkipLinkMust produces Skip-based Link.
func SkipLinkMust(count int) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.Skip(count) }
}

// TakeWhileLinkMust produces TakeWhile-based Link.
func TakeWhileLinkMust(pred common.Predicate) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.TakeWhileMust(pred) }
}

// TakeWhileIdxLinkMust produces TakeWhileIdx-based Link.
func TakeWhileIdxLinkMust(pred common.PredicateIdx) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.TakeWhileIdxMust(pred) }
}

// SkipWhileLinkMust produces SkipWhile-based Link.
func SkipWhileLinkMust(pred common.Predicate) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SkipWhileMust(pred) }
}

// SkipWhileIdxLinkMust produces SkipWhileIdx-based Link.
func SkipWhileIdxLinkMust(pred common.PredicateIdx) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.SkipWhileIdxMust(pred) }
}

// ToDictionarySelLinkMust produces ToDictionarySel-based Link.
// The resulting LinkMust returns enumerable.Enumerable of KeyValues containing in the Dictionary.
func ToDictionarySelLinkMust(ksel, esel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToDictionarySelMust(ksel, esel).Enumerable()
	}
}

// ToDictionaryLinkMust produces ToDictionary-based Link.
// The resulting LinkMust returns enumerable.Enumerable of KeyValues containing in the Dictionary.
func ToDictionaryLinkMust(ksel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.ToDictionaryMust(ksel).Enumerable()
	}
}

// OrderByLinkMust produces OrderBy-based Link.
// The resulting LinkMust returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderBy.
func OrderByLinkMust(ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.OrderByMust(ls).Enumerable() }
}

// OrderBySelLinkMust produces OrderBySel-based Link.
// The resulting LinkMust returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderBySel.
func OrderBySelLinkMust(ls common.Less, ksel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.OrderBySelMust(ls, ksel).Enumerable()
	}
}

// OrderByDescendingLinkMust produces OrderByDescending-based Link.
// The resulting LinkMust returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderByDescending.
func OrderByDescendingLinkMust(ls common.Less) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.OrderByDescendingMust(ls).Enumerable()
	}
}

// OrderByDescendingSelLinkMust produces OrderByDescendingSel-based Link.
// The resulting LinkMust returns enumerable.Enumerable based on the OrderedEnumerable returned by OrderByDescendingSel.
func OrderByDescendingSelLinkMust(ls common.Less, ksel func(common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable {
		return en.OrderByDescendingSelMust(ls, ksel).Enumerable()
	}
}

// ZipLinkMust produces Zip-based Link.
func ZipLinkMust(en2 *enumerable.Enumerable, sel func(common.Elem, common.Elem) common.Elem) LinkMust {
	return func(en *enumerable.Enumerable) *enumerable.Enumerable { return en.ZipMust(en2, sel) }
}
