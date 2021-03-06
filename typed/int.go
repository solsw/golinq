package typed

// generated by github.com/solsw/golinq/typed/typedgen utility

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/errors"
)

// Sliceint is an alias for []int.
type Sliceint = []int

// ElemToint converts Elem to int.
// If 'el' is nil, int's zero value is returned.
func ElemToint(el common.Elem) (int, error) {
	var r0 int
	if el == nil {
		return r0, nil
	}
	r, ok := el.(int)
	if !ok {
		return r0, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(r))
	}
	return r, nil
}

// SliceintToSlice converts Sliceint to Slice.
func SliceintToSlice(ts Sliceint) common.Slice {
	if ts == nil {
		return nil
	}
	r := make(common.Slice, 0, len(ts))
	for _, t := range ts {
		r = append(r, t)
	}
	return r
}

// SliceintToEnumerable converts Sliceint to Enumerable.
func SliceintToEnumerable(ts Sliceint) *enumerable.Enumerable {
	return enumerable.NewElems(SliceintToSlice(ts)...)
}

// SliceToSliceint converts Slice to Sliceint.
func SliceToSliceint(sl common.Slice) (Sliceint, error) {
	if sl == nil {
		return nil, nil
	}
	r := make(Sliceint, 0, len(sl))
	for _, el := range sl {
		t, ok := el.(int)
		if !ok {
			return nil, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(t))
		}
		r = append(r, t)
	}
	return r, nil
}

// EnumerableToSliceint converts Enumerable to Sliceint.
func EnumerableToSliceint(en *enumerable.Enumerable) (Sliceint, error) {
	return SliceToSliceint(en.Slice())
}

// Eqint is an Equality for int.
var Eqint = func(e1, e2 common.Elem) bool {
	// use the appropiate one or implement your own
	return reflect.DeepEqual(e1, e2)
	// return reflect.DeepEqual(e1.(int), e2.(int))
	// return e1.(int) == e2.(int)
}

// EqNilint is an Equality for int, respecting nil parameters.
var EqNilint = func(e1, e2 common.Elem) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	return Eqint(e1, e2)
}

// Lessint is a Less for int.
var Lessint = func(e1, e2 common.Elem) bool {
	// usually requires special implementation
	return e1.(int) < e2.(int)
}

// LessNilint is a Less for int, respecting nil parameters.
var LessNilint = func(e1, e2 common.Elem) bool {
	// must go first
	if e2 == nil {
		return false
	}
	if e1 == nil {
		return true
	}
	return Lessint(e1, e2)
}

// Cmpint is a Comparison for int.
var Cmpint = common.LessToComparison(Lessint)

// CmpNilint is a Comparison for int, respecting nil parameters.
var CmpNilint = common.LessToComparison(LessNilint)
