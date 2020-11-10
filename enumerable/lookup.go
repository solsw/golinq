package enumerable

import (
	"reflect"
	"strings"

	"github.com/solsw/gohelpers/oshelper"
	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Lookup.cs

// Lookup represents a collection of keys each mapped to one or more values.
type Lookup struct {
	// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2

	grgr []Grouping
	// keyEq is equality comparison for grgr's keys.
	keyEq common.Equality
}

// newLookupEq creates new empty Lookup with keys Equality.
func newLookupEq(keq common.Equality) *Lookup {
	return &Lookup{keyEq: keq}
}

// newLookup creates new empty Lookup using reflect.DeepEqual as keys Equality.
func newLookup() *Lookup {
	return newLookupEq(reflect.DeepEqual)
}

func (lk *Lookup) keyIndex(key common.Elem) int {
	for i, g := range lk.grgr {
		if lk.keyEq(g.key, key) {
			return i
		}
	}
	return -1
}

// add adds element 'el' with specified 'key' to 'lk'.
func (lk *Lookup) add(key, el common.Elem) {
	i := lk.keyIndex(key)
	if i >= 0 {
		lk.grgr[i].values = append(lk.grgr[i].values, el)
	} else {
		gr := Grouping{key: key, values: common.Slice{el}}
		lk.grgr = append(lk.grgr, gr)
	}
}

// Count returns number of keys in Lookup.
func (lk *Lookup) Count() int {
	// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.count
	return len(lk.grgr)
}

// Item gets the collection of values indexed by the specified 'key'.
func (lk *Lookup) Item(key common.Elem) *Enumerable {
	// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.item
	i := lk.keyIndex(key)
	if i < 0 {
		return Empty()
	}
	return NewElems(lk.grgr[i].values...)
}

// Contains determines whether a specified 'key' is in the Lookup.
func (lk *Lookup) Contains(key common.Elem) bool {
	// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.contains
	return lk.keyIndex(key) >= 0
}

// Equal checks if two Lookups are equal.
// Keys Equalities do not participate in equality verification,
// since non-nil funcs are always not deeply equal.
func (lk *Lookup) Equal(lk2 *Lookup) bool {
	if lk.Count() != lk2.Count() {
		return false
	}
	for i, g := range lk.grgr {
		g2 := lk2.grgr[i]
		if !lk.keyEq(g.key, g2.key) || !reflect.DeepEqual(g.values, g2.values) {
			return false
		}
	}
	return true
}

// Slice converts Lookup to slice of Groupings.
func (lk *Lookup) Slice() common.Slice {
	r := make(common.Slice, 0, len(lk.grgr))
	for _, g := range lk.grgr {
		r = append(r, g)
	}
	return r
}

// Enumerable converts Lookup to Enumerable of Groupings.
func (lk *Lookup) Enumerable() *Enumerable {
	return NewElems(lk.Slice()...)
}

// String implements the fmt.Stringer interface.
func (lk *Lookup) String() string {
	var b strings.Builder
	for _, g := range lk.grgr {
		if b.Len() > 0 {
			b.WriteString(oshelper.NewLine)
		}
		b.WriteString(g.String())
	}
	return b.String()
}
