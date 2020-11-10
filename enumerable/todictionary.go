package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 25 – ToDictionary
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-todictionary/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

func toDictionaryPrim(en *Enumerable, ksel, vsel func(common.Elem) common.Elem) (Dictionary, error) {
	r := make(Dictionary)
	for en.MoveNext() {
		c := en.Current()
		k := ksel(c)
		if k == nil {
			return nil, errors.NilKey
		}
		if _, ok := r[k]; ok {
			return nil, errors.DupKeys
		}
		if vsel == nil {
			r[k] = c
		} else {
			r[k] = vsel(c)
		}
	}
	return r, nil
}

// ToDictionarySel creates a Dictionary from Enumerable
// according to specified key selector and value selector functions.
//
// Since Dictionary is implemented as map[common.Elem]common.Elem and since Go's map does not support
// equality comparer to determine equality of keys, hence LINQ's key comparer is not implemented.
// Similar to key comparer functionality may be achieved using appropriate key selector.
// Example of custom key selector that mimics case-insensitive equality comparer for string keys
// is presented in TestEnum_CustomEqualityComparer.
func (en *Enumerable) ToDictionarySel(ksel, vsel func(common.Elem) common.Elem) (Dictionary, error) {
	if ksel == nil || vsel == nil {
		return nil, errors.NilSel
	}
	return toDictionaryPrim(en, ksel, vsel)
}

// ToDictionarySelMust is like ToDictionarySel but panics in case of error.
func (en *Enumerable) ToDictionarySelMust(ksel, vsel func(common.Elem) common.Elem) Dictionary {
	r, err := en.ToDictionarySel(ksel, vsel)
	if err != nil {
		panic(err)
	}
	return r
}

// ToDictionary creates a Dictionary from an Enumerable according to a specified key selector function.
// (See help for ToDictionarySelEr.)
func (en *Enumerable) ToDictionary(ksel func(common.Elem) common.Elem) (Dictionary, error) {
	if ksel == nil {
		return nil, errors.NilSel
	}
	return toDictionaryPrim(en, ksel, nil)
}

// ToDictionaryMust is like ToDictionary but panics in case of error.
func (en *Enumerable) ToDictionaryMust(ksel func(common.Elem) common.Elem) Dictionary {
	r, err := en.ToDictionary(ksel)
	if err != nil {
		panic(err)
	}
	return r
}
