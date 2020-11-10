package enumerable

import (
	"fmt"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Grouping.cs
// https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2

// Grouping represents a collection of objects that have a common key.
type Grouping struct {
	key    common.Elem
	values common.Slice
}

// Key returns the key of the Grouping.
func (gr *Grouping) Key() common.Elem {
	return gr.key
}

// Slice returns the Grouping's values as Slice.
func (gr *Grouping) Slice() common.Slice {
	return gr.values
}

// Enumerable returns the Grouping's values as Enumerable.
func (gr *Grouping) Enumerable() *Enumerable {
	return NewElems(gr.values...)
}

// String implements the fmt.Stringer interface.
func (gr *Grouping) String() string {
	return fmt.Sprintf("%v: %v", gr.key, gr.values)
}
