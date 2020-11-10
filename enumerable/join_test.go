package enumerable

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/gohelpers/stringhelper"
	"github.com/solsw/golinq/common"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/JoinTest.cs

func TestEnumerable_JoinMust(t *testing.T) {
	type args struct {
		inner *Enumerable
		oksel func(common.Elem) common.Elem
		iksel func(common.Elem) common.Elem
		rsel  func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name  string
		outer *Enumerable
		args  args
		want  *Enumerable
	}{
		{name: "SimpleJoin",
			outer: NewElems("first", "second", "third"),
			args: args{
				inner: NewElems("essence", "offer", "eating", "psalm"),
				oksel: func(oel common.Elem) common.Elem { return ([]rune(oel.(string)))[0] },
				iksel: func(iel common.Elem) common.Elem { return ([]rune(iel.(string)))[1] },
				rsel:  func(oel, iel common.Elem) common.Elem { return oel.(string) + ":" + iel.(string) },
			},
			want: NewElems("first:offer", "second:essence", "second:psalm")},
		{name: "CustomComparer",
			outer: NewElems("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
			args: args{
				inner: NewElems("000abc", "111gHi", "222333"),
				oksel: func(oel common.Elem) common.Elem {
					s, _ := stringhelper.Substr(oel.(string), 0, 3)
					return strings.ToLower(s)
				},
				iksel: func(iel common.Elem) common.Elem {
					s, _ := stringhelper.SubstrEnd(iel.(string), 3)
					return strings.ToLower(s)
				},
				rsel: func(oel, iel common.Elem) common.Elem { return oel.(string) + ":" + iel.(string) },
			},
			want: NewElems("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.outer.JoinMust(tt.args.inner, tt.args.oksel, tt.args.iksel, tt.args.rsel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.JoinMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_JoinEqMust_CustomComparer(t *testing.T) {
	outer := NewElems("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewElems("000abc", "111gHi", "222333")
	got := outer.JoinEqMust(inner,
		func(oel common.Elem) common.Elem {
			s, _ := stringhelper.SubstrBeg(oel.(string), 3)
			return s
		},
		func(iel common.Elem) common.Elem {
			s, _ := stringhelper.SubstrEnd(iel.(string), 3)
			return s
		},
		func(oel, iel common.Elem) common.Elem { return oel.(string) + ":" + iel.(string) },
		eqCaseInsensitive)
	want := NewElems("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("JoinEqMust_CustomComparer = '%v', want '%v'", got, want)
	}
}

func Test_JoinMust_DifferentSourceTypes(t *testing.T) {
	outer := NewElems(5, 3, 7)
	inner := NewElems("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got := outer.JoinMust(inner,
		common.Identity,
		func(iel common.Elem) common.Elem { return len(iel.(string)) },
		func(oel, iel common.Elem) common.Elem { return fmt.Sprintf("%v:%v", oel, iel) })
	want := NewElems("5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("JoinMust_DifferentSourceTypes = '%v', want '%v'", got, want)
	}
}

func Test_JoinMust_NilKeys(t *testing.T) {
	outer := NewElems("first", "nil", "nothing", "second")
	inner := NewElems("nuff", "second")
	got := outer.JoinMust(inner,
		func(oel common.Elem) common.Elem {
			if strings.HasPrefix(oel.(string), "n") {
				return nil
			}
			return oel
		},
		func(iel common.Elem) common.Elem {
			if strings.HasPrefix(iel.(string), "n") {
				return nil
			}
			return iel
		},
		func(oel, iel common.Elem) common.Elem { return fmt.Sprintf("%v:%v", oel, iel) })
	want := NewElems("second:second")
	if !got.SequenceEqual(want) {
		got.Reset()
		want.Reset()
		t.Errorf("JoinMust_NilKeys = '%v', want '%v'", got, want)
	}
}

func TestEnumerable_JoinSelfMust(t *testing.T) {
	en := NewElems("fs", "sf", "ff", "ss")
	type args struct {
		inner *Enumerable
		oksel func(common.Elem) common.Elem
		iksel func(common.Elem) common.Elem
		rsel  func(common.Elem, common.Elem) common.Elem
	}
	tests := []struct {
		name  string
		outer *Enumerable
		args  args
		want  *Enumerable
	}{
		{name: "SameEnumerable",
			outer: en,
			args: args{
				inner: en,
				oksel: func(oel common.Elem) common.Elem { return ([]rune(oel.(string)))[0] },
				iksel: func(iel common.Elem) common.Elem { return ([]rune(iel.(string)))[1] },
				rsel:  func(oel, iel common.Elem) common.Elem { return oel.(string) + ":" + iel.(string) },
			},
			want: NewElems("fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.outer.JoinSelfMust(tt.args.inner, tt.args.oksel, tt.args.iksel, tt.args.rsel); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Enumerable.JoinSelfMust() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
