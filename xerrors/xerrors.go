package xerrors

import (
	"errors"
	"fmt"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New returns an error object for the code, msg.
func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d,msg=%v", e.Code, e.Msg)
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	return New(UnknownCode, err.Error())
}
