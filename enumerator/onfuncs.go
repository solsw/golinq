package enumerator

import (
	"github.com/solsw/golinq/common"
)

// OnFuncs is Enumerator implementation based on fields-functions.
type OnFuncs struct {
	MvNxt func() bool
	Crrnt func() common.Elem
	Rst   func()
}

// MoveNext implements the Enumerator.MoveNext method.
func (en OnFuncs) MoveNext() bool {
	if en.MvNxt == nil {
		return false
	}
	return en.MvNxt()
}

// Current implements the Enumerator.Current method.
func (en OnFuncs) Current() common.Elem {
	if en.Crrnt == nil {
		return nil
	}
	return en.Crrnt()
}

// Reset implements the Enumerator.Reset method.
func (en OnFuncs) Reset() {
	if en.Rst == nil {
		return
	}
	en.Rst()
}
