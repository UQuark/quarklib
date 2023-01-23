package qerr

import (
	"fmt"
	"runtime"
)

type Error struct {
	code    string
	pattern string
	message string
	cause   error
	stack   []byte

	formatted bool
}

func NewError(code, pattern string) Error {
	return Error{
		code:    code,
		pattern: pattern,
	}
}

func (e Error) New() Error {
	runtime.Stack(e.stack, false)
	return e
}

func (e Error) Cause(err error) Error {
	e.cause = err
	runtime.Stack(e.stack, false)
	return e
}

func (e Error) Format(args ...any) Error {
	e.message = fmt.Sprintf(e.pattern, args...)
	if e.cause != nil {
		if len(e.message) > 0 {
			e.message += "\ncaused by: " + e.cause.Error()
		} else {
			e.message = e.cause.Error()
		}
	}
	e.formatted = true
	runtime.Stack(e.stack, false)
	return e
}

func (e Error) Error() string {
	if !e.formatted {
		e = e.Format()
	}
	return e.message
}

func (e Error) GetCode() string {
	return e.code
}
