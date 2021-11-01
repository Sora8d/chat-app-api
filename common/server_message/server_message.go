package server_message

import (
	"fmt"
	"net/http"
)

//This is just a placeholder for an actual message structure able to be usable across the apps
type Svr_message interface {
	GetStatus() int
	GetMessage() string
	GetFormatted() string
}

type svr_message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ae svr_message) GetStatus() int {
	return ae.Status
}

func (ae svr_message) GetMessage() string {
	return ae.Message
}

func (ae svr_message) GetFormatted() string {
	return fmt.Sprintf("Message %d, '%s'", ae.Status, ae.Message)
}

func NewInternalError() Svr_message {
	return svr_message{Status: http.StatusInternalServerError, Message: "Oops... Algo ha salido mal"}
}

func NewNotFoundError(message string) Svr_message {
	return svr_message{Status: http.StatusNotFound, Message: message}
}

func NewBadRequestError(message string) Svr_message {
	return svr_message{Status: http.StatusBadRequest, Message: message}
}

func NewCustomMessage(status int, message string) Svr_message {
	return svr_message{Status: status, Message: message}
}

//New types of errors should be added as they are needed
