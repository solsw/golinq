package enumerable

import (
	"fmt"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectManyTest.cs

func TestEnumerable_SelectManyMust(t *testing.T) {
	type args struct {
		sel func(common.Elem) *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SimpleFlatten", en: NewElems(3, 5, 20, 15),
			args: args{sel: func(el common.Elem) *Enumerable {
				rr := []rune(fmt.Sprint(el.(int)))
				var ee []common.Elem
				for _, r := range rr {
					ee = append(ee, r)
				}
				return NewElems(ee...)
			}},
			want: NewElems('3', '5', '2', '0', '1', '5'),
		},
		{name: "1", en: enInt4(),
			args: args{sel: func(e common.Elem) *Enumerable {
				i := e.(int)
				return NewElems(i, i*i)
			}},
			want: NewElems(1, 1, 2, 4, 3, 9, 4, 16),
		},
		{name: "2", en: enInt4(),
			args: args{sel: func(e common.Elem) *Enumerable {
				i := e.(int)
				if i%2 == 0 {
					return enEmpty
				}
				return NewElems(i, i*i)
			}},
			want: NewElems(1, 1, 3, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectManyMust(tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectManyMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SelectManyIdxMust(t *testing.T) {
	type args struct {
		sel func(common.Elem, int) *Enumerable
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "SimpleFlattenWithIndex", en: NewElems(3, 5, 20, 15),
			args: args{sel: func(el common.Elem, i int) *Enumerable {
				rr := []rune(fmt.Sprint(el.(int) + i))
				var ee []common.Elem
				for _, r := range rr {
					ee = append(ee, r)
				}
				return NewElems(ee...)
			}},
			want: NewElems('3', '6', '2', '2', '1', '8'),
		},
		{name: "1", en: enInt4(),
			args: args{sel: func(e common.Elem, idx int) *Enumerable {
				if idx%2 == 0 {
					return enEmpty
				}
				i := e.(int)
				return NewElems(i, i*i)
			}},
			want: NewElems(2, 4, 4, 16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectManyIdxMust(tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectManyIdxMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SelectManyCollMust(t *testing.T) {
	type args struct {
		sel1 func(common.Elem) *Enumerable
		sel2 func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "FlattenWithProjection", en: NewElems(3, 5, 20, 15),
			args: args{
				sel1: func(el common.Elem) *Enumerable {
					rr := []rune(fmt.Sprint(el.(int)))
					var ee []common.Elem
					for _, r := range rr {
						ee = append(ee, r)
					}
					return NewElems(ee...)
				},
				sel2: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%d: %s", e1.(int), string(e2.(rune)))
				},
			},
			want: NewElems("3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectManyCollMust(tt.args.sel1, tt.args.sel2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectManyCollMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SelectManyIdxCollMust(t *testing.T) {
	type args struct {
		sel1 func(common.Elem, int) *Enumerable
		sel2 func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "FlattenWithProjectionAndIndex", en: NewElems(3, 5, 20, 15),
			args: args{
				sel1: func(el common.Elem, i int) *Enumerable {
					rr := []rune(fmt.Sprint(el.(int) + i))
					var ee []common.Elem
					for _, r := range rr {
						ee = append(ee, r)
					}
					return NewElems(ee...)
				},
				sel2: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%d: %s", e1.(int), string(e2.(rune)))
				},
			},
			want: NewElems("3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectManyIdxCollMust(tt.args.sel1, tt.args.sel2); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectManyIdxCollMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
