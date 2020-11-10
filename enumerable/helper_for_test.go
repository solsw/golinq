package enumerable

import (
	"strings"

	"github.com/solsw/golinq/common"
)

func enInt4() *Enumerable {
	return NewElems(1, 2, 3, 4)
}

func enStr4() *Enumerable {
	return NewElems("one", "two", "three", "four")
}

func enStr5() *Enumerable {
	return NewElems("one", "two", "three", "four", "five")
}

func enIntStr() *Enumerable {
	return NewElems(1, 2, "three", "four")
}

// emptyEnumerator implements the empty Enumerator interface.
type emptyEnumerator struct{}

// MoveNext implements the Enumerator.MoveNext method.
func (emptyEnumerator) MoveNext() bool {
	return false
}

// Current implements the Enumerator.Current method.
func (emptyEnumerator) Current() common.Elem {
	return nil
}

// Reset implements the Enumerator.Reset method.
func (emptyEnumerator) Reset() {}

var (
	enEmpty      = Empty()
	enEmptyEnmr  = &Enumerable{emptyEnumerator{}}
	enEmptySlice = NewElems()
)

type (
	elel struct {
		e1, e2 common.Elem
	}
	eee struct {
		e1, e2, e3 common.Elem
	}
	eeee struct {
		e1, e2, e3, e4 common.Elem
	}
)

var (
	// eqInt is a Equality for int.
	eqInt = func(e1, e2 common.Elem) bool {
		return e1.(int) == e2.(int)
	}

	// lsInt is a Less for int.
	lsInt = func(e1, e2 common.Elem) bool {
		return e1.(int) < e2.(int)
	}

	// cmpInt is a Comparison for int.
	cmpInt = common.LessToComparison(lsInt)

	// lsStr is a Less for string.
	lsStr = func(e1, e2 common.Elem) bool {
		return e1.(string) < e2.(string)
	}

	// lsNilStr is a Less for string, respecting nil parameters.
	lsNilStr = func(e1, e2 common.Elem) bool {
		if e2 == nil {
			return false
		}
		if e1 == nil {
			return true
		}
		return lsStr(e1, e2)
	}

	// lsCaseSensitive is case-sensitive Less for strings.
	lsCaseSensitive = func(e1, e2 common.Elem) bool {
		return e1.(string) < e2.(string)
	}

	// eqCaseInsensitive is case-insensitive Equality for strings.
	eqCaseInsensitive = func(e1, e2 common.Elem) bool {
		return strings.ToLower(e1.(string)) == strings.ToLower(e2.(string))
	}

	// lsCaseInsensitive is case-insensitive Less for strings.
	lsCaseInsensitive = func(e1, e2 common.Elem) bool {
		return strings.ToLower(e1.(string)) < strings.ToLower(e2.(string))
	}

	// cmpCaseInsensitive is case-insensitive Comparison for strings.
	cmpCaseInsensitive = func(e1, e2 common.Elem) int {
		s1 := strings.ToLower(e1.(string))
		s2 := strings.ToLower(e2.(string))
		if s1 < s2 {
			return -1
		}
		if s1 > s2 {
			return +1
		}
		return 0
	}
)

// enToStrings converts Enumerable to []string.
func enToStrings(en *Enumerable) []string {
	var r []string
	for en.MoveNext() {
		r = append(r, en.Current().(string))
	}
	return r
}
