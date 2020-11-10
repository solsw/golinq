package enumerable

import (
	"reflect"

	"github.com/solsw/golinq/common"
)

// Reimplementing LINQ to Objects: Part 34 – SequenceEqual
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-34-sequenceequal/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

// SequenceEqual determines whether two Enumerables are equal using reflect.DeepEqual as Equality.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use SequenceEqualSelf instead.
func (en *Enumerable) SequenceEqual(en2 *Enumerable) bool {
	return en.SequenceEqualEq(en2, reflect.DeepEqual)
}

// SequenceEqualSelf determines whether two Enumerables are equal using reflect.DeepEqual as Equality.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) SequenceEqualSelf(en2 *Enumerable) bool {
	return en.SequenceEqualSelfEq(en2, reflect.DeepEqual)
}

// SequenceEqualEq determines whether two Enumerables are equal by comparing their elements by using a specified Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// Pointer receiver is required to check if the Enumerable is compared with itself.
// 'en' and 'en2' must NOT be based on the same Enumerable, otherwise use SequenceEqualSelfEq instead.
func (en *Enumerable) SequenceEqualEq(en2 *Enumerable, eq common.Equality) bool {
	if en == en2 {
		return true
	}
	if en.enmr == nil && en2.enmr == nil {
		return true
	}
	if eq == nil {
		eq = reflect.DeepEqual
	}
	for en.MoveNext() {
		if !en2.MoveNext() {
			return false
		}
		if !eq(en.Current(), en2.Current()) {
			return false
		}
	}
	if en2.MoveNext() {
		return false
	}
	return true
}

// SequenceEqualSelfEq determines whether two Enumerables are equal by comparing their elements by using a specified Equality.
// If 'eq' is nil reflect.DeepEqual is used.
// Pointer receiver is required to check if the Enumerable is compared with itself.
// 'en' and 'en2' may be based on the same Enumerable.
func (en *Enumerable) SequenceEqualSelfEq(en2 *Enumerable, eq common.Equality) bool {
	sl2 := en2.Slice()
	en.Reset()
	return en.SequenceEqualEq(NewElems(sl2...), eq)
}
