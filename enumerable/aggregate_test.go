package enumerable

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AggregateTest.cs

func TestEnumerable_Aggregate(t *testing.T) {
	type args struct {
		acc func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        common.Elem
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFuncUnseeded", en: enInt4(), args: args{acc: nil}, wantErr: true, expectedErr: errors.NilAcc},
		{name: "UnseededAggregation", en: NewElems(1, 4, 5),
			args: args{acc: func(ac, el common.Elem) common.Elem { return ac.(int)*2 + el.(int) }}, want: 17},
		{name: "EmptySequenceUnseeded", en: enEmpty,
			args:    args{acc: func(ac, el common.Elem) common.Elem { return ac.(int) + el.(int) }},
			wantErr: true, expectedErr: errors.EmptyEnum},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.Aggregate(tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.Aggregate() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.Aggregate() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.Aggregate() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AggregateMust(t *testing.T) {
	type args struct {
		acc func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want common.Elem
	}{
		{name: "UnseededSingleElementAggregation", en: NewElems(1),
			args: args{acc: func(ac, el common.Elem) common.Elem { return ac.(int)*2 + el.(int) }}, want: 1},
		{name: "FirstElementOfInputIsUsedAsSeedForUnseededOverload", en: NewElems(5, 3, 2),
			args: args{acc: func(ac, el common.Elem) common.Elem { return ac.(int) * el.(int) }}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.AggregateMust(tt.args.acc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.AggregateMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AggregateSeed(t *testing.T) {
	type args struct {
		seed interface{}
		acc  func(interface{}, common.Elem) interface{}
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        interface{}
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFuncSeeded", en: enInt4(),
			args:    args{seed: 5, acc: nil},
			wantErr: true, expectedErr: errors.NilAcc},
		{name: "SeededAggregation", en: NewElems(1, 4, 5),
			args: args{seed: 5, acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int)*2 + el.(int) }},
			want: 57},
		{name: "EmptySequenceSeeded", en: enEmpty,
			args: args{seed: 5, acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int) + el.(int) }},
			want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.AggregateSeed(tt.args.seed, tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.AggregateSeed() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.AggregateSeed() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.AggregateSeedr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AggregateSeedMust(t *testing.T) {
	type args struct {
		seed interface{}
		acc  func(interface{}, common.Elem) interface{}
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want interface{}
	}{
		{name: "DifferentSourceAndAccumulatorTypes", en: NewElems(int32(2000000000), int32(2000000000), int32(2000000000)),
			args: args{seed: int64(0), acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int64) + int64(el.(int32)) }},
			want: int64(6000000000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.AggregateSeedMust(tt.args.seed, tt.args.acc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.AggregateSeedMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_AggregateSeedSel(t *testing.T) {
	type args struct {
		seed interface{}
		acc  func(interface{}, common.Elem) interface{}
		sel  func(interface{}) interface{}
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        interface{}
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFuncSeededWithResultSelector", en: enInt4(),
			args: args{seed: 5,
				acc: nil,
				sel: func(el interface{}) interface{} { return fmt.Sprint(el) }},
			wantErr: true, expectedErr: errors.NilAcc},
		{name: "NilProjectionSeededWithResultSelector", en: enInt4(),
			args: args{seed: 5,
				acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int) + el.(int) },
				sel: nil},
			wantErr: true, expectedErr: errors.NilSel},
		{name: "SeededAggregationWithResultSelector", en: NewElems(1, 4, 5),
			args: args{seed: 5,
				acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int)*2 + el.(int) },
				sel: func(el interface{}) interface{} { return fmt.Sprint(el) }},
			want: "57"},
		{name: "EmptySequenceSeededWithResultSelector", en: enEmpty,
			args: args{seed: 5,
				acc: func(ac interface{}, el common.Elem) interface{} { return ac.(int) + el.(int) },
				sel: func(el interface{}) interface{} { return fmt.Sprint(el) }},
			want: "5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.AggregateSeedSel(tt.args.seed, tt.args.acc, tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.AggregateSeedSel() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.AggregateSeedSel() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.AggregateSeedSel() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
