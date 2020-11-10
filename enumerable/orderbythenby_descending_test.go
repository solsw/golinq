package enumerable

import (
	"math"
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByDescendingTest.cs

func TestEnumerable_OrderByMust(t *testing.T) {
	type args struct {
		ls common.Less
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "1234", en: NewElems(4, 1, 3, 2), args: args{ls: lsInt}, want: NewElems(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.OrderByMust(tt.args.ls).Enumerable(); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.OrderByMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_OrderBySelMust(t *testing.T) {
	type args struct {
		ls   common.Less
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SimpleUniqueKeys",
			en:   NewElems(elel{1, 10}, elel{2, 12}, elel{3, 11}),
			args: args{ls: lsInt, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(1, 3, 2)},
		{name: "NilsAreFirst",
			en:   NewElems(elel{1, "abc"}, elel{2, nil}, elel{3, "def"}),
			args: args{ls: lsNilStr, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(2, 1, 3)},
		{name: "OrderingIsStable",
			en:   NewElems(elel{1, 10}, elel{2, 11}, elel{3, 11}, elel{4, 10}),
			args: args{ls: lsInt, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(1, 4, 2, 3)},
		{name: "CustomLess", en: NewElems(elel{1, 15}, elel{2, -13}, elel{3, 11}),
			args: args{
				ls: func(e1, e2 common.Elem) bool {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					return f1 < f2
				},
				ksel: func(e common.Elem) common.Elem { return e.(elel).e2 },
			},
			want: NewElems(3, 2, 1)},
		{name: "CustomComparer", en: NewElems(elel{1, 15}, elel{2, -13}, elel{3, 11}),
			args: args{
				ls: common.ComparisonToLess(func(e1, e2 common.Elem) int {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
				ksel: func(e common.Elem) common.Elem { return e.(elel).e2 },
			},
			want: NewElems(3, 2, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.OrderBySelMust(tt.args.ls, tt.args.ksel).
				Enumerable().
				SelectMust(func(e common.Elem) common.Elem { return e.(elel).e1 })
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.OrderBySelMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_OrderByDescendingSelMust(t *testing.T) {
	type args struct {
		ls   common.Less
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SimpleUniqueKeys", en: NewElems(elel{1, 10}, elel{2, 12}, elel{3, 11}),
			args: args{ls: lsInt, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(2, 3, 1)},
		{name: "NilsAreLast", en: NewElems(elel{1, "abc"}, elel{2, nil}, elel{3, "def"}),
			args: args{ls: lsNilStr, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(3, 1, 2)},
		{name: "OrderingIsStable", en: NewElems(elel{1, 10}, elel{2, 11}, elel{3, 11}, elel{4, 10}),
			args: args{ls: lsInt, ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(2, 3, 1, 4)},
		// NilComparerIsDefault
		{name: "CustomLess", en: NewElems(elel{1, 15}, elel{2, -13}, elel{3, 11}),
			args: args{
				ls: func(e1, e2 common.Elem) bool {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					return f1 < f2
				},
				ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(1, 2, 3),
		},
		{name: "CustomComparer", en: NewElems(elel{1, 15}, elel{2, -13}, elel{3, 11}),
			args: args{
				ls: common.ComparisonToLess(func(e1, e2 common.Elem) int {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
				ksel: func(e common.Elem) common.Elem { return e.(elel).e2 }},
			want: NewElems(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.en.OrderByDescendingSelMust(tt.args.ls, tt.args.ksel).
				Enumerable().
				SelectMust(func(e common.Elem) common.Elem { return e.(elel).e1 })
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.OrderByDescendingSelMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestOrderedEnumerable_ThenByMust(t *testing.T) {
	type args struct {
		ls common.Less
	}
	tests := []struct {
		name string
		oe   *OrderedEnumerable
		args args
		want *Enumerable
	}{
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			oe: (NewElems(eee{1, 10, 22}, eee{2, 12, 21}, eee{3, 10, 20})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{ls: func(el1, el2 common.Elem) bool { return el1.(eee).e3.(int) < el2.(eee).e3.(int) }},
			want: NewElems(3, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.oe.ThenByMust(tt.args.ls).
				Enumerable().
				SelectMust(func(e common.Elem) common.Elem { return e.(eee).e1 })
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderedEnumerable.ThenByMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestOrderedEnumerable_ThenBySelMust(t *testing.T) {
	type args struct {
		ls   common.Less
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		oe   *OrderedEnumerable
		args args
		want *Enumerable
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			oe: (NewElems(eee{1, 10, 20}, eee{2, 12, 21}, eee{3, 11, 22})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(1, 3, 2)},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			oe: (NewElems(eee{1, 10, 22}, eee{2, 12, 21}, eee{3, 10, 20})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(3, 1, 2)},
		{name: "TertiaryKeys",
			oe: (NewElems(eeee{1, 10, 22, 30}, eeee{2, 12, 21, 31}, eeee{3, 10, 20, 33}, eeee{4, 10, 20, 32})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e2 }).
				ThenBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e3 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eeee).e4 }},
			want: NewElems(4, 3, 1, 2)},
		{name: "ThenByAfterOrderByDescending",
			oe: (NewElems(eee{1, 10, 22}, eee{2, 12, 21}, eee{3, 10, 20})).
				OrderByDescendingSelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(2, 3, 1)},
		{name: "NilsAreFirst",
			oe: (NewElems(eee{1, 1, "abc"}, eee{2, 1, nil}, eee{3, 1, "def"})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsNilStr,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(2, 1, 3)},
		{name: "OrderingIsStable",
			oe: (NewElems(eee{1, 1, 10}, eee{2, 1, 11}, eee{3, 1, 11}, eee{4, 1, 10})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(1, 4, 2, 3)},
		{name: "CustomLess",
			oe: (NewElems(eee{1, 1, 15}, eee{2, 1, -13}, eee{3, 1, 11})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls: func(e1, e2 common.Elem) bool {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					return f1 < f2
				},
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(3, 2, 1)},
		{name: "CustomComparer",
			oe: (NewElems(eee{1, 1, 15}, eee{2, 1, -13}, eee{3, 1, 11})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls: common.ComparisonToLess(func(e1, e2 common.Elem) int {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(3, 2, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.oe.ThenBySelMust(tt.args.ls, tt.args.ksel).Enumerable()
			switch tt.name {
			case "TertiaryKeys":
				got = got.SelectMust(func(e common.Elem) common.Elem { return e.(eeee).e1 })
			default:
				got = got.SelectMust(func(e common.Elem) common.Elem { return e.(eee).e1 })
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderedEnumerable.ThenBySelMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestOrderedEnumerable_ThenByDescendingSelMust(t *testing.T) {
	type args struct {
		ls   common.Less
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		oe   *OrderedEnumerable
		args args
		want *Enumerable
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			oe: (NewElems(eee{1, 10, 20}, eee{2, 12, 21}, eee{3, 11, 22})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(1, 3, 2)},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			oe: (NewElems(eee{1, 10, 19}, eee{2, 12, 21}, eee{3, 10, 20})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(3, 1, 2)},
		{name: "TertiaryKeys",
			oe: (NewElems(eeee{1, 10, 22, 30}, eeee{2, 12, 21, 31}, eeee{3, 10, 20, 33}, eeee{4, 10, 20, 32})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e2 }).
				ThenByDescendingSelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e3 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eeee).e4 }},
			want: NewElems(1, 3, 4, 2)},
		{name: "TertiaryKeysWithMixedOrdering",
			oe: (NewElems(eeee{1, 10, 22, 30}, eeee{2, 12, 21, 31}, eeee{3, 10, 20, 33}, eeee{4, 10, 20, 32})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e2 }).
				ThenBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eeee).e3 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eeee).e4 }},
			want: NewElems(3, 4, 1, 2)},
		{name: "ThenByDescendingAfterOrderByDescending",
			oe: (NewElems(eee{1, 10, 22}, eee{2, 12, 21}, eee{3, 10, 20})).
				OrderByDescendingSelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(2, 1, 3)},
		{name: "NilsAreLast",
			oe: (NewElems(eee{1, 1, "abc"}, eee{2, 1, nil}, eee{3, 1, "def"})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsNilStr,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(3, 1, 2)},
		{name: "OrderingIsStable",
			oe: (NewElems(eee{1, 1, 10}, eee{2, 1, 11}, eee{3, 1, 11}, eee{4, 1, 10})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls:   lsInt,
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 }},
			want: NewElems(2, 3, 1, 4)},
		{name: "CustomLess",
			oe: (NewElems(eee{1, 1, 15}, eee{2, 1, -13}, eee{3, 1, 11})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls: func(e1, e2 common.Elem) bool {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					return f1 < f2
				},
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 },
			},
			want: NewElems(1, 2, 3)},
		{name: "CustomComparer",
			oe: (NewElems(eee{1, 1, 15}, eee{2, 1, -13}, eee{3, 1, 11})).
				OrderBySelMust(lsInt, func(e common.Elem) common.Elem { return e.(eee).e2 }),
			args: args{
				ls: common.ComparisonToLess(func(e1, e2 common.Elem) int {
					f1 := math.Abs(float64(e1.(int)))
					f2 := math.Abs(float64(e2.(int)))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
				ksel: func(e common.Elem) common.Elem { return e.(eee).e3 },
			},
			want: NewElems(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.oe.ThenByDescendingSelMust(tt.args.ls, tt.args.ksel).Enumerable()
			if strings.HasPrefix(tt.name, "TertiaryKeys") {
				got = got.SelectMust(func(e common.Elem) common.Elem { return e.(eeee).e1 })
			} else {
				got = got.SelectMust(func(e common.Elem) common.Elem { return e.(eee).e1 })
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderedEnumerable.ThenByDescendingSelMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
