package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeWhileTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipWhileTest.cs

func TestEnumerable_Take(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "NegativeCount", en: RangeMust(0, 5), args: args{count: -5}, want: enEmpty},
		{name: "ZeroCount", en: RangeMust(0, 5), args: args{count: 0}, want: enEmpty},
		{name: "CountShorterThanSource", en: RangeMust(0, 5), args: args{count: 3}, want: NewElems(0, 1, 2)},
		{name: "CountShorterThanSource2", en: enInt4(), args: args{count: 3}, want: NewElems(1, 2, 3)},
		{name: "CountEqualToSourceLength", en: RangeMust(1, 5), args: args{count: 5}, want: NewElems(1, 2, 3, 4, 5)},
		{name: "CountGreaterThanSourceLength", en: RangeMust(2, 5), args: args{count: 100}, want: NewElems(2, 3, 4, 5, 6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Take(tt.args.count); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Take() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_Skip(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "NegativeCount", en: RangeMust(0, 5), args: args{count: -5}, want: NewElems(0, 1, 2, 3, 4)},
		{name: "ZeroCount", en: RangeMust(0, 5), args: args{count: 0}, want: NewElems(0, 1, 2, 3, 4)},
		{name: "CountShorterThanSource", en: RangeMust(0, 5), args: args{count: 3}, want: NewElems(3, 4)},
		{name: "CountEqualToSourceLength", en: RangeMust(0, 5), args: args{count: 5}, want: enEmpty},
		{name: "CountGreaterThanSourceLength", en: RangeMust(0, 5), args: args{count: 100}, want: enEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Skip(tt.args.count); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Skip() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_TakeWhileMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "PredicateFailingFirstElement",
			en:   NewElems("zero", "one", "two", "three", "four", "five", "six"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) > 4 }},
			want: enEmpty},
		{name: "PredicateMatchingSomeElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) < 5 }},
			want: NewElems("zero", "one", "two")},
		{name: "PredicateMatchingAllElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) < 100 }},
			want: NewElems("zero", "one", "two", "three", "four", "five")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.TakeWhileMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.TakeWhileMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_TakeWhileIdxMust(t *testing.T) {
	type args struct {
		pred common.PredicateIdx
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return i+len(e.(string)) > 4 }},
			want: enEmpty},
		{name: "PredicateWithIndexMatchingSomeElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return len(e.(string)) != i }},
			want: NewElems("zero", "one", "two", "three")},
		{name: "PredicateWithIndexMatchingAllElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return len(e.(string)) < 100 }},
			want: NewElems("zero", "one", "two", "three", "four", "five")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.TakeWhileIdxMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.TakeWhileIdxMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SkipWhileMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "PredicateFailingFirstElement",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) > 4 }},
			want: NewElems("zero", "one", "two", "three", "four", "five")},
		{name: "PredicateMatchingSomeElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) < 5 }},
			want: NewElems("three", "four", "five")},
		{name: "PredicateMatchingAllElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) < 100 }},
			want: enEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SkipWhileMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SkipWhileMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SkipWhileIdxMust(t *testing.T) {
	type args struct {
		pred common.PredicateIdx
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return i+len(e.(string)) > 4 }},
			want: NewElems("zero", "one", "two", "three", "four", "five")},
		{name: "PredicateWithIndexMatchingSomeElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return len(e.(string)) > i }},
			want: NewElems("four", "five")},
		{name: "PredicateWithIndexMatchingAllElements",
			en:   NewElems("zero", "one", "two", "three", "four", "five"),
			args: args{pred: func(e common.Elem, i int) bool { return len(e.(string)) < 100 }},
			want: enEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SkipWhileIdxMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SkipWhileIdxMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
