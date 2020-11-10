package enumerable

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ConcatTest.cs

func TestEnumerable_Concat(t *testing.T) {
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "Empty1", en: enEmpty, args: args{en2: enEmpty}, want: enEmpty},
		{name: "Empty2", en: enEmptyEnmr, args: args{en2: enEmptySlice}, want: enEmpty},
		{name: "Empty3", en: enInt4(), args: args{en2: enEmptySlice}, want: enInt4()},
		{name: "Empty4", en: enEmptyEnmr, args: args{en2: enStr4()}, want: enStr4()},
		{name: "SimpleConcatenation",
			en:   NewElems("a", "b"),
			args: args{en2: NewElems("c", "d")},
			want: NewElems("a", "b", "c", "d")},
		{name: "SimpleConcatenation2",
			en:   RangeMust(1, 2),
			args: args{en2: RepeatMust("3", 1)},
			want: NewElems(1, 2, "3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Concat(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Concat() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ConcatSelf(t *testing.T) {
	i4 := enInt4()
	rs := RepeatMust("q", 2).Skip(1)
	rg := RangeMust(1, 4)
	type args struct {
		en2 *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SameEnumerable",
			en:   i4,
			args: args{en2: i4},
			want: NewElems(1, 2, 3, 4, 1, 2, 3, 4)},
		{name: "SameEnumerable2",
			en:   rs,
			args: args{en2: rs},
			want: NewElems("q", "q")},
		{name: "SameEnumerable3",
			en:   rg.Take(2),
			args: args{en2: rg.Skip(2)},
			want: NewElems(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ConcatSelf(tt.args.en2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ConcatSelf() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
