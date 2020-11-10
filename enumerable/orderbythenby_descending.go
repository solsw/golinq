package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// OrderBy sorts (using 'ls') the elements of Enumerable in ascending order.
func (en *Enumerable) OrderBy(ls common.Less) (*OrderedEnumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	return &OrderedEnumerable{en, ls}, nil
}

// OrderByMust is like OrderBy but panics in case of error.
func (en *Enumerable) OrderByMust(ls common.Less) *OrderedEnumerable {
	r, err := en.OrderBy(ls)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderBySel sorts (using 'ls') the elements of Enumerable in ascending order of keys obtained by 'ksel'.
func (en *Enumerable) OrderBySel(ls common.Less, ksel func(common.Elem) common.Elem) (*OrderedEnumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	if ksel == nil {
		return nil, errors.NilSel
	}
	ls2 := projectionLess(ls, ksel)
	return &OrderedEnumerable{en, ls2}, nil
}

// OrderBySelMust is like OrderBySel but panics in case of error.
func (en *Enumerable) OrderBySelMust(ls common.Less, ksel func(common.Elem) common.Elem) *OrderedEnumerable {
	r, err := en.OrderBySel(ls, ksel)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescending sorts (using 'ls') the elements of Enumerable in descending order.
func (en *Enumerable) OrderByDescending(ls common.Less) (*OrderedEnumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	ls2 := reverseLess(ls)
	return &OrderedEnumerable{en, ls2}, nil
}

// OrderByDescendingMust is like OrderByDescending but panics in case of error.
func (en *Enumerable) OrderByDescendingMust(ls common.Less) *OrderedEnumerable {
	r, err := en.OrderByDescending(ls)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescendingSel sorts (using 'ls') the elements of Enumerable in descending order of keys obtained by 'ksel'.
func (en *Enumerable) OrderByDescendingSel(ls common.Less, ksel func(common.Elem) common.Elem) (*OrderedEnumerable, error) {
	if ls == nil {
		return nil, errors.NilLess
	}
	if ksel == nil {
		return nil, errors.NilSel
	}
	ls2 := projectionLess(ls, ksel)
	ls3 := reverseLess(ls2)
	return &OrderedEnumerable{en, ls3}, nil
}

// OrderByDescendingSelMust is like OrderByDescendingSel but panics in case of error.
func (en *Enumerable) OrderByDescendingSelMust(ls common.Less, ksel func(common.Elem) common.Elem) *OrderedEnumerable {
	r, err := en.OrderByDescendingSel(ls, ksel)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenBy performs a subsequent ordering (using 'ls') of OrderedEnumerable in ascending order.
func (oe *OrderedEnumerable) ThenBy(ls common.Less) (*OrderedEnumerable, error) {
	if oe.ls == nil || ls == nil {
		return nil, errors.NilLess
	}
	ls2 := compoundLess(oe.ls, ls)
	return &OrderedEnumerable{oe.en, ls2}, nil
}

// ThenByMust is like ThenBy but panics in case of error.
func (oe *OrderedEnumerable) ThenByMust(ls common.Less) *OrderedEnumerable {
	r, err := oe.ThenBy(ls)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenBySel performs a subsequent ordering (using 'ls') of OrderedEnumerable in ascending order of keys obtained by 'ksel'.
func (oe *OrderedEnumerable) ThenBySel(ls common.Less, ksel func(common.Elem) common.Elem) (*OrderedEnumerable, error) {
	if oe.ls == nil || ls == nil {
		return nil, errors.NilLess
	}
	if ksel == nil {
		return nil, errors.NilSel
	}
	ls2 := projectionLess(ls, ksel)
	ls3 := compoundLess(oe.ls, ls2)
	return &OrderedEnumerable{oe.en, ls3}, nil
}

// ThenBySelMust is like ThenBySel but panics in case of error.
func (oe *OrderedEnumerable) ThenBySelMust(ls common.Less, ksel func(common.Elem) common.Elem) *OrderedEnumerable {
	r, err := oe.ThenBySel(ls, ksel)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDescending performs a subsequent ordering (using 'ls') of OrderedEnumerable in descending order.
func (oe *OrderedEnumerable) ThenByDescending(ls common.Less) (*OrderedEnumerable, error) {
	if oe.ls == nil || ls == nil {
		return nil, errors.NilLess
	}
	ls2 := reverseLess(ls)
	ls3 := compoundLess(oe.ls, ls2)
	return &OrderedEnumerable{oe.en, ls3}, nil
}

// ThenByDescendingMust is like ThenByDescending but panics in case of error.
func (oe *OrderedEnumerable) ThenByDescendingMust(ls common.Less) *OrderedEnumerable {
	r, err := oe.ThenByDescending(ls)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDescendingSel performs a subsequent ordering (using 'ls') of OrderedEnumerable in descending order of keys obtained by 'ksel'.
func (oe *OrderedEnumerable) ThenByDescendingSel(ls common.Less, ksel func(common.Elem) common.Elem) (*OrderedEnumerable, error) {
	if oe.ls == nil || ls == nil {
		return nil, errors.NilLess
	}
	if ksel == nil {
		return nil, errors.NilSel
	}
	ls2 := projectionLess(ls, ksel)
	ls3 := reverseLess(ls2)
	ls4 := compoundLess(oe.ls, ls3)
	return &OrderedEnumerable{oe.en, ls4}, nil
}

// ThenByDescendingSelMust is like ThenByDescendingSel but panics in case of error.
func (oe *OrderedEnumerable) ThenByDescendingSelMust(ls common.Less, ksel func(common.Elem) common.Elem) *OrderedEnumerable {
	r, err := oe.ThenByDescendingSel(ls, ksel)
	if err != nil {
		panic(err)
	}
	return r
}
