package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 9 – SelectMany
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-9-selectmany/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

// SelectMany projects each element of 'en' into Enumerable
// and flattens the resulting Enumerables into one Enumerable.
func (en *Enumerable) SelectMany(sel func(common.Elem) *Enumerable) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	t := Empty()
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !en.MoveNext() {
						return false
					}
					t = sel(en.Current())
				}
			},
			Crrnt: func() common.Elem { return t.Current() },
			Rst:   func() { t = Empty(); en.Reset() },
		}},
		nil
}

// SelectManyMust is like SelectMany but panics in case of error.
func (en *Enumerable) SelectManyMust(sel func(common.Elem) *Enumerable) *Enumerable {
	r, err := en.SelectMany(sel)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyIdx projects each element of 'en' and its index into Enumerable
// and flattens the resulting Enumerables into one Enumerable.
func (en *Enumerable) SelectManyIdx(sel func(common.Elem, int) *Enumerable) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	var (
		i int         = -1
		t *Enumerable = Empty()
	)
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !en.MoveNext() {
						return false
					}
					i++
					t = sel(en.Current(), i)
				}
			},
			Crrnt: func() common.Elem { return t.Current() },
			Rst:   func() { i = -1; t = Empty(); en.Reset() },
		}},
		nil
}

// SelectManyIdxMust is like SelectManyIdx but panics in case of error.
func (en *Enumerable) SelectManyIdxMust(sel func(common.Elem, int) *Enumerable) *Enumerable {
	r, err := en.SelectManyIdx(sel)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyColl projects (using 'sel1') each element of 'en' into Enumerable,
// flattens the resulting Enums into one Enumerable and
// projects (using 'sel2') initial element with each corresponding projected elements into elements of resulting Enumerable.
func (en *Enumerable) SelectManyColl(sel1 func(common.Elem) *Enumerable, sel2 func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	if sel1 == nil || sel2 == nil {
		return nil, errors.NilSel
	}
	var e1 common.Elem
	t := Empty()
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !en.MoveNext() {
						return false
					}
					e1 = en.Current()
					t = sel1(e1)
				}
			},
			Crrnt: func() common.Elem { return sel2(e1, t.Current()) },
			Rst:   func() { t = Empty(); en.Reset() },
		}},
		nil
}

// SelectManyCollMust is like SelectManyColl but panics in case of error.
func (en *Enumerable) SelectManyCollMust(sel1 func(common.Elem) *Enumerable, sel2 func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := en.SelectManyColl(sel1, sel2)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyIdxColl projects (using 'sel1') each element of 'en' and its index into Enumerable,
// flattens the resulting Enums into one Enumerable and
// projects (using 'sel2') initial element with each corresponding projected elements into elements of resulting Enumerable.
func (en *Enumerable) SelectManyIdxColl(sel1 func(common.Elem, int) *Enumerable, sel2 func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	if sel1 == nil || sel2 == nil {
		return nil, errors.NilSel
	}
	var e1 common.Elem
	i := -1
	t := Empty()
	return &Enumerable{enumerator.OnFuncs{
			MvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !en.MoveNext() {
						return false
					}
					e1 = en.Current()
					i++
					t = sel1(e1, i)
				}
			},
			Crrnt: func() common.Elem { return sel2(e1, t.Current()) },
			Rst:   func() { i = -1; t = Empty(); en.Reset() },
		}},
		nil
}

// SelectManyIdxCollMust is like SelectManyIdxColl but panics in case of error.
func (en *Enumerable) SelectManyIdxCollMust(sel1 func(common.Elem, int) *Enumerable, sel2 func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := en.SelectManyIdxColl(sel1, sel2)
	if err != nil {
		panic(err)
	}
	return r
}
