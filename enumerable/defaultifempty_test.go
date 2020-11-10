package enumerable

import (
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DefaultIfEmptyTest.cs

func TestEnumerable_DefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want *Enumerable
	}{
		{name: "EmptySequenceNoDefaultValue", en: enEmpty, want: NewElems(nil)},
		{name: "EmptySequenceNoDefaultValue2", en: enEmptyEnmr, want: NewElems(nil)},
		{name: "EmptySequenceNoDefaultValue3", en: enEmptySlice, want: NewElems(nil)},
		{name: "NonEmptySequenceNoDefaultValue", en: NewElems(1, 2), want: NewElems(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.DefaultIfEmpty(); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.DefaultIfEmpty() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_DefaultIfEmptyDef(t *testing.T) {
	type args struct {
		def common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "EmptySequenceWithDefaultValue", en: enEmpty, args: args{def: 5}, want: NewElems(5)},
		{name: "EmptySequenceWithDefaultValue2", en: enEmptyEnmr, args: args{def: 5}, want: NewElems(5)},
		{name: "EmptySequenceWithDefaultValue3", en: enEmptySlice, args: args{def: 5}, want: NewElems(5)},
		{name: "NonEmptySequenceWithDefaultValue", en: NewElems(1, 2), args: args{def: 5}, want: NewElems(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.DefaultIfEmptyDef(tt.args.def); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.DefaultIfEmptyDef() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
