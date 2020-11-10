package enumerable

import (
	"math"
	"reflect"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MinTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MaxTest.cs

func TestEnumerable_MinSel(t *testing.T) {
	type args struct {
		ls  common.Less
		sel func(common.Elem) common.Elem
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSelector",
			en:      enEmpty,
			args:    args{ls: func(e1, e2 common.Elem) bool { return true }},
			wantErr: true, expectedErr: errors.NilSel},
		{name: "EmptySequenceWithSelector",
			en:      enEmpty,
			args:    args{sel: common.Identity, ls: func(e1, e2 common.Elem) bool { return true }},
			wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SimpleSequenceWithSelector",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return len(e.(string)) },
				ls:  func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: 1},
		{name: "SimpleSequenceWithSelector2",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return []rune(e.(string))[0] },
				ls:  func(e1, e2 common.Elem) bool { return e1.(rune) < e2.(rune) }},
			want: '0'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.MinSel(tt.args.ls, tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.MinSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.MinSel() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MinSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_MinSelElMust(t *testing.T) {
	type args struct {
		ls  common.Less
		sel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "MinElement",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return len(e.(string)) },
				ls:  func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: "0"},
		{name: "MinElement2",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return []rune(e.(string))[0] },
				ls:  func(e1, e2 common.Elem) bool { return e1.(rune) < e2.(rune) }},
			want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.MinSelElMust(tt.args.ls, tt.args.sel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MinSelElMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_MinMust(t *testing.T) {
	type args struct {
		ls common.Less
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "SimpleSequenceNoSelector",
			en:   NewElems(5, 10, 6, 2, 13, 8),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: 2},
		{name: "AllNilSequence",
			en: NewElems(nil, nil, nil),
			args: args{ls: func(e1, e2 common.Elem) bool {
				if e2 == nil {
					return false
				}
				if e1 == nil {
					return true
				}
				return e1.(int) < e2.(int)
			}},
			want: nil},
		{name: "SequenceContainingBothInfinities",
			en:   NewElems(1., math.Inf(+1), math.Inf(-1)),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(float64) < e2.(float64) }},
			want: math.Inf(-1)},
		{name: "SequenceContainingNaN",
			en:   NewElems(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(float64) < e2.(float64) }},
			want: math.Inf(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.MinMust(tt.args.ls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MinMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_MaxSelMust(t *testing.T) {
	type args struct {
		ls  common.Less
		sel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "SimpleSequenceWithSelector",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return len(e.(string)) },
				ls:  func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: 5},
		{name: "AllNilSequenceWithSelector",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return nil },
				ls: func(e1, e2 common.Elem) bool {
					if e2 == nil {
						return false
					}
					if e1 == nil {
						return true
					}
					return e1.(int) < e2.(int)
				}},
			want: nil},
		{name: "SimpleSequenceWithSelector",
			en: NewElems("zyx", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return []rune(e.(string))[0] },
				ls:  func(e1, e2 common.Elem) bool { return e1.(rune) < e2.(rune) }},
			want: 'z'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.MaxSelMust(tt.args.ls, tt.args.sel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MaxSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_MaxSelElMust(t *testing.T) {
	type args struct {
		ls  common.Less
		sel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "MaxElement",
			en: NewElems("xyz", "ab", "abcde", "0"),
			args: args{
				sel: func(e common.Elem) common.Elem { return len(e.(string)) },
				ls:  func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.MaxSelElMust(tt.args.ls, tt.args.sel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MaxSelElMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_MaxMust(t *testing.T) {
	type args struct {
		ls common.Less
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "SimpleSequenceNoSelector",
			en:   NewElems(5, 10, 6, 2, 13, 8),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(int) < e2.(int) }},
			want: 13},
		{name: "SequenceIncludingNilsNoSelector",
			en: NewElems(5, nil, 10, nil, 6, nil, 2, 13, 8),
			args: args{ls: func(e1, e2 common.Elem) bool {
				if e2 == nil {
					return false
				}
				if e1 == nil {
					return true
				}
				return e1.(int) < e2.(int)
			}},
			want: 13},
		{name: "SimpleSequenceFloat64",
			en:   NewElems(-2.5, 2.5, 0.),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(float64) < e2.(float64) }},
			want: 2.5},
		{name: "SequenceContainingBothInfinities",
			en:   NewElems(1., math.Inf(+1), math.Inf(-1)),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(float64) < e2.(float64) }},
			want: math.Inf(+1)},
		{name: "SequenceContainingNaN",
			en:   NewElems(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			args: args{ls: func(e1, e2 common.Elem) bool { return e1.(float64) < e2.(float64) }},
			want: math.Inf(+1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.MaxMust(tt.args.ls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.MaxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
