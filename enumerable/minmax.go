package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 29 – Min/Max
// https://codeblog.jonskeet.uk/2011/01/09/reimplementing-linq-to-objects-part-29-min-max/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

// 'sel' projects each element of 'en'
// 'ls' compares projected values
// if 'min', function searches for minimum, otherwise - for maximum
// if 'el', element of sequence is returned, otherwise - projected value
func minMaxPrim(en *Enumerable, sel func(common.Elem) common.Elem, ls common.Less, min, el bool) (common.Elem, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	if ls == nil {
		return nil, errors.NilLess
	}
	first := true
	var re, rs common.Elem
	for en.MoveNext() {
		if first {
			first = false
			re = en.Current()
			rs = sel(re)
			continue
		}
		e := en.Current()
		s := sel(e)
		if (min && ls(s, rs)) || (!min && ls(rs, s)) {
			re = e
			rs = s
		}
	}
	if first {
		return nil, errors.EmptyEnum
	}
	if el {
		return re, nil
	}
	return rs, nil
}

// MinSel invokes a transform function
// on each element of Enumerable and returns the minimum resulting value.
func (en *Enumerable) MinSel(ls common.Less, sel func(common.Elem) common.Elem) (common.Elem, error) {
	return minMaxPrim(en, sel, ls, true, false)
}

// MinSelMust is like MinSel but panics in case of error.
func (en *Enumerable) MinSelMust(ls common.Less, sel func(common.Elem) common.Elem) common.Elem {
	r, err := en.MinSel(ls, sel)
	if err != nil {
		panic(err)
	}
	return r
}

// MinSelEl invokes a transform function on each element
// of Enumerable and returns the element which produces the minimum resulting value.
func (en *Enumerable) MinSelEl(ls common.Less, sel func(common.Elem) common.Elem) (common.Elem, error) {
	return minMaxPrim(en, sel, ls, true, true)
}

// MinSelElMust is like MinSelEl but panics in case of error.
func (en *Enumerable) MinSelElMust(ls common.Less, sel func(common.Elem) common.Elem) common.Elem {
	r, err := en.MinSelEl(ls, sel)
	if err != nil {
		panic(err)
	}
	return r
}

// Min returns the minimum Enumerable's element.
func (en *Enumerable) Min(ls common.Less) (common.Elem, error) {
	return minMaxPrim(en, common.Identity, ls, true, true)
}

// MinMust is like Min but panics in case of error.
func (en *Enumerable) MinMust(ls common.Less) common.Elem {
	r, err := en.Min(ls)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxSel invokes a transform function
// on each element of Enumerable and returns the maximum resulting value.
func (en *Enumerable) MaxSel(ls common.Less, sel func(common.Elem) common.Elem) (common.Elem, error) {
	return minMaxPrim(en, sel, ls, false, false)
}

// MaxSelMust is like MaxSel but panics in case of error.
func (en *Enumerable) MaxSelMust(ls common.Less, sel func(common.Elem) common.Elem) common.Elem {
	r, err := en.MaxSel(ls, sel)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxSelEl invokes a transform function on each element
// of Enumerable and returns the element which produces the maximum resulting value.
func (en *Enumerable) MaxSelEl(ls common.Less, sel func(common.Elem) common.Elem) (common.Elem, error) {
	return minMaxPrim(en, sel, ls, false, true)
}

// MaxSelElMust is like MaxSelEl but panics in case of error.
func (en *Enumerable) MaxSelElMust(ls common.Less, sel func(common.Elem) common.Elem) common.Elem {
	r, err := en.MaxSelEl(ls, sel)
	if err != nil {
		panic(err)
	}
	return r
}

// Max returns the maximum Enumerable's element.
func (en *Enumerable) Max(ls common.Less) (common.Elem, error) {
	return minMaxPrim(en, common.Identity, ls, false, true)
}

// MaxMust is like Max but panics in case of error.
func (en *Enumerable) MaxMust(ls common.Less) common.Elem {
	r, err := en.Max(ls)
	if err != nil {
		panic(err)
	}
	return r
}
