package domain

import "github.com/Sora8d/common/server_message"

type Response struct {
	Response server_message.Svr_message `json:"response"`
	Data     interface{}                `json:"data"`
}

func (r Response) CreateResponse(data interface{}, resp server_message.Svr_message) Response {
	r.Data = data
	r.Response = resp
	return r
}

func (r Response) CreateJSON() {}
