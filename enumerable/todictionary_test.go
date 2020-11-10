package enumerable

import (
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToDictionaryTest.cs

func TestEnumerable_ToDictionary(t *testing.T) {
	type args struct {
		ksel func(common.Elem) common.Elem
	}
	tests := []struct {
		name        string
		en          *Enumerable
		args        args
		want        Dictionary
		wantErr     bool
		expectedErr error
	}{
		{name: "NilKeySelectorNoComparerNoElementSelector",
			en:      enEmpty,
			args:    args{ksel: nil},
			wantErr: true, expectedErr: errors.NilSel},
		{name: "JustKeySelector",
			en:   NewElems("zero", "one", "two"),
			args: args{ksel: func(e common.Elem) common.Elem { return []rune(e.(string))[0] }},
			want: Dictionary{'z': "zero", 'o': "one", 't': "two"}},
		{name: "DuplicateKey",
			en:      NewElems("zero", "One", "Two", "three"),
			args:    args{ksel: func(e common.Elem) common.Elem { return strings.ToLower(string([]rune(e.(string))[:1])) }},
			wantErr: true, expectedErr: errors.DupKeys},
		{name: "NilEntryKeyCausesException",
			en:      NewElems("a", "b", nil),
			args:    args{ksel: common.Identity},
			wantErr: true, expectedErr: errors.NilKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.en.ToDictionary(tt.args.ksel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerable.ToDictionary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Enumerable.ToDictionary() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.ToDictionary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerable_ToDictionarySelMust(t *testing.T) {
	type args struct {
		ksel func(common.Elem) common.Elem
		vsel func(common.Elem) common.Elem
	}
	tests := []struct {
		name string
		en   *Enumerable
		args args
		want Dictionary
	}{
		{name: "KeyAndElementSelector",
			en: NewElems("zero", "one", "two"),
			args: args{
				ksel: func(e common.Elem) common.Elem { return []rune(e.(string))[0] },
				vsel: func(e common.Elem) common.Elem { return len(e.(string)) }},
			want: Dictionary{'z': 4, 'o': 3, 't': 3}},
		{name: "NilEntryValueIsAllowed",
			en: NewElems("a", "b"),
			args: args{
				ksel: common.Identity,
				vsel: func(e common.Elem) common.Elem { return nil }},
			want: Dictionary{"a": nil, "b": nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.en.ToDictionarySelMust(tt.args.ksel, tt.args.vsel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enumerable.ToDictionarySelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnum_CustomSelector(t *testing.T) {
	en := NewElems("zero", "one", "THREE")
	ksel := func(e common.Elem) common.Elem { return strings.ToLower(string([]rune(e.(string))[0])) }
	vsel := func(e common.Elem) common.Elem { return len(e.(string)) }
	got := en.ToDictionarySelMust(ksel, vsel)
	if len(got) != 3 {
		t.Errorf("len(NewElems.ToDictionarySelMust()) = %v, want 3", len(got))
	}
	want := Dictionary{"z": 4, "o": 3, "t": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewElems.ToDictionarySelMust() = '%v', want '%v'", got, want)
	}
}
