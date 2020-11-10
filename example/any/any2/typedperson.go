package main

// generated by github.com/solsw/golinq/typed/typedgen utility

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/errors"
)

// Sliceperson is an alias for []person.
type Sliceperson = []person

// ElemToperson converts Elem to person.
// If 'el' is nil, person's zero value is returned.
func ElemToperson(el common.Elem) (person, error) {
	var r0 person
	if el == nil {
		return r0, nil
	}
	r, ok := el.(person)
	if !ok {
		return r0, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(r))
	}
	return r, nil
}

// SlicepersonToSlice converts Sliceperson to Slice.
func SlicepersonToSlice(ts Sliceperson) common.Slice {
	if ts == nil {
		return nil
	}
	r := make(common.Slice, 0, len(ts))
	for _, t := range ts {
		r = append(r, t)
	}
	return r
}

// SlicepersonToEnumerable converts Sliceperson to Enumerable.
func SlicepersonToEnumerable(ts Sliceperson) *enumerable.Enumerable {
	return enumerable.NewElems(SlicepersonToSlice(ts)...)
}

// SliceToSliceperson converts Slice to Sliceperson.
func SliceToSliceperson(sl common.Slice) (Sliceperson, error) {
	if sl == nil {
		return nil, nil
	}
	r := make(Sliceperson, 0, len(sl))
	for _, el := range sl {
		t, ok := el.(person)
		if !ok {
			return nil, errors.WrongType(reflect.TypeOf(el), reflect.TypeOf(t))
		}
		r = append(r, t)
	}
	return r, nil
}

// EnumerableToSliceperson converts Enumerable to Sliceperson.
func EnumerableToSliceperson(en *enumerable.Enumerable) (Sliceperson, error) {
	return SliceToSliceperson(en.Slice())
}

// Eqperson is an Equality for person.
var Eqperson = func(e1, e2 common.Elem) bool {
	// use the appropiate one or implement your own
	return reflect.DeepEqual(e1, e2)
	// return reflect.DeepEqual(e1.(person), e2.(person))
	// return e1.(person) == e2.(person)
}

// EqNilperson is an Equality for person, respecting nil parameters.
var EqNilperson = func(e1, e2 common.Elem) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	return Eqperson(e1, e2)
}
