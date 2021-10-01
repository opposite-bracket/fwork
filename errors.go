package fwork

import (
	"fmt"
)

type ApiError struct {
	Status  int    // http status code
	Code    string // file code
	Message string // message
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[%v:%v] %v", e.Status, e.Code, e.Message)
}
