package enumerator

import (
	"github.com/solsw/golinq/common"
)

// OnSlice is Enumerator implementation based on slice of common.Elem.
type OnSlice struct {
	// indx-1 - index of the current element in elel
	//     ^^ because initially enumerator is positioned before the first element in the collection
	indx int
	// elements of the OnSlice instance
	elel common.Slice
}

// NewOnSlice creates a new OnSlice with the specified contents.
func NewOnSlice(ee ...common.Elem) *OnSlice {
	var en OnSlice
	en.elel = make(common.Slice, len(ee))
	if len(ee) > 0 {
		copy(en.elel, ee)
	}
	return &en
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *OnSlice) MoveNext() bool {
	if en.indx == len(en.elel) {
		return false
	}
	en.indx++
	return true
}

// Current implements the Enumerator.Current method.
func (en *OnSlice) Current() common.Elem {
	return en.Item(en.indx - 1)
}

// Reset implements the Enumerator.Reset method.
func (en *OnSlice) Reset() {
	en.indx = 0
}

// Count implements the common.Counter interface.
func (en *OnSlice) Count() int {
	return len(en.elel)
}

// Item implements the common.Itemer interface.
func (en *OnSlice) Item(i int) common.Elem {
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.current#remarks
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.list-1.item#exceptions
	if !(0 <= i && i < len(en.elel)) {
		return nil
	}
	return en.elel[i]
}
