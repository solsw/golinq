package enumerable

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/EmptyTest.cs

func TestEmpty(t *testing.T) {
	tests := []struct {
		name string
		want *Enumerable
	}{
		{name: "EmptyContainsNoElements", want: enEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Empty() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
