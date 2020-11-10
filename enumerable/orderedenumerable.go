package enumerable

import (
	"sort"

	"github.com/solsw/golinq/common"
)

// Reimplementing LINQ to Objects: Part 26a – IOrderedEnumerable
// https://codeblog.jonskeet.uk/2011/01/04/reimplementing-linq-to-objects-part-26a-iorderedenumerable/
// https://docs.microsoft.com/dotnet/api/system.linq.iorderedenumerable-1

// OrderedEnumerable represents a sorted Enumerable.
//
// OrderedEnumerable itself does NOT contain sorted data.
// Instead sorted Enumerable is obtained from OrderedEnumerable with the help of Enumerable() method.
type OrderedEnumerable struct {
	en *Enumerable
	ls common.Less
}

// Enumerable converts OrderedEnumerable to Enumerable using sort.SliceStable for sorting.
func (oe *OrderedEnumerable) Enumerable() *Enumerable {
	var r common.Slice
	for oe.en.MoveNext() {
		r = append(r, oe.en.Current())
	}
	sort.SliceStable(r, func(i, j int) bool {
		return oe.ls(r[i], r[j])
	})
	return NewElems(r...)
}
