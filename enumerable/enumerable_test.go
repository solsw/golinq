package enumerable

import (
	"reflect"
	"testing"

	"github.com/solsw/golinq/common"
)

func TestEnumerable_Slice(t *testing.T) {
	en3 := NewElems(1, 2, 3, 4)
	en3.MoveNext()
	tests := []struct {
		name string
		en   *Enumerable
		want common.Slice
	}{
		{name: "01", en: enEmpty, want: nil},
		{name: "02", en: enEmptyEnmr, want: nil},
		{name: "03", en: enEmptySlice, want: nil},
		{name: "1", en: NewElems(1), want: common.Slice{1}},
		{name: "2", en: NewElems("1", "2", "3", " 4"), want: common.Slice{"1", "2", "3", " 4"}},
		{name: "3", en: en3, want: common.Slice{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_String(t *testing.T) {
	en3 := NewElems(1, 2, 3, 4)
	en3.MoveNext()
	tests := []struct {
		name string
		en   *Enumerable
		want string
	}{
		{name: "01", en: enEmpty, want: ""},
		{name: "02", en: enEmptyEnmr, want: ""},
		{name: "03", en: enEmptySlice, want: ""},
		{name: "1", en: NewElems(1), want: "1"},
		{name: "2", en: NewElems("1", "2", "3", " 4"), want: "1 2 3  4"},
		{name: "3", en: en3, want: "2 3 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.String(); got != tt.want {
				t.Errorf("Enumerable.String() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
