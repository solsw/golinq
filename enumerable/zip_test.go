package enumerable

import (
	"fmt"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ZipTest.cs

func TestEnumerable_ZipMust(t *testing.T) {
	type args struct {
		en2 *Enumerable
		sel func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "ShortFirst",
			en: NewElems("a", "b", "c"),
			args: args{
				en2: RangeMust(5, 10),
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%s:%d", e1.(string), e2.(int))
				}},
			want: NewElems("a:5", "b:6", "c:7")},
		{name: "ShortSecond",
			en: NewElems("a", "b", "c", "d", "e"),
			args: args{
				en2: RangeMust(5, 3),
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%s:%d", e1.(string), e2.(int))
				}},
			want: NewElems("a:5", "b:6", "c:7")},
		{name: "EqualLengthSequences",
			en: NewElems("a", "b", "c"),
			args: args{
				en2: RangeMust(5, 3),
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%s:%d", e1.(string), e2.(int))
				}},
			want: NewElems("a:5", "b:6", "c:7")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ZipMust(tt.args.en2, tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ZipMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ZipSelfMust(t *testing.T) {
	ee := NewElems("a", "b", "c", "d", "e")
	r1 := RepeatMust("q", 2)
	r2 := RangeMust(1, 4)
	type args struct {
		en2 *Enumerable
		sel func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want *Enumerable
	}{
		{name: "AdjacentElements",
			en: ee,
			args: args{
				en2: ee.Skip(1),
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%s%s", e1.(string), e2.(string))
				}},
			want: NewElems("ab", "bc", "cd", "de")},
		{name: "SameEnumerable1",
			en: r1,
			args: args{
				en2: r1,
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%s:%s", e1.(string), e2.(string))
				}},
			want: NewElems("q:q", "q:q")},
		{name: "SameEnumerable2",
			en: r2.Skip(2),
			args: args{
				en2: r2,
				sel: func(e1, e2 common.Elem) common.Elem {
					return fmt.Sprintf("%d:%d", e1.(int), e2.(int))
				}},
			want: NewElems("3:1", "4:2")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ZipSelfMust(tt.args.en2, tt.args.sel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.ZipSelfMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
