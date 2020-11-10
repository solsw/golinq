package enumerable

import (
	"math"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SumTest.cs

func TestEnumerable_SumIntMust(t *testing.T) {
	type args struct {
		sel func(common.Elem) int
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want int
	}{
		{name: "EmptySequenceIntWithSelector",
			en:   enEmpty,
			args: args{sel: func(e common.Elem) int { return len(e.(string)) }},
			want: 0},
		{name: "SimpleSumIntWithSelector",
			en:   NewElems("x", "abc", "de"),
			args: args{sel: func(e common.Elem) int { return len(e.(string)) }},
			want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SumIntMust(tt.args.sel); got != tt.want {
				t.Errorf("Enumerable.SumIntMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestEnumerable_SumFloat64Must(t *testing.T) {
	type args struct {
		sel func(common.Elem) float64
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want float64
	}{
		{name: "EmptySequenceFloat64WithSelector",
			en:   enEmpty,
			args: args{sel: func(e common.Elem) float64 { return float64(len(e.(string))) }},
			want: 0},
		{name: "SimpleSumFloat64WithSelector",
			en:   NewElems("x", "abc", "de"),
			args: args{sel: func(e common.Elem) float64 { return float64(len(e.(string))) }},
			want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SumFloat64Must(tt.args.sel); got != tt.want {
				t.Errorf("Enumerable.SumFloat64Must() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnum_SumFloat64MustIsNaN(t *testing.T) {
	type args struct {
		sel func(common.Elem) float64
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "SimpleSumFloat64WithSelectorWithNan",
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
			if want := math.IsNaN(tt.en.SumFloat64Must(tt.args.sel)); want != tt.want {
				t.Errorf("IsNaN(Enumerable.SumFloat64Must()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestEnum_SumFloat64Must(t *testing.T) {
	type args struct {
		sel func(common.Elem) float64
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "OverflowToNegInfinityFloat64",
			en:   NewElems(-math.MaxFloat64, -math.MaxFloat64),
			args: args{sel: func(e common.Elem) float64 { return e.(float64) }},
			want: true},
		{name: "OverflowToInfinityFloat64",
			en:   NewElems(math.MaxFloat64, math.MaxFloat64),
			args: args{sel: func(e common.Elem) float64 { return e.(float64) }},
			want: true},
		{name: "OverflowToInfinityFloat64WithSelector",
			en:   NewElems("x", "y"),
			args: args{sel: func(e common.Elem) float64 { return math.MaxFloat64 }},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.SumFloat64Must(tt.args.sel)
			var want bool
			switch tt.name {
			case "OverflowToNegInfinityFloat64":
				want = math.IsInf(got, -1)
			case "OverflowToInfinityFloat64":
				want = math.IsInf(got, +1)
			case "OverflowToInfinityFloat64WithSelector":
				want = math.IsInf(got, +1)
			}
			if want != tt.want {
				t.Errorf("IsInf(Enumerable.SumFloat64Must()) = %v, want %v", want, tt.want)
			}
		})
	}
}
