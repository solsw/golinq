package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RepeatTest.cs

func TestRepeat(t *testing.T) {
	type args struct {
		el    common.Elem
		count int
	}
	tests := []struct {
		name        string
		args        args
		want        *Enumerable
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleRepeat", args: args{el: "foo", count: 3}, want: NewElems("foo", "foo", "foo")},
		{name: "EmptyRepeat", args: args{el: "foo", count: 0}, want: enEmpty},
		{name: "NilElement", args: args{el: nil, count: 2}, want: NewElems(nil, nil)},
		{name: "NegativeCount", args: args{el: "foo", count: -1}, wantErr: true, expectedErr: errors.NegCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Repeat(tt.args.el, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repeat() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Repeat() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Repeat() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestRepeatMust(t *testing.T) {
	type args struct {
		el    common.Elem
		count int
	}
	tests := []struct {
		name string
		args args
		want *Enumerable
	}{
		{name: "1", args: args{el: 0, count: 0}, want: enEmpty},
		{name: "2", args: args{el: 2, count: 2}, want: NewElems(2, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatMust(tt.args.el, tt.args.count); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("RepeatMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
