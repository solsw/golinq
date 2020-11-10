package enumerable

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/gohelpers/stringhelper"
	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupJoinTest.cs

func Test_GroupJoinMust_SimpleGroupJoin(t *testing.T) {
	outer := NewElems("first", "second", "third")
	inner := NewElems("essence", "offer", "eating", "psalm")
	got := outer.GroupJoinMust(inner,
		func(oel common.Elem) common.Elem { return []rune(oel.(string))[0] },
		func(iel common.Elem) common.Elem { return []rune(iel.(string))[1] },
		func(oel common.Elem, iels *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", oel, strings.Join(enToStrings(iels), ";"))
		})
	want := NewElems("first:offer", "second:essence;psalm", "third:")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinMust_SimpleGroupJoin = '%v', want '%v'", got, want)
	}
}

func Test_GroupJoinSelfMust_SameEnumerable(t *testing.T) {
	outer := NewElems("fs", "sf", "ff", "ss")
	inner := outer
	got := outer.GroupJoinSelfMust(inner,
		func(oel common.Elem) common.Elem { return []rune(oel.(string))[0] },
		func(iel common.Elem) common.Elem { return []rune(iel.(string))[1] },
		func(oel common.Elem, iels *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", oel, strings.Join(enToStrings(iels), ";"))
		}).Slice()
	want := common.Slice{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinSelfMust_SameEnumerable = %v, want %v", got, want)
	}
}

func Test_GroupJoinEqMust_CustomComparer(t *testing.T) {
	outer := NewElems("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewElems("000abc", "111gHi", "222333", "333AbC")
	got := outer.GroupJoinEqMust(inner,
		func(oel common.Elem) common.Elem {
			r, _ := stringhelper.SubstrBeg(oel.(string), 3)
			return r
		},
		func(iel common.Elem) common.Elem {
			r, _ := stringhelper.SubstrEnd(iel.(string), 3)
			return r
		},
		func(oel common.Elem, iels *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", oel, strings.Join(enToStrings(iels), ";"))
		},
		eqCaseInsensitive)
	want := NewElems("ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinEqMust_CustomComparer = '%v', want '%v'", got, want)
	}
}

func Test_GroupJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := NewElems(5, 3, 7, 4)
	inner := NewElems("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got := outer.GroupJoinMust(inner, common.Identity,
		func(iel common.Elem) common.Elem { return len(iel.(string)) },
		func(oel common.Elem, iels *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", oel, strings.Join(enToStrings(iels), ";"))
		},
	)
	want := NewElems("5:tiger", "3:bee;cat;dog", "7:giraffe", "4:")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinMust_DifferentSourceTypes = '%v', want '%v'", got, want)
	}
}

func Test_GroupJoinMust_NilKeys(t *testing.T) {
	outer := NewElems("first", nil, "second")
	inner := NewElems("first", "nil", "nothing")
	got := outer.GroupJoinMust(inner, common.Identity,
		func(iel common.Elem) common.Elem {
			if strings.HasPrefix(iel.(string), "n") {
				return nil
			}
			return iel
		},
		func(oel common.Elem, iels *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", oel, strings.Join(enToStrings(iels), ";"))
		},
	)
	want := NewElems("first:first", "<nil>:", "second:")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinMust_NilKeys = '%v', want '%v'", got, want)
	}
}
