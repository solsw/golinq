package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
)

// Reimplementing LINQ to Objects: Part 32 – Contains
// https://codeblog.jonskeet.uk/2011/01/12/reimplementing-linq-to-objects-part-32-contains/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

// Contains determines whether Enumerable contains the specified element using reflect.DeepEqual as Equality.
func (en *Enumerable) Contains(el common.Elem) bool {
	return en.ContainsEq(el, reflect.DeepEqual)
}

// ContainsEq determines whether Enumerable contains the specified element
// using Equality ('el' is passed to 'eq' as first argument).
// If 'eq' is nil reflect.DeepEqual is used.
func (en *Enumerable) ContainsEq(el common.Elem, eq common.Equality) bool {
	if eq == nil {
		eq = reflect.DeepEqual
	}
	for en.MoveNext() {
		if eq(el, en.Current()) {
			return true
		}
	}
	return false
}
