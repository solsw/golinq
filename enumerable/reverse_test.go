package enumerable

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ReverseTest.cs

func TestEnumerable_Reverse(t *testing.T) {
	tests := []struct {
		name string
		en   *Enumerable
		want *Enumerable
	}{
		{name: "ReversedRange", en: RangeMust(5, 5), want: NewElems(9, 8, 7, 6, 5)},
		{name: "ReversedStrs", en: NewElems("one", "two"), want: NewElems("two", "one")},
		{name: "EmptyInput", en: enEmpty, want: enEmpty},
		{name: "1", en: NewElems("1"), want: NewElems("1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Reverse(); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.Reverse() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
