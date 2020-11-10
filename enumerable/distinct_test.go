package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DistinctTest.cs

var (
	testStr1 = "test"
	testStr2 = "test"
)

func TestEnumerable_Distinct(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want *Enumerable
	}{
		{name: "1",
			en:   NewElems("A", "a", "b", "c", "b"),
			want: NewElems("A", "a", "b", "c")},
		{name: "2",
			en:   NewElems("b", "a", "d", "a"),
			want: NewElems("b", "a", "d")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Distinct(); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_DistinctEq(t *testing.T) {
	type args struct {
		eq common.Equality
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			en:   NewElems("xyz", testStr1, "XYZ", testStr2, "def"),
			args: args{eq: eqCaseInsensitive},
			want: NewElems("xyz", testStr1, "def")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.DistinctEq(tt.args.eq); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.DistinctEq() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_DistinctCmpMust(t *testing.T) {
	type args struct {
		cmp common.Comparison
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			en:   NewElems("xyz", testStr1, "XYZ", testStr2, "def"),
			args: args{cmp: cmpCaseInsensitive},
			want: NewElems("xyz", testStr1, "def")},
		{name: "EmptyEnumerable",
			en:   enEmpty,
			args: args{cmp: cmpInt},
			want: enEmpty},
		{name: "EmptyEnumerable2",
			en:   enEmptyEnmr,
			args: args{cmp: cmpInt},
			want: enEmpty},
		{name: "EmptyEnumerable3",
			en:   enEmptySlice,
			args: args{cmp: cmpInt},
			want: enEmpty},
		{name: "1",
			en:   enInt4(),
			args: args{cmp: cmpInt},
			want: enInt4()},
		{name: "2",
			en:   enInt4().Concat(enInt4()),
			args: args{cmp: cmpInt},
			want: enInt4()},
		{name: "3",
			en:   NewElems("A", "a", "b", "c", "b"),
			args: args{cmp: cmpCaseInsensitive},
			want: NewElems("A", "b", "c")},
		{name: "4",
			en:   NewElems("b", "a", "d", "a"),
			args: args{cmp: cmpCaseInsensitive},
			want: NewElems("b", "a", "d")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.DistinctCmpMust(tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.DistinctCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_DistinctLsMust(t *testing.T) {
	type args struct {
		ls common.Less
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "DistinctCmpWithCaseInsensitiveComparer",
			en:   NewElems("a", "b", "b"),
			args: args{ls: lsCaseSensitive},
			want: NewElems("a", "b")},
		{name: "DistinctCmpWithCaseInsensitiveComparer2",
			en:   NewElems("a", "b", "B"),
			args: args{ls: lsCaseInsensitive},
			want: NewElems("a", "b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.DistinctLsMust(tt.args.ls); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.DistinctLsMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_DistinctEq_Reset(t *testing.T) {
	en := NewElems("xyz", testStr1, "XYZ", testStr2, "def").DistinctEq(eqCaseInsensitive)
	got1 := NewElems(en.Slice()...)
	en.Reset()
	got2 := NewElems(en.Slice()...)
	if !got1.SequenceEqual(got2) {
		got1.Reset()
		got2.Reset()
		t.Errorf("Reset error: '%v' != '%v'", got1, got2)
	}
}
