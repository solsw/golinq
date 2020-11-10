package enumerable

import (
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ContainsTest.cs

func TestEnumerable_Contains(t *testing.T) {
	type args struct {
		el common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "NoMatchNoComparer",
			en:   NewElems("foo", "bar", "baz"),
			args: args{el: "BAR"},
			want: false},
		{name: "MatchNoComparer",
			en:   NewElems("foo", "bar", "baz"),
			args: args{el: strings.ToLower("BAR")},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.Contains(tt.args.el); got != tt.want {
				t.Errorf("Enumerable.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ContainsEq(t *testing.T) {
	type args struct {
		el common.Elem
		eq common.Equality
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want bool
	}{
		{name: "NoMatchWithCustomComparer",
			en:   NewElems("foo", "bar", "baz"),
			args: args{el: "gronk", eq: eqCaseInsensitive},
			want: false},
		{name: "MatchWithCustomComparer",
			en:   NewElems("foo", "bar", "baz"),
			args: args{el: "BAR", eq: eqCaseInsensitive},
			want: true},
		{name: "ImmediateReturnWhenMatchIsFound",
			en:   NewElems(10, 1, 5, 0),
			args: args{el: 2, eq: func(e1, e2 common.Elem) bool { return e1.(int) == 10/e2.(int) }},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ContainsEq(tt.args.el, tt.args.eq); got != tt.want {
				t.Errorf("Enumerable.ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
