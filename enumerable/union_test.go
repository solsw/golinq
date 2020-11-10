package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func TestEnumerable_Union(t *testing.T) {
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "UnionWithTwoEmptySequences",
			en:   enEmpty,
			args: args{en2: enEmpty},
			want: enEmpty},
		{name: "FirstEmpty",
			en:   enEmpty,
			args: args{en2: enInt4()},
			want: enInt4()},
		{name: "SecondEmpty",
			en:   enInt4(),
			args: args{en2: enEmpty},
			want: enInt4()},
		{name: "UnionWithoutComparer",
			en:   NewElems("a", "b", "B", "c", "b"),
			args: args{en2: NewElems("d", "e", "d", "a")},
			want: NewElems("a", "b", "B", "c", "d", "e")},
		{name: "UnionWithoutComparer2",
			en:   NewElems("a", "b"),
			args: args{en2: NewElems("b", "a")},
			want: NewElems("a", "b")},
		{name: "UnionWithEmptyFirstSequence",
			en:   enEmpty,
			args: args{en2: NewElems("d", "e", "d", "a")},
			want: NewElems("d", "e", "a")},
		{name: "UnionWithEmptySecondSequence",
			en:   NewElems("a", "b", "B", "c", "b"),
			args: args{en2: enEmpty},
			want: NewElems("a", "b", "B", "c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Union(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Union() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_UnionSelf(t *testing.T) {
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
			en:   e2.Take(1),
			args: args{en2: e2.Skip(3)},
			want: NewElems(1, 4)},
		{name: "SameEnumerable3",
			en:   e3.Skip(2),
			args: args{en2: e3},
			want: NewElems(3, 4, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.UnionSelf(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.UnionSelf() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_UnionEq(t *testing.T) {
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
		{name: "UnionWithIntEquality",
			en: NewElems(1, 2),
			args: args{
				en2: NewElems(2, 3),
				eq:  eqInt},
			want: NewElems(1, 2, 3)},
		{name: "UnionWithCaseInsensitiveComparerEq",
			en: NewElems("a", "b", "B", "c", "b"),
			args: args{
				en2: NewElems("d", "e", "d", "a"),
				eq:  eqCaseInsensitive},
			want: NewElems("a", "b", "c", "d", "e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.UnionEq(tt.args.en2, tt.args.eq); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.UnionEq() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_UnionCmpMust(t *testing.T) {
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
		{name: "UnionWithIntComparer1",
			en: NewElems(1, 2, 2),
			args: args{
				en2: enEmpty,
				cmp: cmpInt},
			want: NewElems(1, 2)},
		{name: "UnionWithIntComparer2",
			en: NewElems(1, 2),
			args: args{
				en2: NewElems(2, 3),
				cmp: cmpInt},
			want: NewElems(1, 2, 3)},
		{name: "UnionWithCaseInsensitiveComparerCmp",
			en: NewElems("a", "b", "B", "c", "b"),
			args: args{
				en2: NewElems("d", "e", "d", "a"),
				cmp: cmpCaseInsensitive},
			want: NewElems("a", "b", "c", "d", "e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.UnionCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.UnionCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_UnionSelfCmpMust(t *testing.T) {
	e1 := enInt4()
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
			en: e1,
			args: args{
				en2: e1,
				cmp: cmpInt},
			want: enInt4()},
		{name: "SameEnumerable2",
			en: e2.Skip(2),
			args: args{
				en2: e2.Take(1),
				cmp: cmpInt},
			want: NewElems(3, 4, 1)},
		{name: "SameEnumerable3",
			en: e3.Skip(2),
			args: args{
				en2: e3,
				cmp: cmpInt},
			want: NewElems(3, 4, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.UnionSelfCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.UnionSelfCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_UnionSelfLsMust(t *testing.T) {
	e1 := enInt4()
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
			en: e1,
			args: args{
				en2: e1,
				ls:  lsInt},
			want: enInt4()},
		{name: "SameEnumerable2",
			en: e2.Take(2),
			args: args{
				en2: e2.Skip(3),
				ls:  lsInt},
			want: NewElems(1, 2, 4)},
		{name: "SameEnumerable3",
			en: e3.Skip(2),
			args: args{
				en2: e3,
				ls:  lsInt},
			want: NewElems(3, 4, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.UnionSelfLsMust(tt.args.en2, tt.args.ls); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.UnionSelfLsMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
