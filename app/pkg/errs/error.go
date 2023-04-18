package errs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

import "github.com/gin-gonic/gin"

const (
	codePrefix    = "code: "
	separator     = ", "
	messagePrefix = "message: "
)

// Error type describes the internal representation of an error in the system
type Error struct {
	// Code contains code according to the error dictionary
	Code string `json:"code"`
	// Message text description of the error
	Message string `json:"message"`
}

// jsonError non-imported type for wrapping Error in JSON
type jsonError struct {
	Error Error `json:"error,omitempty"`
}

// New returns the declared type Error
func New() *Error {
	return &Error{}
}

// FromError accepts an error and tries to cast it to an internal type Error
func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if !strings.HasPrefix(err.Error(), codePrefix) {
		return &Error{
			Code:    Internal,
			Message: err.Error(),
		}
	}

	var code string
	var message string

	errComponents := strings.Split(err.Error(), separator)
	for i, component := range errComponents {
		switch i {
		case 0:
			code = component[len(codePrefix):]
		case 1:
			message = component[len(messagePrefix):]
		default:
			message += component
		}
	}

	return &Error{
		Code:    code,
		Message: message,
	}
}

// FromBody accepts a bytes.Buffer and converts the content to an internal type Error
func FromBody(body *bytes.Buffer) *Error {
	if body == nil {
		return nil
	}

	// response has not an API error ¯\_(ツ)_/¯
	if !strings.Contains(body.String(), `{"error":{"code"`) {
		return nil
	}

	arrByte := make([]byte, len(body.Bytes()))
	copy(arrByte, body.Bytes())

	var jsError jsonError
	err := json.Unmarshal(arrByte, &jsError)
	if err != nil {
		return &Error{
			Code:    Internal,
			Message: "unable get the error from the body",
		}
	}

	if len(jsError.Error.Code) == 0 {
		return nil
	}

	return &jsError.Error
}

// FromBytes accepts a []byte and converts the content to an internal type Error
func FromBytes(bb []byte) *Error {
	if bb == nil {
		return nil
	}

	// response has not an API error ¯\_(ツ)_/¯
	if !strings.Contains(string(bb), `{"error":{"code"`) {
		return nil
	}

	arrByte := make([]byte, len(bb))
	copy(arrByte, bb)

	var jsError jsonError
	err := json.Unmarshal(arrByte, &jsError)
	if err != nil {
		return &Error{
			Code:    Internal,
			Message: "unable get the error from the bytes",
		}
	}

	if len(jsError.Error.Code) == 0 {
		return nil
	}

	return &jsError.Error
}

// SetCode sets error code
func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

// SetMsg sets error description
func (e *Error) SetMsg(format string, v ...interface{}) *Error {
	e.Message = fmt.Sprintf(format, v...)
	return e
}

// Error converts a type to an error
func (e *Error) Error() error {
	return errors.New(e.String())
}

// String converts a type to a string
func (e *Error) String() string {
	return fmt.Sprintf("%s%s%s%s%s", codePrefix, e.Code, separator, messagePrefix, e.Message)
}

// Marshal converts a type to a json []byte
func (e *Error) Marshal() []byte {
	jsError := jsonError{Error: *e}
	js, err := json.Marshal(jsError)
	if err != nil {
		return nil
	}

	return js
}

// GinJSON converts a type to a gin shortcut for map[string]interface{} for transmission by transport
func (e *Error) GinJSON() gin.H {
	return gin.H{"error": *e}
}
