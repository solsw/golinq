package enumerable

import (
	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// Reimplementing LINQ to Objects: Part 13 – Aggregate
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-13-aggregate/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

// Aggregate applies an accumulator function over Enumerable.
func (en *Enumerable) Aggregate(acc func(common.Elem, common.Elem) common.Elem) (common.Elem, error) {
	if acc == nil {
		return nil, errors.NilAcc
	}
	if !en.MoveNext() {
		return nil, errors.EmptyEnum
	}
	r := en.Current()
	for en.MoveNext() {
		r = acc(r, en.Current())
	}
	return r, nil
}

// AggregateMust is like Aggregate but panics in case of error.
func (en *Enumerable) AggregateMust(acc func(common.Elem, common.Elem) common.Elem) common.Elem {
	r, err := en.Aggregate(acc)
	if err != nil {
		panic(err)
	}
	return r
}

// AggregateSeed applies an accumulator function over Enumerable.
// The specified seed value is used as the initial accumulator value.
func (en *Enumerable) AggregateSeed(seed interface{}, acc func(interface{}, common.Elem) interface{}) (interface{}, error) {
	if acc == nil {
		return nil, errors.NilAcc
	}
	r := seed
	for en.MoveNext() {
		r = acc(r, en.Current())
	}
	return r, nil
}

// AggregateSeedMust is like AggregateSeed but panics in case of error.
func (en *Enumerable) AggregateSeedMust(seed interface{}, acc func(interface{}, common.Elem) interface{}) interface{} {
	r, err := en.AggregateSeed(seed, acc)
	if err != nil {
		panic(err)
	}
	return r
}

// AggregateSeedSel applies an accumulator function over Enumerable.
// The specified seed value is used as the initial accumulator value,
// and the specified selector is used to select the result value.
func (en *Enumerable) AggregateSeedSel(seed interface{}, acc func(interface{}, common.Elem) interface{},
	sel func(interface{}) interface{}) (interface{}, error) {
	if acc == nil {
		return nil, errors.NilAcc
	}
	if sel == nil {
		return nil, errors.NilSel
	}
	r := seed
	for en.MoveNext() {
		r = acc(r, en.Current())
	}
	return sel(r), nil
}

// AggregateSeedSelMust is like AggregateSeedSel but panics in case of error.
func (en *Enumerable) AggregateSeedSelMust(seed interface{}, acc func(interface{}, common.Elem) interface{}, sel func(interface{}) interface{}) interface{} {
	r, err := en.AggregateSeedSel(seed, acc, sel)
	if err != nil {
		panic(err)
	}
	return r
}
