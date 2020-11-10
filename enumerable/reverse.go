package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
)

// Reimplementing LINQ to Objects: Part 27 – Reverse
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-27-reverse/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse

// Reverse inverts the order of the elements in a sequence.
func (en *Enumerable) Reverse() *Enumerable {
	sl := en.Slice()
	i := len(sl)
	return &Enumerable{enumerator.OnFuncs{
		MvNxt: func() bool {
			if i > 0 {
				i--
				return true
			}
			return false
		},
		Crrnt: func() common.Elem { return sl[i] },
		Rst:   func() { i = len(sl); en.Reset() },
	}}
}
