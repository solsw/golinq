package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
)

// Reimplementing LINQ to Objects: Part 8 – Concat
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-8-concat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat

// Concat concatenates two Enumerables.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use ConcatSelf instead.
func (en *Enumerable) Concat(en2 *Enumerable) *Enumerable {
	from1 := true
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			if from1 && en.MoveNext() {
				return true
			}
			from1 = false
			return en2.MoveNext()
		},
		Crrnt: func() common.Elem {
			if from1 {
				return en.Current()
			}
			return en2.Current()
		},
		Rst: func() { from1 = true; en.Reset(); en2.Reset() },
	}}
}

// ConcatSelf concatenates two Enumerables.
// 'en' and 'en2' may be based on the same Enumerable,
// but this Enumerable must have fully implemented Reset method.
// (See TestEnumerable_ConcatSelf for examples.)
func (en *Enumerable) ConcatSelf(en2 *Enumerable) *Enumerable {
	sl2 := en2.Slice()
	en.Reset()
	return en.Concat(NewElems(sl2...))
}
