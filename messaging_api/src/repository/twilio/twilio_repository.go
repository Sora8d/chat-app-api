package twilio

import (
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
	JoinParticipant(string, string) (*string, server_message.Svr_message)
	CreateMessage(string, message.Message) (*string, server_message.Svr_message)
	UpdateMessage(string, *message.Message) server_message.Svr_message
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
	return resp.Sid, nil
}

func (twio twioRepository) JoinParticipant(conversation_identifier, participant_identifier string) (*string, server_message.Svr_message) {
	client := twilio.GetTwilioClient()
	params := conversation.CreateConversationParticipantParams{
		Identity: &participant_identifier,
	}
	resp, err := client.ConversationsV1.CreateConversationParticipant(conversation_identifier, &params)
	if err != nil {
		logger.Error("error in CreateConversationParticipant function in the twilio repository", err)
		return nil, server_message.NewInternalError()
	}
	return resp.Sid, nil
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
	return resp.Sid, nil
}

func (twio twioRepository) UpdateMessage(conversation_identifier string, msg *message.Message) server_message.Svr_message {
	client := twilio.GetTwilioClient()
	params := conversation.UpdateConversationMessageParams{
		Body: &msg.Text,
	}
	_, err := client.ConversationsV1.UpdateConversationMessage(conversation_identifier, msg.TwilioSid, &params)
	if err != nil {
		logger.Error("message updated in db, but not in twilio... error in UpdateMessage function in the twilio repository", err)
		return server_message.NewInternalError()
	}
	return nil
}

//This all will be sunject to change once i get a better idea of Tomas' plans
/*
func (twio twioRepository) DeleteConversation(conversation_identifier) {

}

func (twio twioRepository) DeleteParticipant(conversation_identifier, participant_identifier string) server_message.Svr_message {
	client := twilio.GetTwilioClient()
	if err := client.ConversationsV1.DeleteConversationParticipant(conversation_identifier, participant_identifier, nil); err != nil {

	}

}

func (twio twioRepository) DeleteMessage(conversation_identifier, message_identifier string) {
	client := twilio.GetTwilioClient()
	err := client.ConversationsV1.DeleteConversationMessage(conversation_identifier, message_identifier, nil)

}
*/
