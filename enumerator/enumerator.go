package enumerator

import (
	"github.com/solsw/golinq/common"
)

// Enumerator supports a simple iteration over a collection
// (https://docs.microsoft.com/dotnet/api/system.collections.ienumerator).
type Enumerator interface {
	// MoveNext advances the enumerator to the next element of the collection.
	MoveNext() bool

	// Current returns the element in the collection at the current position of the enumerator.
	Current() common.Elem

	// Reset sets the enumerator to its initial position, which is before the first element in the collection
	// (see https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.reset#remarks,
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator#remarks).
	Reset()
}
