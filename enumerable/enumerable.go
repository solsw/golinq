package enumerable

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerator"
)

// Enumerable mimics .NET's Enumerable Class
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable)
// for querying a sequence of common.Elem and implements the enumerator.Enumerator interface.
type Enumerable struct {
	enmr enumerator.Enumerator
}

// MoveNext implements the enumerator.Enumerator.MoveNext method.
func (en *Enumerable) MoveNext() bool {
	if en.enmr == nil {
		return false
	}
	return en.enmr.MoveNext()
}

// Current implements the enumerator.Enumerator.Current method.
func (en *Enumerable) Current() common.Elem {
	if en.enmr == nil {
		return nil
	}
	return en.enmr.Current()
}

// Reset implements the enumerator.Enumerator.Reset method.
func (en *Enumerable) Reset() {
	if en.enmr == nil {
		return
	}
	en.enmr.Reset()
}

// NewEmpty creates a new empty Enumerable.
func NewEmpty() *Enumerable {
	return &Enumerable{}
}

// NewElems creates a new Enumerable based on enumerator.OnSlice with 'ee' contents.
func NewElems(ee ...common.Elem) *Enumerable {
	return &Enumerable{enumerator.NewOnSlice(ee...)}
}

// Slice converts Enumerable to common.Slice.
func (en *Enumerable) Slice() common.Slice {
	var r common.Slice
	for en.MoveNext() {
		r = append(r, en.Current())
	}
	return r
}

// String implements the fmt.Stringer interface.
func (en *Enumerable) String() string {
	const sep string = " "
	var b strings.Builder
	for en.MoveNext() {
		if b.Len() > 0 {
			b.WriteString(sep)
		}
		b.WriteString(fmt.Sprint(en.Current()))
	}
	return b.String()
}

// CloneEmpty creates a new Enumerable with empty Enumerator of the same type as in 'en'.
func (en *Enumerable) CloneEmpty() *Enumerable {
	// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
	t := reflect.TypeOf(en.enmr)
	var v reflect.Value
	if t.Kind() == reflect.Ptr {
		v = reflect.New(t.Elem())
	} else {
		v = reflect.Zero(t)
	}
	i := v.Interface()
	e, ok := i.(enumerator.Enumerator)
	if !ok {
		panic("e, ok := i.(enumerator.Enumerator)")
	}
	return &Enumerable{e}
}
