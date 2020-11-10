package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 31 – ElementAt / ElementAtOrDefault
// https://codeblog.jonskeet.uk/2011/01/11/reimplementing-linq-to-objects-part-31-elementat-elementatordefault/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault

// ElementAt returns the element at a specified index in Enumerable.
func (en *Enumerable) ElementAt(idx int) (common.Elem, error) {
	if idx < 0 {
		return nil, errors.IdxRange
	}
	if c, ok := (en.enmr).(common.Counter); ok {
		if idx >= c.Count() {
			return nil, errors.IdxRange
		}
		if i, ok := (en.enmr).(common.Itemer); ok {
			return i.Item(idx), nil
		}
	}
	i := 0
	for en.MoveNext() {
		if i == idx {
			return en.Current(), nil
		}
		i++
	}
	return nil, errors.IdxRange
}

// ElementAtMust is like ElementAt but panics in case of error.
func (en *Enumerable) ElementAtMust(idx int) common.Elem {
	r, err := en.ElementAt(idx)
	if err != nil {
		panic(err)
	}
	return r
}

// ElementAtOrDefault returns the element at a specified index in Enumerable or nil if the index is out of range.
func (en *Enumerable) ElementAtOrDefault(idx int) common.Elem {
	r, err := en.ElementAt(idx)
	if err != nil {
		return nil
	}
	return r
}
