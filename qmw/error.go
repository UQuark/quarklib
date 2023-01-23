package qmw

import (
	"encoding/json"
	"fmt"
)

type statusError struct {
	status int
	e      error
}

func E(status int, e error) error {
	return statusError{status, e}
}

func (e statusError) Error() string {
	return fmt.Sprintf("%d %s", e.status, e.e.Error())
}

type ErrBody struct {
	Code    string
	Message string
}

func (e ErrBody) Json() []byte {
	b, _ := json.Marshal(e)
	return b
}
