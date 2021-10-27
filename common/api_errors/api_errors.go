package api_errors

import (
	"errors"
	"fmt"
	"net/http"
)

//This is just a placeholder for an actual error structure able to be usable across the apps
type Api_error interface {
	GetError() error
	GetStatus() int
	GetMessage() string
	GetFormatted() string
}

type api_error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

func (ae api_error) GetError() error {
	return ae.Err
}

func (ae api_error) GetStatus() int {
	return ae.Status
}

func (ae api_error) GetMessage() string {
	return ae.Message
}

func (ae api_error) GetFormatted() string {
	return fmt.Sprintf("Error %d, '%s'", ae.Status, ae.Message)
}

func NewInternalError(message string, err error) Api_error {
	return api_error{Status: http.StatusInternalServerError, Message: message, Err: err}
}

func NewNotFoundError(message string) Api_error {
	return api_error{Status: http.StatusNotFound, Message: message, Err: errors.New("not found")}
}

func NewBadRequestError(message string) Api_error {
	return api_error{Status: http.StatusBadRequest, Message: message, Err: errors.New("bad request")}
}

//New types of errors should be added as they are needed
