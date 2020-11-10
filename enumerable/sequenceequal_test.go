package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SequenceEqualTest.cs

func TestEnumerable_SequenceEqual(t *testing.T) {
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "001", want: true},
		{name: "002", en: enEmptyEnmr, args: args{en2: enEmpty}, want: true},
		{name: "003", en: enEmptySlice, args: args{en2: enEmpty}, want: true},
		{name: "004", en: enEmpty, args: args{en2: enEmptyEnmr}, want: true},
		{name: "005", en: enEmpty, args: args{en2: enEmptySlice}, want: true},
		{name: "01", en: NewElems(1), args: args{en2: enEmpty}, want: false},
		{name: "02", en: enEmpty, args: args{en2: NewElems(1)}, want: false},
		{name: "1", en: NewElems(1), args: args{en2: NewElems(1)}, want: true},
		{name: "2", en: enStr4(), args: args{en2: enStr4()}, want: true},
		{name: "4", en: NewElems("a", "b"), args: args{en2: NewElems("a")}, want: false},
		{name: "5", en: NewElems("a"), args: args{en2: NewElems("a", "b")}, want: false},
		{name: "UnequalLengthsBothArrays", en: NewElems(1, 5, 3), args: args{en2: NewElems(1, 5, 3, 10)}, want: false},
		{name: "UnequalLengthsBothRangesFirstLonger", en: RangeMust(0, 11), args: args{en2: RangeMust(0, 10)}, want: false},
		{name: "UnequalLengthsBothRangesSecondLonger", en: RangeMust(0, 10), args: args{en2: RangeMust(0, 11)}, want: false},
		{name: "UnequalData", en: NewElems(1, 5, 3, 9), args: args{en2: NewElems(1, 5, 3, 10)}, want: false},
		{name: "EqualDataBothArrays", en: NewElems(1, 5, 3, 10), args: args{en2: NewElems(1, 5, 3, 10)}, want: true},
		{name: "EqualDataBothRanges", en: RangeMust(0, 10), args: args{en2: RangeMust(0, 10)}, want: true},
		{name: "OrderMatters", en: NewElems(1, 2), args: args{en2: NewElems(2, 1)}, want: false},
		{name: "ReturnAtFirstDifference",
			en:   NewElems(1, 5, 10, 2, 0).SelectMust(func(e common.Elem) common.Elem { return 10 / e.(int) }),
			args: args{en2: NewElems(1, 5, 10, 1, 0).SelectMust(func(e common.Elem) common.Elem { return 10 / e.(int) })},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SequenceEqual(tt.args.en2); got != tt.want {
				t.Errorf("Enumerable.SequenceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SequenceEqualEq(t *testing.T) {
	type args struct {
		en2 *Enumerable
		eq  common.Equality
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "1",
			en: NewElems("a", "b"),
			args: args{en2: NewElems("a", "B"),
				eq: eqCaseInsensitive},
			want: true},
		{name: "CustomEqualityComparer",
			en: NewElems("foo", "BAR", "baz"),
			args: args{en2: NewElems("FOO", "bar", "Baz"),
				eq: eqCaseInsensitive},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SequenceEqualEq(tt.args.en2, tt.args.eq); got != tt.want {
				t.Errorf("Enumerable.SequenceEqualEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SequenceEqualSelf(t *testing.T) {
	r0 := RangeMust(0, 0)
	r1 := RangeMust(0, 1)
	r2 := RangeMust(0, 2)
	r3 := RepeatMust(1, 4)
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "Same0", en: r0, args: args{en2: r0}, want: true},
		{name: "Same1", en: r1, args: args{en2: r1}, want: true},
		{name: "Same2", en: r2, args: args{en2: r2}, want: true},
		{name: "Same3", en: r3, args: args{en2: r3.Skip(2)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SequenceEqualSelf(tt.args.en2); got != tt.want {
				t.Errorf("Enumerable.SequenceEqualSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
