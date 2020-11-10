package errcode

// golinq error codes
const (
	OK int = -iota
	DupKeys
	EmptyEnum
	IdxRange
	MultiElems
	MultiMatch
	NegCount
	NilAcc
	NilAction
	NilCmp
	NilEq
	NilKey
	NilLess
	NilPred
	NilSel
	NoMatch
	NotSameType
	WrongStrtCnt
	WrongType
)
