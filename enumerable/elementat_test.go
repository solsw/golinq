package enumerable

import (
	"reflect"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtOrDefaultTest.cs

func TestEnumerable_ElementAt(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NegativeIndex",
			en:      enInt4(),
			args:    args{idx: -1},
			wantErr: true, expectedErr: errors.IdxRange},
		{name: "OvershootIndex",
			en:      enInt4(),
			args:    args{idx: 4},
			wantErr: true, expectedErr: errors.IdxRange},
		{name: "ValidIndex",
			en:   enStr4(),
			args: args{idx: 2},
			want: "three"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.ElementAt(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.ElementAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.ElementAt() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.ElementAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ElementAtOrDefault(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "NegativeIndex",
			en:   enInt4(),
			args: args{idx: -1},
			want: nil},
		{name: "OvershootIndex",
			en:   enInt4(),
			args: args{idx: 4},
			want: nil},
		{name: "ValidIndex",
			en:   enStr4(),
			args: args{idx: 2},
			want: "three"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ElementAtOrDefault(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.ElementAtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
