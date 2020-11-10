package chain

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
)

func TestChain(t *testing.T) {
	type args struct {
		en    *enumerable.Enumerable
		links []LinkMust
	}
	tests := []struct {
		name string
		args args
		want *enumerable.Enumerable
	}{
		{name: "1",
			args: args{
				en: enumerable.NewElems(1, 2, 3, 4),
				links: []LinkMust{
					WhereLinkMust(func(e common.Elem) bool { return e.(int)%2 == 0 }),
					SelectLinkMust(func(e common.Elem) common.Elem { return e.(int) * 2 }),
				},
			},
			want: enumerable.NewElems(4, 8)},
		{name: "2",
			args: args{
				en: enumerable.NewElems("one", "two", "three", "four"),
				links: []LinkMust{
					WhereLinkMust(func(e common.Elem) bool { return strings.HasPrefix(e.(string), "t") }),
					SelectLinkMust(func(e common.Elem) common.Elem { return len(e.(string)) }),
				},
			},
			want: enumerable.NewElems(3, 5)},
		{name: "3",
			args: args{
				en: enumerable.NewElems("one", "two", "three", "four", "five"),
				links: []LinkMust{
					SelectLinkMust(func(e common.Elem) common.Elem { return fmt.Sprintf("%d%s", len(e.(string)), e.(string)) }),
					WhereLinkMust(func(e common.Elem) bool { return strings.HasPrefix(e.(string), "4") }),
				},
			},
			want: enumerable.NewElems("4four", "4five")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChainMust(tt.args.en, tt.args.links...); !got.SequenceEqual(tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Chain() = %v, want %v", got, tt.want)
			}
		})
	}
}
