package enumerable

import (
	"math"
	"testing"

	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RangeTest.cs

func TestRange(t *testing.T) {
	type args struct {
		start int
		count int
	}
	tests := []struct {
		name        string
		args        args
		want        *Enumerable
		wantErr     bool
		expectedErr error
	}{
		{name: "NegativeCount", args: args{start: 10, count: -1}, wantErr: true, expectedErr: errors.NegCount},
		{name: "CountTooLarge1", args: args{start: math.MaxInt32, count: 2}, wantErr: true, expectedErr: errors.WrongStrtCnt},
		{name: "CountTooLarge2", args: args{start: 2, count: math.MaxInt32}, wantErr: true, expectedErr: errors.WrongStrtCnt},
		{name: "CountTooLarge3", args: args{start: math.MaxInt32 / 2, count: math.MaxInt32/2 + 3},
			wantErr: true, expectedErr: errors.WrongStrtCnt},
		{name: "LargeButValidCount1", args: args{start: math.MaxInt32, count: 1}, want: NewElems(math.MaxInt32)},
		// max length of Enumerable depends on available memory
		// {name: "LargeButValidCount2", args: args{start: 1, count: math.MaxInt32 / 4}}, // works about 3 minutes on my notebook
		// {name: "LargeButValidCount3", args: args{start: math.MaxInt32 / 2, count: math.MaxInt32/2 + 2}, wantErr: false},
		{name: "ValidRange", args: args{start: 5, count: 3}, want: NewElems(5, 6, 7)},
		{name: "NegativeStart", args: args{start: -2, count: 5}, want: NewElems(-2, -1, 0, 1, 2)},
		{name: "EmptyRange", args: args{start: 100, count: 0}, want: enEmpty},
		{name: "SingleValueOfMaxInt32", args: args{start: math.MaxInt32, count: 1}, want: NewElems(math.MaxInt32)},
		{name: "EmptyRangeStartingAtMinInt32", args: args{start: math.MinInt32, count: 0}, want: enEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Range(tt.args.start, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Range() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Range() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
