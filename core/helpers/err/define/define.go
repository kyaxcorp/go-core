package define

import (
	"errors"
	"fmt"
)

type GeneralError struct {
	Code             int
	Err              error  // This is the original defined error!
	AdditionalErrMsg string // This is additional, when we have an existing error, but we want to complete it with another message
}

func (e *GeneralError) Error() string {
	// https://golang.org/pkg/fmt/
	defaultFormat := "code -> %v | message -> %v"
	if e.AdditionalErrMsg != "" {
		return fmt.Sprintf(defaultFormat+" | additional: %v", e.Code, e.Err, e.AdditionalErrMsg)
	} else {
		return fmt.Sprintf(defaultFormat, e.Code, e.Err)
	}
}

func (e *GeneralError) SetAdditionalError(additionalErrorMsg string) {
	e.AdditionalErrMsg = additionalErrorMsg
}

// Code Return the Error code
func Code(err error) int {
	if err == nil {
		// No error
		return 0
	}
	var generalError *GeneralError
	if errors.As(err, &generalError) {
		return generalError.Code
	} else {
		// It may be a different error?
		// or simply it's not one...
		return 0
	}
}

func Is(err error) bool {
	if err == nil {
		// No error
		return false
	}
	// TODO: maybe catch here other errors

	var generalError *GeneralError
	if errors.As(err, &generalError) {
		fmt.Println("General Error:", generalError.Err)
		return true
	} else {
		// There is no error
		return false
	}
}

// Err -> This is used for defining at runtime errors!
func Err(code int, message ...interface{}) *GeneralError {
	e := &GeneralError{
		Code: code,
		//Err:  errors.New(strings.Join(message, " -> ")),
		Err: errors.New(fmt.Sprint(message...)),
	}
	return e
}
