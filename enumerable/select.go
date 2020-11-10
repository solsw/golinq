package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 3 – "Select" (and a rename...)
// https://codeblog.jonskeet.uk/2010/12/23/reimplementing-linq-to-objects-part-3-quot-select-quot-and-a-rename/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

// see example/selectmany/selectmanyex3

func selectPrim(en *Enumerable, sel func(common.Elem) common.Elem) *Enumerable {
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: en.MoveNext,
		Crrnt: func() common.Elem { return sel(en.Current()) },
		Rst:   en.Reset,
	}}
}

// Select projects each element of the Enumerable into a new form.
func (en *Enumerable) Select(sel func(common.Elem) common.Elem) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	return selectPrim(en, sel), nil
}

// SelectMust is like Select but panics in case of error.
func (en *Enumerable) SelectMust(sel func(common.Elem) common.Elem) *Enumerable {
	r, err := en.Select(sel)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectIdx projects each element of the Enumerable into a new form by incorporating the element's index.
func (en *Enumerable) SelectIdx(sel func(common.Elem, int) common.Elem) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	var i int = -1 // position before the first element
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool { i++; return en.MoveNext() },
			Crrnt: func() common.Elem { return sel(en.Current(), i) },
			Rst:   func() { i = -1; en.Reset() },
		}},
		nil
}

// SelectIdxMust is like SelectIdx but panics in case of error.
func (en *Enumerable) SelectIdxMust(sel func(common.Elem, int) common.Elem) *Enumerable {
	r, err := en.SelectIdx(sel)
	if err != nil {
		panic(err)
	}
	return r
}
