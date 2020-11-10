package enumerable

import (
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/WhereTest.cs

func TestEnumerable_Where(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        *Enumerable
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySource",
			en:   enEmpty,
			args: args{pred: func(e common.Elem) bool { return e.(int) > 5 }},
			want: enEmpty},
		{name: "NilPredicateThrowsNilArgumentException",
			en:          enInt4(),
			wantErr:     true,
			expectedErr: errors.NilPred},
		{name: "SimpleFiltering",
			en:   NewElems(1, 3, 4, 2, 8, 1),
			args: args{pred: func(e common.Elem) bool { return e.(int) < 4 }},
			want: NewElems(1, 3, 2, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.Where(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.Where() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.Where() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Where() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_WhereMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "1", en: enInt4(), args: args{pred: func(common.Elem) bool { return false }}, want: enEmpty},
		{name: "2", en: enInt4(), args: args{pred: func(common.Elem) bool { return true }}, want: enInt4()},
		{name: "3", en: enInt4(), args: args{pred: func(e common.Elem) bool { return e.(int)%2 == 1 }}, want: NewElems(1, 3)},
		{name: "4", en: enStr5(), args: args{pred: func(common.Elem) bool { return true }}, want: enStr5()},
		{name: "5", en: enStr5(),
			args: args{pred: func(e common.Elem) bool { return strings.HasPrefix(e.(string), "t") }},
			want: NewElems("two", "three")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.WhereMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.WhereMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_WhereIdxMust(t *testing.T) {
	type args struct {
		pred common.PredicateIdx
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "EmptySourceWithIndex",
			en:   enEmpty,
			args: args{pred: func(e common.Elem, i int) bool { return e.(int) > 5 }},
			want: enEmpty},
		{name: "WithIndexSimpleFiltering",
			en:   NewElems(1, 3, 4, 2, 8, 1),
			args: args{pred: func(e common.Elem, i int) bool { return e.(int) < i }},
			want: NewElems(2, 1)},
		{name: "1",
			en:   enStr5(),
			args: args{pred: func(e common.Elem, i int) bool { return len(e.(string)) == i }},
			want: NewElems("five")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.WhereIdxMust(tt.args.pred); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.WhereIdxMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
