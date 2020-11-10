package enumerable

import (
	"fmt"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectTest.cs

func TestEnumerable_Select(t *testing.T) {
	var count int
	type args struct {
		sel func(common.Elem) common.Elem
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
			args: args{sel: common.Identity},
			want: enEmpty},
		{name: "NilProjectionThrowsNilArgumentException",
			en:          enInt4(),
			wantErr:     true,
			expectedErr: errors.NilSel},
		{name: "SimpleProjection",
			en:   NewElems(1, 5, 2),
			args: args{sel: func(e common.Elem) common.Elem { return e.(int) * 2 }},
			want: NewElems(2, 10, 4)},
		{name: "SimpleProjectionToDifferentType",
			en:   NewElems(1, 5, 2),
			args: args{sel: func(e common.Elem) common.Elem { return fmt.Sprint(e.(int)) }},
			want: NewElems("1", "5", "2")},
		{name: "SideEffectsInProjection1",
			en: NewElems(3, 2, 1), // Actual values won't be relevant
			args: args{sel: func(common.Elem) common.Elem {
				count++
				return count
			}},
			want: NewElems(1, 2, 3)},
		{name: "SideEffectsInProjection2",
			en: NewElems(1, 2, 3), // Actual values won't be relevant
			args: args{sel: func(common.Elem) common.Elem {
				count++
				return count
			}},
			want: NewElems(4, 5, 6)},
		{name: "SideEffectsInProjection3",
			en: NewElems(1, 2, 3), // Actual values won't be relevant
			args: args{sel: func(common.Elem) common.Elem {
				count++
				return count
			}},
			want: NewElems(11, 12, 13)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.Select(tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.Select() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.Select() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Select() = '%v', want '%v'", got, tt.want)
			}
		})
		if tt.name == "SideEffectsInProjection2" {
			count = 10
		}
	}
}

func TestEnumerable_SelectMust(t *testing.T) {
	type args struct {
		sel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "1",
			en:   enInt4(),
			args: args{sel: func(e common.Elem) common.Elem { return fmt.Sprint(e) }},
			want: NewElems("1", "2", "3", "4")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectMust(tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SelectIdx(t *testing.T) {
	type args struct {
		sel func(common.Elem, int) common.Elem
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
			args: args{sel: func(e common.Elem, i int) common.Elem { return e.(int) + i }},
			want: enEmpty},
		{name: "WithIndexNilProjectionThrowsNilArgumentException",
			en:          enInt4(),
			wantErr:     true,
			expectedErr: errors.NilSel},
		{name: "WithIndexSimpleProjection",
			en:   NewElems(1, 5, 2),
			args: args{sel: func(e common.Elem, i int) common.Elem { return e.(int) + i*10 }},
			want: NewElems(1, 15, 22)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.SelectIdx(tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.SelectIdx() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.SelectIdx() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectIdx() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_SelectIdxMust(t *testing.T) {
	type args struct {
		sel func(common.Elem, int) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "1",
			en:   enStr4(),
			args: args{sel: func(e common.Elem, i int) common.Elem { return len(e.(string)) + i }},
			want: NewElems(3, 4, 7, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.SelectIdxMust(tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.SelectIdxMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
