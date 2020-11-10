package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

func TestEnumerable_Count(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want int
	}{
		{name: "NonCollectionCount", en: RangeMust(2, 5), want: 5},
		{name: "0", en: enEmpty, want: 0},
		{name: "1", en: enInt4(), want: 4},
		{name: "2", en: enStr4(), want: 4},
		{name: "3", en: enStr5(), want: 5},
		{name: "4", en: enIntStr(), want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Count(); got != tt.want {
				t.Errorf("Enumerable.Count() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_CountPred(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "PredicatedNilPredicateThrowsArgumentNilException",
			en:      enInt4(),
			args:    args{pred: nil},
			wantErr: true, expectedErr: errors.NilPred},
		{name: "PredicatedCount",
			en:   RangeMust(2, 5),
			args: args{pred: func(e common.Elem) bool { return e.(int)%2 == 0 }},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.CountPred(tt.args.pred)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.CountPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.CountPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Enumerable.CountPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_CountPredMust(t *testing.T) {
	type args struct {
		pred common.Predicate
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want int
	}{
		{name: "11", en: enInt4(), args: args{pred: func(common.Elem) bool { return false }}, want: 0},
		{name: "12", en: enInt4(), args: args{pred: func(common.Elem) bool { return true }}, want: 4},
		{name: "21", en: enStr4(), args: args{pred: func(e common.Elem) bool { return len(e.(string)) == 3 }}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.CountPredMust(tt.args.pred); got != tt.want {
				t.Errorf("Enumerable.CountPredMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
