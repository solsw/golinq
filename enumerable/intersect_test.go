package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/IntersectTest.cs

func TestEnumerable_Intersect(t *testing.T) {
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "1",
			en:   NewElems(1, 2),
			args: args{en2: NewElems(2, 3)},
			want: NewElems(2)},
		{name: "NoComparerSpecified",
			en:   NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a")},
			want: NewElems("a", "b")},
		{name: "IntWithoutComparer",
			en:   NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10)},
			want: NewElems(4, 5, 6, 7, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Intersect(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Intersect() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_IntersectSelf(t *testing.T) {
	e1 := enInt4()
	e2 := enInt4()
	e3 := enInt4()
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SameEnumerable1",
			en:   e1,
			args: args{en2: e1},
			want: enInt4()},
		{name: "SameEnumerable2",
			en:   e2,
			args: args{en2: e2.Skip(1)},
			want: NewElems(2, 3, 4)},
		{name: "SameEnumerable3",
			en:   e3.Skip(3),
			args: args{en2: e3},
			want: NewElems(4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.IntersectSelf(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.IntersectSelf() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_IntersectEq(t *testing.T) {
	type args struct {
		en2 *Enumerable
		eq  common.Equality
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "CaseInsensitiveComparerSpecified",
			en: NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a"),
				eq: eqCaseInsensitive},
			want: NewElems("A", "b")},
		{name: "IntComparerSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				eq: eqInt},
			want: NewElems(4, 5, 6, 7, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.IntersectEq(tt.args.en2, tt.args.eq); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.IntersectEq() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_IntersectCmpMust(t *testing.T) {
	type args struct {
		en2 *Enumerable
		cmp common.Comparison
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "CaseInsensitiveComparerSpecified",
			en: NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a"),
				cmp: cmpCaseInsensitive},
			want: NewElems("A", "b")},
		{name: "IntComparerSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp: cmpInt},
			want: NewElems(4, 5, 6, 7, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.IntersectCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.IntersectCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_IntersectSelfCmpMust(t *testing.T) {
	e1 := NewElems(4, 3, 2, 1)
	e2 := enInt4()
	e3 := enInt4()
	type args struct {
		en2 *Enumerable
		cmp common.Comparison
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SameEnumerable1",
			en:   e1,
			args: args{en2: e1, cmp: cmpInt},
			want: NewElems(4, 3, 2, 1)},
		{name: "SameEnumerable2",
			en:   e2,
			args: args{en2: e2.Skip(1), cmp: cmpInt},
			want: NewElems(2, 3, 4)},
		{name: "SameEnumerable3",
			en:   e3.Skip(3),
			args: args{en2: e3, cmp: cmpInt},
			want: NewElems(4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.IntersectSelfCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.IntersectSelfCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_IntersectSelfLsMust(t *testing.T) {
	e1 := NewElems(3, 2, 4, 1)
	e2 := enInt4()
	e3 := enInt4()
	type args struct {
		en2 *Enumerable
		ls  common.Less
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SameEnumerable1",
			en:   e1,
			args: args{en2: e1, ls: lsInt},
			want: NewElems(3, 2, 4, 1)},
		{name: "SameEnumerable2",
			en:   e2,
			args: args{en2: e2.Skip(1), ls: lsInt},
			want: NewElems(2, 3, 4)},
		{name: "SameEnumerable3",
			en:   e3.Skip(3),
			args: args{en2: e3, ls: lsInt},
			want: NewElems(4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.IntersectSelfLsMust(tt.args.en2, tt.args.ls); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.IntersectSelfLsMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
