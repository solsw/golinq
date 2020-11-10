package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AnyTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AllTest.cs

func TestEnumerable_Any(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want bool
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, want: false},
		{name: "NonEmptySequenceWithoutPredicate", en: NewElems(nil), want: true},
		{name: "1", en: enIntStr(), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Any(); got != tt.want {
				t.Errorf("Enumerable.Any() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AnyPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        bool
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: NewElems(1, 3, 5), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 10 }}, want: false},
		{name: "NonEmptySequenceWithPredicateMatchingElement", en: NewElems(1, 5, 20, 30),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 10 }}, want: true},
		{name: "NonEmptySequenceWithPredicateNotMatchingElement", en: NewElems(1, 5, 8, 9),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 10 }}, want: false},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch", en: NewElems(10, 2, 0, 3),
			args: args{pred: func(e common.Elem) bool {
				if e.(int) == 0 {
					panic("e.(int) == 0")
				}
				return true
			}},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.AnyPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.AnyPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.AnyPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Enumerable.AnyPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AnyPredMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "1", en: enInt4(),
			args: args{pred: func(e common.Elem) bool { return e.(int) == 4 }},
			want: true},
		{name: "2", en: enStr4(),
			args: args{pred: func(e common.Elem) bool { return len(e.(string)) == 4 }},
			want: true},
		{name: "3", en: enIntStr(),
			args: args{pred: func(e common.Elem) bool { _, ok := e.(int); return ok }},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.AnyPredMust(tt.args.pred); got != tt.want {
				t.Errorf("Enumerable.AnyPredMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_All(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        bool
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate",
			en:          NewElems(1, 3, 5),
			args:        args{pred: nil},
			wantErr:     true,
			expectedErr: errors.NilPred},
		{name: "EmptySequenceReturnsTrue",
			en:   enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 10 }},
			want: true},
		{name: "PredicateMatchingNoElements",
			en:   NewElems(1, 5, 20, 30),
			args: args{pred: func(e common.Elem) bool { return e.(int) < 0 }},
			want: false},
		{name: "PredicateMatchingSomeElements",
			en:   NewElems(1, 5, 8, 9),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }},
			want: false},
		{name: "PredicateMatchingAllElements",
			en:   NewElems(1, 5, 8, 9),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 0 }},
			want: true},
		{name: "SequenceIsNotEvaluatedAfterFirstNonMatch",
			en: NewElems(2, 10, 0, 3),
			args: args{pred: func(e common.Elem) bool {
				i := e.(int)
				if i == 0 {
					panic("e.(int) == 0")
				}
				return i > 2
			}},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.All(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.All() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.All() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Enumerable.All() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AllMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "1", en: enStr4(), args: args{pred: func(e common.Elem) bool { return len(e.(string)) >= 3 }}, want: true},
		{name: "2", en: enStr4(), args: args{pred: func(e common.Elem) bool { return len(e.(string)) > 3 }}, want: false},
		{name: "3", en: enIntStr(), args: args{pred: func(e common.Elem) bool { _, ok := e.(int); return ok }}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.AllMust(tt.args.pred); got != tt.want {
				t.Errorf("Enumerable.AllMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
