package domain

import "github.com/flydevs/chat-app-api/common/server_message"

type Response struct {
	Response server_message.Svr_message `json:"data"`
	Data     interface{}                `json:"response"`
}

func (r Response) CreateResponse(data interface{}, resp server_message.Svr_message) Response {
	r.Data = data
	r.Response = resp
	return r
}

func (r Response) CreateJSON() {}
