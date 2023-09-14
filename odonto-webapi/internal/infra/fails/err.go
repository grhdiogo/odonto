package fails

import (
	"fmt"
)

func checker(msg string, stck Stack) bool {
	if stck.Message() == msg {
		return true
	} else if stck.ErrStack() != nil {
		return checker(msg, stck.ErrStack())
	} else {
		return false
	}
}

// -------------------------------------------------------
// Stack interface
// -------------------------------------------------------
type Stack interface {
	Code() int
	Message() string
	Error() string
	ErrStack() Stack
}

// -------------------------------------------------------
// Stack
// -------------------------------------------------------

type errStack struct {
	Cod int    `json:"code"`
	Msg string `json:"message"`
	Err Stack  `json:"errorStack"`
}

func (e *errStack) sprintErrStack() string {
	if err0, ok := e.Err.(*errStack); ok {
		return fmt.Sprintf(`code: %d, message: %s, errorStack: [%s]`, e.Cod, e.Msg, err0.sprintErrStack())
	}
	return fmt.Sprintf(`code: %d, message: %s`, e.Cod, e.Msg)
}

func (e *errStack) Code() int {
	return e.Cod
}

func (e *errStack) Message() string {
	return e.Msg
}

func (e *errStack) Error() string {
	return e.sprintErrStack()
}

func (e *errStack) ErrStack() Stack {
	return e.Err
}

func (e *errStack) Contains(err error) bool {
	return checker(err.Error(), e)
}

func NewError(code int, msg string, stack Stack) Stack {
	if stack != nil {
		return &errStack{
			Cod: code,
			Msg: msg,
			Err: stack,
		}
	}
	return &errStack{
		Cod: code,
		Msg: msg,
		Err: nil,
	}
}
