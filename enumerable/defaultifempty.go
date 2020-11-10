package enumerable

import (
	"github.com/solsw/golinq/common"
)

// Reimplementing LINQ to Objects: Part 12 – DefaultIfEmpty
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-12-defaultifempty/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

// DefaultIfEmpty returns 'en' or nil in a singleton Enumerable if 'en' is empty.
// 'en' may need Reset() for further use.
func (en *Enumerable) DefaultIfEmpty() *Enumerable {
	return en.DefaultIfEmptyDef(nil)
}

// DefaultIfEmptyDef returns 'en' or the specified value in a singleton Enumerable if 'en' is empty.
// 'en' may need Reset() for further use.
func (en *Enumerable) DefaultIfEmptyDef(def common.Elem) *Enumerable {
	if !en.Any() {
		return NewElems(def)
	}
	return en
}
