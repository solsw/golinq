package enumerable

import (
	"reflect"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstOrDefaultTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleOrDefaultTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastOrDefaultTest.cs

func TestEnumerable_First(t *testing.T) {
	tests := []struct {
		name        string
		en          *Enumerable
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.First()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.First() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.First() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.First() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_FirstPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate2", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate3", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 2, 5, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "EarlyOutAfterFirstElementWithPredicate", en: NewElems(15, 1, 0, 3),
			args: args{pred: func(e common.Elem) bool {
				if e.(int) == 0 {
					panic("e.(int) == 0")
				}
				return true
			}},
			want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.FirstPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.FirstPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.FirstPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.FirstPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_FirstOrDefault(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want common.Elem
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, want: nil},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, want: nil},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, want: nil},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.FirstOrDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.FirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_FirstOrDefaultPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate2", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate3", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 2, 5, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "EarlyOutAfterFirstElementWithPredicate", en: NewElems(15, 1, 0, 3),
			args: args{pred: func(e common.Elem) bool {
				if e.(int) == 0 {
					panic("e.(int) == 0")
				}
				return true
			}},
			want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.FirstOrDefaultPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.FirstOrDefaultPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.FirstOrDefaultPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.FirstOrDefaultPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_Single(t *testing.T) {
	tests := []struct {
		name        string
		en          *Enumerable
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), wantErr: true, expectedErr: errors.MultiElems},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.Single()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.Single() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.Single() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.Single() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SinglePred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate2", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate3", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 3, 5, 4, 2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 4 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.MultiMatch},
		{name: "EarlyOutWithPredicate", en: NewElems(1, 2, 0),
			args: args{pred: func(e common.Elem) bool {
				if e.(int) == 0 {
					panic("e.(int) == 0")
				}
				return true
			}},
			wantErr: true, expectedErr: errors.MultiMatch},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.SinglePred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.SinglePred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.SinglePred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.SinglePred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SingleOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		en          *Enumerable
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, want: nil},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, want: nil},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, want: nil},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), wantErr: true, expectedErr: errors.MultiElems},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.SingleOrDefault()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.SingleOrDefault() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.SingleOrDefault() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.SingleOrDefault() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SingleOrDefaultPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 2, 5, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.MultiMatch},
		{name: "EarlyOutWithPredicate", en: NewElems(1, 2, 0),
			args: args{pred: func(e common.Elem) bool {
				if e.(int) == 0 {
					panic("e.(int) == 0")
				}
				return true
			}},
			wantErr: true, expectedErr: errors.MultiMatch},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.SingleOrDefaultPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.SingleOrDefaultPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.SingleOrDefaultPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.SingleOrDefaultPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_Last(t *testing.T) {
	tests := []struct {
		name        string
		en          *Enumerable
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.Last()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.Last() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.Last() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.Last() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_LastPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate2", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "EmptySequenceWithPredicate3", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, wantErr: true, expectedErr: errors.NoMatch},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 2, 5, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.LastPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.LastPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.LastPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.LastPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_LastOrDefaultMust(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want common.Elem
	}{
		{name: "EmptySequenceWithoutPredicate", en: enEmpty, want: nil},
		{name: "EmptySequenceWithoutPredicate2", en: enEmptyEnmr, want: nil},
		{name: "EmptySequenceWithoutPredicate3", en: enEmptySlice, want: nil},
		{name: "SingleElementSequenceWithoutPredicate", en: NewElems(5), want: 5},
		{name: "MultipleElementSequenceWithoutPredicate", en: NewElems(5, 10), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.LastOrDefaultMust(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.LastOrDefaultMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_LastOrDefaultPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilPredicate", en: enInt4(), args: args{pred: nil}, wantErr: true, expectedErr: errors.NilPred},
		{name: "EmptySequenceWithPredicate", en: enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate2", en: enEmptyEnmr,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "EmptySequenceWithPredicate3", en: enEmptySlice,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "SingleElementSequenceWithMatchingPredicate", en: NewElems(5),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "SingleElementSequenceWithNonMatchingPredicate", en: NewElems(2),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithNoPredicateMatches", en: NewElems(1, 2, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: nil},
		{name: "MultipleElementSequenceWithSinglePredicateMatch", en: NewElems(1, 2, 5, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 5},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches", en: NewElems(1, 2, 5, 10, 2, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) > 3 }}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.LastOrDefaultPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.LastOrDefaultPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.LastOrDefaultPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.LastOrDefaultPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
