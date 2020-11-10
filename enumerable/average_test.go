package enumerable

import (
	"math"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AverageTest.cs

func TestEnumerable_AverageInt(t *testing.T) {
	type args struct {
		sel func(common.Elem) int
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        float64
		wantErr     bool
		expectedErr error
	}{
		{name: "SourceStrNilSelector", en: enStr4(), wantErr: true, expectedErr: errors.NilSel},
		{name: "EmptySequenceIntNoSelector", en: enEmpty, wantErr: true, expectedErr: errors.NilSel},
		{name: "EmptySequenceIntWithSelector",
			en:      enEmpty,
			args:    args{sel: func(e common.Elem) int { return len(e.(string)) }},
			wantErr: true, expectedErr: errors.EmptyEnum},
		{name: "SimpleAverageInt",
			en:   NewElems(5, 10, 0, 15),
			args: args{sel: func(e common.Elem) int { return e.(int) }},
			want: 7.5},
		{name: "SimpleAverageIntWithSelector",
			en:   NewElems("", "abcd", "a", "b"),
			args: args{sel: func(e common.Elem) int { return len(e.(string)) }},
			want: 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.AverageInt(tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.AverageInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.AverageInt() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Enumerable.AverageInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnum_AverageFloat64MustIsNaN(t *testing.T) {
	type args struct {
		sel func(common.Elem) float64
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "SequenceContainingNan",
			en: NewElems("x", "abc", "de"),
			args: args{sel: func(e common.Elem) float64 {
				l := len(e.(string))
				if l == 3 {
					return math.NaN()
				}
				return float64(l)
			}},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.AverageFloat64Must(tt.args.sel)
			want := math.IsNaN(got)
			if want != tt.want {
				t.Errorf("IsNaN(Enumerable.AverageFloat64Must()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestEnum_AverageFloat64MustIsInf(t *testing.T) {
	type args struct {
		sel func(common.Elem) float64
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "Float64OverflowsToInfinity",
			en:   NewElems(math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64),
			args: args{sel: func(e common.Elem) float64 { return e.(float64) }},
			want: true},
		{name: "Float64OverflowsToNegInfinity",
			en:   NewElems(-math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, math.MaxFloat64),
			args: args{sel: func(e common.Elem) float64 { return e.(float64) }},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.AverageFloat64Must(tt.args.sel)
			var want bool
			switch tt.name {
			case "Float64OverflowsToInfinity":
				want = math.IsInf(got, +1)
			case "Float64OverflowsToNegInfinity":
				want = math.IsInf(got, -1)
			}
			if want != tt.want {
				t.Errorf("IsInf(Enumerable.AverageFloat64Must()) = %v, want %v", want, tt.want)
			}
		})
	}
}
