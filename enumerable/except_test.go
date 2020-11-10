package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func TestEnumerable_Except(t *testing.T) {
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "NoComparerSpecified",
			en:   NewElems("A", "a", "b", "c", "b", "c"),
			args: args{en2: NewElems("b", "a", "d", "a")},
			want: NewElems("A", "c")},
		{name: "IntWithoutComparer",
			en:   NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10)},
			want: NewElems(1, 2, 3)},
		{name: "IdenticalEnumerable",
			en:   enInt4(),
			args: args{en2: enInt4()},
			want: enEmpty},
		{name: "IdenticalEnumerable2",
			en:   enInt4(),
			args: args{en2: enInt4().Skip(2)},
			want: NewElems(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.Except(tt.args.en2)
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Except() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptSelf(t *testing.T) {
	i4 := enInt4()
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "NoComparerSpecified",
			en:   NewElems("A", "a", "b", "c", "b", "c"),
			args: args{en2: NewElems("b", "a", "d", "a")},
			want: NewElems("A", "c")},
		{name: "IntWithoutComparer",
			en:   NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10)},
			want: NewElems(1, 2, 3)},
		{name: "IdenticalEnumerable",
			en:   enInt4(),
			args: args{en2: enInt4()},
			want: enEmpty},
		{name: "IdenticalEnumerable2",
			en:   enInt4(),
			args: args{en2: enInt4().Skip(2)},
			want: NewElems(1, 2)},
		{name: "SameEnumerable",
			en:   i4,
			args: args{en2: i4.Skip(2)},
			want: NewElems(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptSelf(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptSelf() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptEq(t *testing.T) {
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
			want: NewElems("c")},
		{name: "IntComparerSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				eq: eqInt},
			want: NewElems(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptEq(tt.args.en2, tt.args.eq); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptEq() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptCmpMust(t *testing.T) {
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
		{name: "IntComparerSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp: cmpInt},
			want: NewElems(1, 2, 3)},
		{name: "CaseInsensitiveComparerSpecified",
			en: NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a"),
				cmp: cmpCaseInsensitive},
			want: NewElems("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptSelfCmpMust(t *testing.T) {
	i4 := enInt4()
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
		{name: "IntComparerSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp: cmpInt},
			want: NewElems(1, 2, 3)},
		{name: "CaseInsensitiveComparerSpecified",
			en: NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a"),
				cmp: cmpCaseInsensitive},
			want: NewElems("c")},
		{name: "SameEnumerable",
			en: i4,
			args: args{en2: i4.Skip(2),
				cmp: cmpInt},
			want: NewElems(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptSelfCmpMust(tt.args.en2, tt.args.cmp); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptSelfCmpMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptLsMust(t *testing.T) {
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
		{name: "IntLessSpecified",
			en: NewElems(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			args: args{en2: NewElems(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				ls: lsInt},
			want: NewElems(1, 2, 3)},
		{name: "CaseInsensitiveLessSpecified",
			en: NewElems("A", "a", "b", "c", "b"),
			args: args{en2: NewElems("b", "a", "d", "a"),
				ls: lsCaseInsensitive},
			want: NewElems("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptLsMust(tt.args.en2, tt.args.ls); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptLsMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ExceptSelfLsMust(t *testing.T) {
	i4 := enInt4()
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
		{name: "SameEnumerable",
			en: i4,
			args: args{en2: i4.Skip(3),
				ls: lsInt},
			want: NewElems(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ExceptSelfLsMust(tt.args.en2, tt.args.ls); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ExceptSelfLsMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
