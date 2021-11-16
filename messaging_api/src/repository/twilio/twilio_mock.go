package twilio

import (
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
)

type twiliomock struct {
}

func GetTwilioMock() TwilioRepository {
	return &twiliomock{}
}

func (twiom twiliomock) CreateConversation() (*string, server_message.Svr_message) {
	st := ""
	return &st, nil
}
func (twiom twiliomock) JoinParticipant(string, string) (*string, server_message.Svr_message) {
	st := ""
	return &st, nil
}
func (twiom twiliomock) CreateMessage(string, message.Message) (*string, server_message.Svr_message) {
	st := ""
	return &st, nil
}
func (twiom twiliomock) UpdateMessage(string, *message.Message) server_message.Svr_message {
	return nil
}
