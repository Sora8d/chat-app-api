package services

import (
	"net/http"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/uuids"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/twilio"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
)

type MessagingService interface {
	CreateConversation(conversation.Conversation) (*uuids.Uuid, server_message.Svr_message)
	CreateMessage(message.Message) (*uuids.Uuid, server_message.Svr_message)
	CreateUserConversation(conversation.UserConversation) (*uuids.Uuid, server_message.Svr_message)

	GetConversationsByUser(string) ([]conversation.ConversationResponse, server_message.Svr_message)
	GetConversationByUuid(string) (*conversation.Conversation, server_message.Svr_message)
	UpdateConversationInfo(string, conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message)

	GetMessagesByConversation(string, string) ([]message.Message, server_message.Svr_message)
	UpdateMessage(string, string) (*message.Message, server_message.Svr_message)
}
type messagingService struct {
	dbrepo      db.MessagingDBRepository
	proto_users users_client.UserProtoClientRepository
	twiorepo    twilio.TwilioRepository
}

func NewMessagingService(dbrepo db.MessagingDBRepository, users_client users_client.UserProtoClientRepository, twiorepo twilio.TwilioRepository) MessagingService {
	return &messagingService{dbrepo: dbrepo, proto_users: users_client, twiorepo: twiorepo}
}

func (ms *messagingService) CreateConversation(convo conversation.Conversation) (*uuids.Uuid, server_message.Svr_message) {
	//Twilio---
	sid, err := ms.twiorepo.CreateConversation()
	if err != nil {
		return nil, err
	}
	convo.TwilioSid = *sid
	//---
	uuid, err := ms.dbrepo.CreateConversation(convo)
	if err != nil {
		return nil, err
	}
	return uuid, server_message.NewCustomMessage(http.StatusOK, "conversation created")
}
func (ms *messagingService) CreateMessage(msg message.Message) (*uuids.Uuid, server_message.Svr_message) {
	user, response_msg := ms.proto_users.GetUser(msg.AuthorUuid)
	if response_msg.GetStatus() != 200 {
		return nil, response_msg
	}
	msg.AuthorId = user.Id

	convo, response_msg := ms.GetConversationByUuid(msg.ConversationUuid)
	if response_msg.GetStatus() != 200 {
		return nil, response_msg
	}
	msg.ConversationId = convo.Id
	//Twilio---
	sid, err := ms.twiorepo.CreateMessage(convo.TwilioSid, msg)
	if err != nil {
		return nil, err
	}
	msg.TwilioSid = *sid
	//---
	uuid, err := ms.dbrepo.CreateMessage(msg)
	if err != nil {
		return nil, err
	}
	//Updating Conversations last msgs from here.
	_, err = ms.dbrepo.UpdateConversationLastMsg(msg.ConversationUuid, uuid.Uuid)
	if err != nil {
		return nil, err
	}

	return uuid, server_message.NewCustomMessage(http.StatusOK, "message created")
}
func (ms *messagingService) CreateUserConversation(userconvo conversation.UserConversation) (*uuids.Uuid, server_message.Svr_message) {
	//Validation
	user, response_msg := ms.proto_users.GetUser(userconvo.UserUuid)
	if response_msg.GetStatus() != 200 {
		return nil, response_msg
	}
	userconvo.UserId = user.Id

	convo, response_msg := ms.GetConversationByUuid(userconvo.ConversationUuid)
	if response_msg.GetStatus() != 200 {
		return nil, response_msg
	}
	userconvo.ConversationId = convo.Id
	//---
	//Twilio---
	sid, err := ms.twiorepo.JoinParticipant(convo.TwilioSid, userconvo.UserUuid)
	if err != nil {
		return nil, err
	}
	userconvo.TwilioSid = *sid
	//---
	uuid, err := ms.dbrepo.CreateUserConversation(userconvo)
	if err != nil {
		return nil, err
	}
	return uuid, server_message.NewCustomMessage(http.StatusOK, "user_conversation created")
}

//

func (ms *messagingService) GetConversationsByUser(user_uuid string) ([]conversation.ConversationResponse, server_message.Svr_message) {
	conversations, err := ms.dbrepo.GetConversationsByUser(user_uuid)
	if err != nil {
		return nil, err
	}

	return conversations, server_message.NewCustomMessage(http.StatusOK, "conversations retrieved")
}

func (ms *messagingService) GetConversationByUuid(uuid string) (*conversation.Conversation, server_message.Svr_message) {
	conversation, err := ms.dbrepo.GetConversationByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return conversation, server_message.NewCustomMessage(http.StatusOK, "conversation retrieved")
}
func (ms *messagingService) UpdateConversationInfo(uuid string, conv_info conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message) {
	convo, err := ms.dbrepo.UpdateConversationInfo(uuid, conv_info)
	if err != nil {
		return nil, err
	}
	return convo, server_message.NewCustomMessage(http.StatusOK, "conversation info updated")

}

//

func (ms *messagingService) GetMessagesByConversation(uc_uuid string, convo_uuid string) ([]message.Message, server_message.Svr_message) {
	messages, err := ms.dbrepo.GetMessagesByConversation(convo_uuid)
	if err != nil {
		return nil, err
	}
	_, err = ms.dbrepo.UserConversationUpdateLastAccess(uc_uuid, messages[len(messages)-1].Uuid)
	if err != nil {
		return nil, err
	}

	return messages, server_message.NewCustomMessage(http.StatusOK, "messages retrieved")
}
func (ms *messagingService) UpdateMessage(uuid string, text string) (*message.Message, server_message.Svr_message) {
	message, err := ms.dbrepo.UpdateMessage(uuid, text)
	if err != nil {
		return nil, err
	}
	//Twilio
	convo, msg := ms.GetConversationByUuid(message.ConversationUuid)
	if msg.GetStatus() != 200 {
		return nil, msg
	}
	ms.twiorepo.UpdateMessage(convo.TwilioSid, message)
	//----

	message.ConversationUuid = ""
	return message, server_message.NewCustomMessage(http.StatusOK, "message updated")

}
