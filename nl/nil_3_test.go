package nl

import (
	"fmt"
	"testing"
)

type APIError struct {
	Message string
}

func (a *APIError) Error() string {
	return a.Message
}

func doSomething() error {
	return nil
}

func handle_3() *APIError {
	var apierr *APIError
	err := doSomething()
	if err != nil {
		apierr = &APIError{
			Message: err.Error(),
		}
	}
	return apierr
}

func TestNil_3(t *testing.T) {
	var err error
	err = handle_3()
	if err != nil {
		fmt.Println("error occurred")
	} else {
		fmt.Println("success")
	}
}
