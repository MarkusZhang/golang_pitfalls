package quiz

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

func TestQuiz_4(t *testing.T) {
	err := getError()
	fmt.Println(err == nil)

	var err2 error
	err2 = getError()
	fmt.Println(err2 == nil)

	err3 := getError()
	fmt.Println(err3.(*APIError) == nil)
}

func getError() error {
	var err *APIError
	return err
}