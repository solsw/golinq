package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 35 – Zip
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-35-zip/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

func zipPrim(en1, en2 *Enumerable, sel func(common.Elem, common.Elem) common.Elem) *Enumerable {
	// 'en1' and 'en2' must NOT be based on the same Enumerable
	var c common.Elem
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			if en1.MoveNext() && en2.MoveNext() {
				c = sel(en1.Current(), en2.Current())
				return true
			}
			return false
		},
		Crrnt: func() common.Elem { return c },
		Rst:   func() { en1.Reset(); en2.Reset() },
	}}
}

// Zip applies a specified selector to the corresponding elements
// of two Enumerables, producing Enumerable of the results.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ZipSelf instead.
func (en *Enumerable) Zip(en2 *Enumerable, sel func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	return zipPrim(en, en2, sel), nil
}

// ZipMust is like Zip but panics in case of error.
func (en *Enumerable) ZipMust(en2 *Enumerable, sel func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := en.Zip(en2, sel)
	if err != nil {
		panic(err)
	}
	return r
}

// ZipSelf applies a specified selector to the corresponding elements
// of two Enumerables, producing Enumerable of the results.
// 'en' and 'en2' may be based on the same Enumerable,
// this Enumerable must have fully implemented Reset method.
// (See TestEnumerable_ZipSelf for examples.)
func (en *Enumerable) ZipSelf(en2 *Enumerable, sel func(common.Elem, common.Elem) common.Elem) (*Enumerable, error) {
	if sel == nil {
		return nil, errors.NilSel
	}
	sl2 := en2.Slice()
	en.Reset()
	return zipPrim(en, NewElems(sl2...), sel), nil
}

// ZipSelfMust is like ZipSelf but panics in case of error.
func (en *Enumerable) ZipSelfMust(en2 *Enumerable, sel func(common.Elem, common.Elem) common.Elem) *Enumerable {
	r, err := en.ZipSelf(en2, sel)
	if err != nil {
		panic(err)
	}
	return r
}
