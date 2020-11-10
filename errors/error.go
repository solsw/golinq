package errors

import (
	"fmt"
	"reflect"

	"github.com/solsw/golinq/errcode"
)

// Error represents golinq error.
type Error struct {
	Code int
	Msg  string
}

// Error implements the 'error' interface.
func (e Error) Error() string {
	r := fmt.Sprintf("(%d)", e.Code)
	if len(e.Msg) > 0 {
		r = r + " " + e.Msg
	}
	return r
}

// Error variables.
var (
	DupKeys      = Error{Code: errcode.DupKeys, Msg: "duplicate keys"}
	EmptyEnum    = Error{Code: errcode.EmptyEnum, Msg: "enumerable is empty"}
	IdxRange     = Error{Code: errcode.IdxRange, Msg: "index is out of range"}
	MultiElems   = Error{Code: errcode.MultiElems, Msg: "enumerable contains multiple elements"}
	MultiMatch   = Error{Code: errcode.MultiMatch, Msg: "enumerable contains multiple matching elements"}
	NegCount     = Error{Code: errcode.NegCount, Msg: "count is negative"}
	NilAcc       = Error{Code: errcode.NilAcc, Msg: "accumulator is nil"}
	NilAction    = Error{Code: errcode.NilAction, Msg: "action is nil"}
	NilCmp       = Error{Code: errcode.NilCmp, Msg: "comparison is nil"}
	NilEq        = Error{Code: errcode.NilEq, Msg: "equality is nil"}
	NilKey       = Error{Code: errcode.NilKey, Msg: "key is nil"}
	NilLess      = Error{Code: errcode.NilLess, Msg: "less is nil"}
	NilPred      = Error{Code: errcode.NilPred, Msg: "predicate is nil"}
	NilSel       = Error{Code: errcode.NilSel, Msg: "selector is nil"}
	NoMatch      = Error{Code: errcode.NoMatch, Msg: "enumerable contains no matching elements"}
	WrongStrtCnt = Error{Code: errcode.WrongStrtCnt, Msg: "start and/or count are out of range"}
)

// NotSameType denotes erroneous situation when elements have not same type.
func NotSameType(t1, t2 reflect.Type) Error {
	return Error{Code: errcode.NotSameType, Msg: fmt.Sprintf("elements have not same type: '%s', '%s'", t1, t2)}
}

// WrongType denotes erroneous situation when element has wrong type.
func WrongType(t1, t2 reflect.Type) Error {
	return Error{Code: errcode.WrongType, Msg: fmt.Sprintf("element has type '%s' instead of '%s'", t1, t2)}
}
