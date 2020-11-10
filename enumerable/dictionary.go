package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// KeyValue represents key, value pair of Dictionary.
type KeyValue struct {
	key   common.Elem
	value common.Elem
}

// Key returns the key of the KeyValue.
func (kv *KeyValue) Key() common.Elem {
	return kv.key
}

// Value returns the value of the KeyValue.
func (kv *KeyValue) Value() common.Elem {
	return kv.value
}

// Dictionary represents map[common.Elem]common.Elem.
type Dictionary map[common.Elem]common.Elem

// Enumerable converts Dictionary to Enumerable of KeyValues.
func (d Dictionary) Enumerable() *Enumerable {
	r := make(common.Slice, 0, len(d))
	for k, v := range d {
		r = append(r, KeyValue{k, v})
	}
	return NewElems(r...)
}

// Dictionary converts Enumerable of KeyValues to Dictionary.
func (en *Enumerable) Dictionary() (Dictionary, error) {
	r := make(Dictionary)
	for en.MoveNext() {
		kv, ok := en.Current().(KeyValue)
		if !ok {
			return nil, errors.WrongType(reflect.TypeOf(en.Current()), reflect.TypeOf(kv))
		}
		r[kv.key] = kv.value
	}
	return r, nil
}
