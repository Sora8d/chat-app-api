package twilio

import (
	"net/http"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/clients/twilio"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	conversation "github.com/twilio/twilio-go/rest/conversations/v1"
)

type twioRepository struct {
}

type TwilioRepository interface {
	CreateConversation() (*string, server_message.Svr_message)
	JoinParticipant(string, string) server_message.Svr_message
	CreateMessage(string, message.Message) (*string, server_message.Svr_message)
}

func NewTwilioRepository() TwilioRepository {
	return &twioRepository{}
}

func (twio twioRepository) CreateConversation() (*string, server_message.Svr_message) {
	client := twilio.GetTwilioClient()
	params := conversation.CreateConversationParams{}
	resp, err := client.ConversationsV1.CreateConversation(&params)
	if err != nil {
		logger.Error("error in CreateConversation function in the twilio repository", err)
		return nil, server_message.NewInternalError()
	}
	return resp.Sid, server_message.NewCustomMessage(http.StatusOK, "")
}

func (twio twioRepository) JoinParticipant(conversation_identifier, participant_identifier string) server_message.Svr_message {
	client := twilio.GetTwilioClient()
	params := conversation.CreateConversationParticipantParams{
		Identity: &conversation_identifier,
	}
	_, err := client.ConversationsV1.CreateConversationParticipant(conversation_identifier, &params)
	if err != nil {
		logger.Error("error in CreateConversationParticipant function in the twilio repository", err)
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(http.StatusOK, "")
}

func (twio twioRepository) CreateMessage(conversation_identifier string, msg message.Message) (*string, server_message.Svr_message) {
	client := twilio.GetTwilioClient()
	params := conversation.CreateConversationMessageParams{
		Author: &msg.AuthorUuid,
		Body:   &msg.Text,
	}
	resp, err := client.ConversationsV1.CreateConversationMessage(conversation_identifier, &params)
	if err != nil {
		logger.Error("error in CreateMessage function in the twilio repository", err)
		return nil, server_message.NewInternalError()
	}
	return resp.Sid, server_message.NewCustomMessage(http.StatusOK, "")
}
