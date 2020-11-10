package enumerable

import (
	"sort"

	"github.com/solsw/golinq/common"
)

func elInElelEq(el common.Elem, ee []common.Elem, eq common.Equality) bool {
	for _, e := range ee {
		if eq(e, el) {
			return true
		}
	}
	return false
}

func elIdxInElelCmp(el common.Elem, ee []common.Elem, cmp common.Comparison) int {
	return sort.Search(len(ee), func(i int) bool {
		return cmp(el, ee[i]) <= 0
	})
}

func elInElelCmp(el common.Elem, ee []common.Elem, cmp common.Comparison) bool {
	i := elIdxInElelCmp(el, ee, cmp)
	return i < len(ee) && cmp(el, ee[i]) == 0
}

func elIntoElelAtIdx(el common.Elem, ee *[]common.Elem, i int) {
	*ee = append(*ee, nil)
	if i < len(*ee)-1 {
		copy((*ee)[i+1:], (*ee)[i:])
	}
	(*ee)[i] = el
}

func projectionLess(ls common.Less, sel func(common.Elem) common.Elem) common.Less {
	return func(x, y common.Elem) bool {
		return ls(sel(x), sel(y))
	}
}

func reverseLess(ls common.Less) common.Less {
	return func(x, y common.Elem) bool {
		return ls(y, x)
	}
}

func compoundLess(ls1, ls2 common.Less) common.Less {
	return func(x, y common.Elem) bool {
		if ls1(x, y) {
			return true
		}
		if ls1(y, x) {
			return false
		}
		return ls2(x, y)
	}
}
