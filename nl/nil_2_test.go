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

func handle_2() *APIError {
	var apierr *APIError
	err := doSomething()
	if err != nil {
		apierr = &APIError{
			Message: err.Error(),
		}
	}
	return apierr
}

func TestNil_2(t *testing.T) {
	err := handle_2()
	if err != nil {
		fmt.Println("error occurred")
	} else {
		fmt.Println("success")
	}
}
