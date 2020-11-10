package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 18 – ToLookup
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-18-tolookup/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup

// for 'withNilKeys' purpose see Addendum at
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-18-tolookup/
func toLookupPrim(en *Enumerable, ksel, esel func(common.Elem) common.Elem, keq common.Equality, withNilKeys bool) *Lookup {
	lk := newLookupEq(keq)
	for en.MoveNext() {
		c := en.Current()
		k := ksel(c)
		if !withNilKeys && k == nil {
			continue
		}
		if esel == nil {
			lk.add(k, c)
		} else {
			lk.add(k, esel(c))
		}
	}
	return lk
}

// ToLookupSelEq creates a Lookup from Enumerable.
// 'ksel' - key selector. 'esel' - element selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) ToLookupSelEq(ksel, esel func(common.Elem) common.Elem, keq common.Equality) (*Lookup, error) {
	if ksel == nil || esel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return toLookupPrim(en, ksel, esel, keq, true), nil
}

// ToLookupSelEqMust is like ToLookupSelEq but panics in case of error.
func (en *Enumerable) ToLookupSelEqMust(ksel, esel func(common.Elem) common.Elem, keq common.Equality) *Lookup {
	r, err := en.ToLookupSelEq(ksel, esel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSel creates a Lookup from Enumerable.
// 'ksel' - key selector. 'esel' - element selector.
// reflect.DeepEqual is used as keys Equality.
func (en *Enumerable) ToLookupSel(ksel, esel func(common.Elem) common.Elem) (*Lookup, error) {
	if ksel == nil || esel == nil {
		return nil, errors.NilSel
	}
	return toLookupPrim(en, ksel, esel, reflect.DeepEqual, true), nil
}

// ToLookupSelMust is like ToLookupSel but panics in case of error.
func (en *Enumerable) ToLookupSelMust(ksel, esel func(common.Elem) common.Elem) *Lookup {
	r, err := en.ToLookupSel(ksel, esel)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupEq creates a Lookup from Enumerable.
// 'ksel' - key selector. 'keq' - keys Equality.
// If 'keq' is nil reflect.DeepEqual is used.
func (en *Enumerable) ToLookupEq(ksel func(common.Elem) common.Elem, keq common.Equality) (*Lookup, error) {
	if ksel == nil {
		return nil, errors.NilSel
	}
	if keq == nil {
		keq = reflect.DeepEqual
	}
	return toLookupPrim(en, ksel, nil, keq, true), nil
}

// ToLookupEqMust is like ToLookupEq but panics in case of error.
func (en *Enumerable) ToLookupEqMust(ksel func(common.Elem) common.Elem, keq common.Equality) *Lookup {
	r, err := en.ToLookupEq(ksel, keq)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookup creates a Lookup from Enumerable.
// 'ksel' - key selector. reflect.DeepEqual is used as keys Equality.
func (en *Enumerable) ToLookup(ksel func(common.Elem) common.Elem) (*Lookup, error) {
	if ksel == nil {
		return nil, errors.NilSel
	}
	return toLookupPrim(en, ksel, nil, reflect.DeepEqual, true), nil
}

// ToLookupMust is like ToLookup but panics in case of error.
func (en *Enumerable) ToLookupMust(ksel func(common.Elem) common.Elem) *Lookup {
	r, err := en.ToLookup(ksel)
	if err != nil {
		panic(err)
	}
	return r
}
