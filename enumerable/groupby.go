package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 21 - GroupBy
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-21-groupby/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

func groupByPrim(en *Enumerable,
	ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem,
	keq common.Equality) *Enumerable {
	lk := toLookupPrim(en, ksel, esel, keq, true)
	r := lk.Enumerable()
	if rsel == nil {
		return r
	}
	return r.SelectMust(func(el common.Elem) common.Elem {
		gr := el.(Grouping)
		return rsel(gr.key, gr.Enumerable())
	})
}

// GroupBySelRes groups the elements of Enumerable according to key selector 'ksel'
// and creates (using 'rsel') a result value from each group and its key.
// The elements of each group are projected using 'esel'.
// reflect.DeepEqual is used as keys Equality.
func (en *Enumerable) GroupBySelRes(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) (*Enumerable, error) {
	if ksel == nil || esel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	return groupByPrim(en, ksel, esel, rsel, reflect.DeepEqual), nil
}

// GroupBySelResMust is like GroupBySelRes but panics in case of error.
func (en *Enumerable) GroupBySelResMust(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) *Enumerable {
	r, err := en.GroupBySelRes(ksel, esel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelResEq groups the elements of Enumerable according to key selector 'ksel'
// and creates (using 'rsel') a result value from each group and its key.
// Key values are compared using 'keq' and the elements of each group are projected using 'esel'.
// If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) GroupBySelResEq(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem,
	keq common.Equality) (*Enumerable, error) {
	if ksel == nil || esel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return groupByPrim(en, ksel, esel, rsel, keq), nil
}

// GroupBySelResEqMust is like GroupBySelResEq but panics in case of error.
func (en *Enumerable) GroupBySelResEqMust(ksel, esel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem,
	keq common.Equality) *Enumerable {
	r, err := en.GroupBySelResEq(ksel, esel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySel groups the elements of Enumerable according to key selector 'ksel'
// and projects the elements of each group using 'esel'.
// reflect.DeepEqual is used as keys Equality.
func (en *Enumerable) GroupBySel(ksel, esel func(common.Elem) common.Elem) (*Enumerable, error) {
	if ksel == nil || esel == nil {
		return nil, errors.NilSel
	}
	return groupByPrim(en, ksel, esel, nil, reflect.DeepEqual), nil
}

// GroupBySelMust is like GroupBySel but panics in case of error.
func (en *Enumerable) GroupBySelMust(ksel, esel func(common.Elem) common.Elem) *Enumerable {
	r, err := en.GroupBySel(ksel, esel)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelEq groups the elements of Enumerable according to key selector 'ksel'.
// The keys are compared using 'keq' and each group's elements are projected using 'esel'.
// If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) GroupBySelEq(ksel, esel func(common.Elem) common.Elem, keq common.Equality) (*Enumerable, error) {
	if ksel == nil || esel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return groupByPrim(en, ksel, esel, nil, keq), nil
}

// GroupBySelEqMust is like GroupBySelEq but panics in case of error.
func (en *Enumerable) GroupBySelEqMust(ksel, esel func(common.Elem) common.Elem, keq common.Equality) *Enumerable {
	r, err := en.GroupBySelEq(ksel, esel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByRes groups the elements of Enumerable according to key selector 'ksel'
// and creates (using 'rsel') a result value from each group and its key.
// reflect.DeepEqual is used as keys Equality.
func (en *Enumerable) GroupByRes(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) (*Enumerable, error) {
	if ksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	return groupByPrim(en, ksel, nil, rsel, reflect.DeepEqual), nil
}

// GroupByResMust is like GroupByRes but panics in case of error.
func (en *Enumerable) GroupByResMust(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem) *Enumerable {
	r, err := en.GroupByRes(ksel, rsel)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByResEq groups the elements of Enumerable according to key selector 'ksel'
// and creates (using 'rsel') a result value from each group and its key.
// The keys are compared using 'keq'. If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) GroupByResEq(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) (*Enumerable, error) {
	if ksel == nil || rsel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return groupByPrim(en, ksel, nil, rsel, keq), nil
}

// GroupByResEqMust is like GroupByResEq but panics in case of error.
func (en *Enumerable) GroupByResEqMust(ksel func(common.Elem) common.Elem,
	rsel func(common.Elem, *Enumerable) common.Elem, keq common.Equality) *Enumerable {
	r, err := en.GroupByResEq(ksel, rsel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBy groups the elements of Enumerable according to key selector 'ksel' using reflect.DeepEqual as keys Equality.
func (en *Enumerable) GroupBy(ksel func(common.Elem) common.Elem) (*Enumerable, error) {
	if ksel == nil {
		return nil, errors.NilSel
	}
	return groupByPrim(en, ksel, nil, nil, reflect.DeepEqual), nil
}

// GroupByMust is like GroupBy but panics in case of error.
func (en *Enumerable) GroupByMust(ksel func(common.Elem) common.Elem) *Enumerable {
	r, err := en.GroupBy(ksel)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByEq groups the elements of Enumerable according to key selector 'ksel' using 'keq' as keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) GroupByEq(ksel func(common.Elem) common.Elem, keq common.Equality) (*Enumerable, error) {
	if ksel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return groupByPrim(en, ksel, nil, nil, keq), nil
}

// GroupByEqMust is like GroupByEq but panics in case of error.
func (en *Enumerable) GroupByEqMust(ksel func(common.Elem) common.Elem, keq common.Equality) *Enumerable {
	r, err := en.GroupByEq(ksel, keq)
	if err != nil {
		panic(err)
	}
	return r
}
