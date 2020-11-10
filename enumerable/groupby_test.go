package enumerable

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupByTest.cs

func Test_GroupByMust(t *testing.T) {
	en := NewElems("abc", "hello", "def", "there", "four")
	grs := en.GroupByMust(func(el common.Elem) common.Elem { return len(el.(string)) }).Slice()
	if len(grs) != 3 {
		t.Errorf("len(Enumerable.GroupByMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].(Grouping).values)
	if lg0 != 2 {
		t.Errorf("len(Enumerable.GroupByMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0].(Grouping)
	if gr0.key != 3 {
		t.Errorf("Enumerable.GroupByMust[0].Key = '%v', want '%v'", gr0.key, 3)
	}
	got0 := NewElems(gr0.values...)
	want0 := NewElems("abc", "def")
	if !got0.SequenceEqual(want0) {
		got0.Reset()
		want0.Reset()
		t.Errorf("Enumerable.GroupByMust[0].values = '%v', want '%v'", got0, want0)
	}

	gr1 := grs[1].(Grouping)
	if gr1.key != 5 {
		t.Errorf("Enumerable.GroupByMust[1].Key = '%v', want '%v'", gr1.key, 5)
	}
	got1 := NewElems(gr1.values...)
	want1 := NewElems("hello", "there")
	if !got1.SequenceEqual(want1) {
		got1.Reset()
		want1.Reset()
		t.Errorf("Enumerable.GroupByMust[1].values = '%v', want '%v'", got1, want1)
	}

	gr2 := grs[2].(Grouping)
	if gr2.key != 4 {
		t.Errorf("Enumerable.GroupByMust[2].Key = '%v', want '%v'", gr2.key, 4)
	}
	got2 := NewElems(gr2.values...)
	want2 := NewElems("four")
	if !got2.SequenceEqual(want2) {
		got2.Reset()
		want2.Reset()
		t.Errorf("Enumerable.GroupByMust[2].values = '%v', want '%v'", got2, want2)
	}
}

func Test_GroupBySelMust(t *testing.T) {
	en := NewElems("abc", "hello", "def", "there", "four")
	grs := en.GroupBySelMust(
		func(el common.Elem) common.Elem { return len(el.(string)) },
		func(el common.Elem) common.Elem { return []rune(el.(string))[0] }).Slice()
	if len(grs) != 3 {
		t.Errorf("len(Enumerable.GroupBySelMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].(Grouping).values)
	if lg0 != 2 {
		t.Errorf("len(Enumerable.GroupBySelMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0].(Grouping)
	if gr0.key != 3 {
		t.Errorf("Enumerable.GroupBySelMust[0].Key = '%v', want '%v'", gr0.key, 3)
	}
	got0 := NewElems(gr0.values...)
	want0 := NewElems('a', 'd')
	if !got0.SequenceEqual(want0) {
		got0.Reset()
		want0.Reset()
		t.Errorf("Enumerable.GroupBySelMust[0].values = '%v', want '%v'", got0, want0)
	}

	gr1 := grs[1].(Grouping)
	if gr1.key != 5 {
		t.Errorf("Enumerable.GroupBySelMust[1].Key = '%v', want '%v'", gr1, 3)
	}
	got1 := NewElems(gr1.values...)
	want1 := NewElems('h', 't')
	if !got1.SequenceEqual(want1) {
		got1.Reset()
		want1.Reset()
		t.Errorf("Enumerable.GroupBySelMust[1].values = '%v', want '%v'", got1, want1)
	}

	gr2 := grs[2].(Grouping)
	if gr2.key != 4 {
		t.Errorf("Enumerable.GroupBySelMust[2].Key = '%v', want '%v'", gr2, 3)
	}
	got2 := NewElems(gr2.values...)
	want2 := NewElems('f')
	if !got2.SequenceEqual(want2) {
		got2.Reset()
		want2.Reset()
		t.Errorf("Enumerable.GroupBySelMust[2].values = '%v', want '%v'", got2, want2)
	}
}

func Test_GroupByResMust(t *testing.T) {
	en := NewElems("abc", "hello", "def", "there", "four")
	grs := en.GroupByResMust(
		func(el common.Elem) common.Elem { return len(el.(string)) },
		func(el common.Elem, en *Enumerable) common.Elem {
			return fmt.Sprintf("%v:%v", el.(int), strings.Join(enToStrings(en), ";"))
		}).Slice()
	got := NewElems(grs...)
	want := NewElems("3:abc;def", "5:hello;there", "4:four")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("Enumerable.GroupByResMust = '%v', want '%v'", got, want)
	}
}

func Test_GroupBySelResMust(t *testing.T) {
	en := NewElems("abc", "hello", "def", "there", "four")
	grs := en.GroupBySelResMust(
		func(el common.Elem) common.Elem { return len(el.(string)) },
		func(el common.Elem) common.Elem { return []rune(el.(string))[0] },
		func(el common.Elem, en *Enumerable) common.Elem {
			vv := func() []string {
				r := make([]string, 0)
				for en.MoveNext() {
					r = append(r, string(en.Current().(rune)))
				}
				return r
			}()
			return fmt.Sprintf("%v:%v", el.(int), strings.Join(vv, ";"))
		}).Slice()
	got := NewElems(grs...)
	want := NewElems("3:a;d", "5:h;t", "4:f")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("Enumerable.GroupBySelResMust = '%v', want '%v'", got, want)
	}
}

func Test_GroupByResMust_NilKeys(t *testing.T) {
	en := NewElems("first", "nil", "nothing", "second")
	grs := en.GroupByResMust(
		func(el common.Elem) common.Elem {
			s := el.(string)
			if strings.HasPrefix(s, "n") {
				return nil
			}
			return s
		},
		func(el common.Elem, en *Enumerable) common.Elem {
			var k string
			if el == nil {
				k = ""
			} else {
				k = el.(string)
			}
			return fmt.Sprintf("%v:%v", k, strings.Join(enToStrings(en), ";"))
		}).Slice()
	got := NewElems(grs...)
	want := NewElems("first:first", ":nil;nothing", "second:second")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("Enumerable.GroupByResMust_NilKeys = '%v', want '%v'", got, want)
	}
}
