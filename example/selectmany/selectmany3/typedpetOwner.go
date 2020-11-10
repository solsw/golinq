package main

// generated by github.com/solsw/golinq/typed/typedgen utility

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/errors"
)

// SlicepetOwner is an alias for []petOwner.
type SlicepetOwner = []petOwner

// ElemTopetOwner converts Elem to petOwner.
// If 'el' is nil, petOwner's zero value is returned.
func ElemTopetOwner(el common.Elem) (petOwner, error) {
	var r0 petOwner
	if el == nil {
		return r0, nil
	}
	r, ok := el.(petOwner)
	if !ok {
		return r0, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(r))
	}
	return r, nil
}

// SlicepetOwnerToSlice converts SlicepetOwner to Slice.
func SlicepetOwnerToSlice(ts SlicepetOwner) common.Slice {
	if ts == nil {
		return nil
	}
	r := make(common.Slice, 0, len(ts))
	for _, t := range ts {
		r = append(r, t)
	}
	return r
}

// SlicepetOwnerToEnumerable converts SlicepetOwner to Enumerable.
func SlicepetOwnerToEnumerable(ts SlicepetOwner) *enumerable.Enumerable {
	return enumerable.NewElems(SlicepetOwnerToSlice(ts)...)
}

// SliceToSlicepetOwner converts Slice to SlicepetOwner.
func SliceToSlicepetOwner(sl common.Slice) (SlicepetOwner, error) {
	if sl == nil {
		return nil, nil
	}
	r := make(SlicepetOwner, 0, len(sl))
	for _, el := range sl {
		t, ok := el.(petOwner)
		if !ok {
			return nil, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(t))
		}
		r = append(r, t)
	}
	return r, nil
}

// EnumerableToSlicepetOwner converts Enumerable to SlicepetOwner.
func EnumerableToSlicepetOwner(en *enumerable.Enumerable) (SlicepetOwner, error) {
	return SliceToSlicepetOwner(en.Slice())
}

// EqpetOwner is an Equality for petOwner.
var EqpetOwner = func(e1, e2 common.Elem) bool {
	// use the appropiate one or implement your own
	return reflect.DeepEqual(e1, e2)
	// return reflect.DeepEqual(e1.(petOwner), e2.(petOwner))
	// return e1.(petOwner) == e2.(petOwner)
}

// EqNilpetOwner is an Equality for petOwner, respecting nil parameters.
var EqNilpetOwner = func(e1, e2 common.Elem) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	return EqpetOwner(e1, e2)
}
