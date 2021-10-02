package fwork

import (
	"fmt"
	"log"
)

type ApiError struct {
	Status  int    // http status code
	Code    string // file code
	Message string // message
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[%v:%v] %v", e.Status, e.Code, e.Message)
}

// FatalfIfError is equivalent to Fatalf if error exists
func FatalfIfError(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}
