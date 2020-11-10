package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToLookupTest.cs

func TestEnumerable_ToLookupMust(t *testing.T) {
	lk1 := newLookup()
	lk1.add(3, "abc")
	lk1.add(3, "def")
	lk1.add(1, "x")
	lk1.add(1, "y")
	lk1.add(3, "ghi")
	lk1.add(1, "z")
	lk1.add(2, "00")
	lk2 := newLookup()
	lk2.add("abc", "abc")
	lk2.add("def", "def")
	lk2.add("ABC", "ABC")
	type args struct {
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Lookup
	}{
		{name: "LookupWithNoComparerOrElementSelector",
			en:   NewElems("abc", "def", "x", "y", "ghi", "z", "00"),
			args: args{ksel: func(el common.Elem) common.Elem { return len(el.(string)) }},
			want: lk1},
		{name: "LookupWithNilComparerButNoElementSelector",
			en:   NewElems("abc", "def", "ABC"),
			args: args{ksel: common.Identity},
			want: lk2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ToLookupMust(tt.args.ksel); !got.Equal(tt.want) {
				t.Errorf("Enumerable.ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ToLookupSelMust(t *testing.T) {
	lk := newLookup()
	lk.add(3, "a")
	lk.add(3, "d")
	lk.add(1, "x")
	lk.add(1, "y")
	lk.add(3, "g")
	lk.add(1, "z")
	lk.add(2, "0")
	type args struct {
		ksel func(common.Elem) common.Elem
		esel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Lookup
	}{
		{name: "LookupWithElementSelectorButNoComparer",
			en: NewElems("abc", "def", "x", "y", "ghi", "z", "00"),
			args: args{
				ksel: func(el common.Elem) common.Elem { return len(el.(string)) },
				esel: func(el common.Elem) common.Elem { return string(el.(string)[0]) }},
			want: lk},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ToLookupSelMust(tt.args.ksel, tt.args.esel); !got.Equal(tt.want) {
				t.Errorf("Enumerable.ToLookupSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ToLookupEqMust(t *testing.T) {
	lk := newLookup()
	lk.add("abc", "abc")
	lk.add("def", "def")
	lk.add("abc", "ABC")
	type args struct {
		ksel func(common.Elem) common.Elem
		keq  common.Equality
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Lookup
	}{
		{name: "LookupWithComparerButNoElementSelector",
			en: NewElems("abc", "def", "ABC"),
			args: args{
				ksel: common.Identity,
				keq:  eqCaseInsensitive,
			},
			want: lk},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ToLookupEqMust(tt.args.ksel, tt.args.keq); !got.Equal(tt.want) {
				t.Errorf("Enumerable.ToLookupEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
