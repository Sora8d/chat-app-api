package services

import (
	"net/http"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/uuids"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
)

type MessagingService interface {
	CreateConversation(conversation.Conversation) (*uuids.Uuid, server_message.Svr_message)
	CreateMessage(message.Message) (*uuids.Uuid, server_message.Svr_message)
	CreateUserConversation(conversation.UserConversation) (*uuids.Uuid, server_message.Svr_message)

	GetConversationsByUser(string) ([]conversation.Conversation, server_message.Svr_message)
	GetConversationByUuid(string) (*conversation.Conversation, server_message.Svr_message)
	UpdateConversationLastMsg(string, string) (*conversation.Conversation, server_message.Svr_message)
	UpdateConversationInfo(string, conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message)

	GetMessagesByAuthor(string) ([]message.Message, server_message.Svr_message)
	GetMessagesByConversation(string) ([]message.Message, server_message.Svr_message)
	GetMessageByUuid(string) (*message.Message, server_message.Svr_message)
	UpdateMessage(string, string) (*message.Message, server_message.Svr_message)

	GetUserConversationsForUser(string) ([]conversation.UserConversation, server_message.Svr_message)
	GetUserConversationsForConversation(string) ([]conversation.UserConversation, server_message.Svr_message)
	GetUserConversationByUuid(string) (*conversation.UserConversation, server_message.Svr_message)
	UserConversationUpdateLastAccess(string, string) (*conversation.UserConversation, server_message.Svr_message)
}
type messagingService struct {
	dbrepo      db.MessagingDBRepository
	proto_users users_client.UserProtoClientRepository
}

func NewMessagingService(dbrepo db.MessagingDBRepository, users_client users_client.UserProtoClientRepository) MessagingService {
	return &messagingService{dbrepo: dbrepo, proto_users: users_client}
}

func (ms *messagingService) CreateConversation(convo conversation.Conversation) (*uuids.Uuid, server_message.Svr_message) {
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

	uuid, err := ms.dbrepo.CreateMessage(msg)
	if err != nil {
		return nil, err
	}
	//Later decide whether conversation_last_message will be updated here or elsewhere
	return uuid, server_message.NewCustomMessage(http.StatusOK, "message created")
}
func (ms *messagingService) CreateUserConversation(userconvo conversation.UserConversation) (*uuids.Uuid, server_message.Svr_message) {
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

	uuid, err := ms.dbrepo.CreateUserConversation(userconvo)
	if err != nil {
		return nil, err
	}
	return uuid, server_message.NewCustomMessage(http.StatusOK, "user_conversation created")
}

//

func (ms *messagingService) GetConversationsByUser(user_uuid string) ([]conversation.Conversation, server_message.Svr_message) {
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
func (ms *messagingService) UpdateConversationLastMsg(uuid string, last_msg string) (*conversation.Conversation, server_message.Svr_message) {
	result, err := ms.dbrepo.UpdateConversationLastMsg(uuid, last_msg)
	if err != nil {
		return nil, err
	}
	return result, server_message.NewCustomMessage(http.StatusOK, "conversation updatee retrieved")
}

func (ms *messagingService) UpdateConversationInfo(uuid string, conv_info conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message) {
	convo, err := ms.dbrepo.UpdateConversationInfo(uuid, conv_info)
	if err != nil {
		return nil, err
	}
	return convo, server_message.NewCustomMessage(http.StatusOK, "conversation info updated")

}

//

func (ms *messagingService) GetMessageByUuid(uuid string) (*message.Message, server_message.Svr_message) {
	message, err := ms.dbrepo.GetMessageByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return message, server_message.NewCustomMessage(http.StatusOK, "message retrieved")
}
func (ms *messagingService) GetMessagesByAuthor(uuid string) ([]message.Message, server_message.Svr_message) {
	user, msg := ms.proto_users.GetUser(uuid)
	if msg.GetStatus() != 200 {
		return nil, msg
	}

	messages, err := ms.dbrepo.GetMessagesByAuthor(user.Id)
	if err != nil {
		return nil, err
	}
	return messages, server_message.NewCustomMessage(http.StatusOK, "messages retrieved")
}
func (ms *messagingService) GetMessagesByConversation(uuid string) ([]message.Message, server_message.Svr_message) {
	messages, err := ms.dbrepo.GetMessagesByConversation(uuid)
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
	return message, server_message.NewCustomMessage(http.StatusOK, "message updated")

}

//

func (ms *messagingService) GetUserConversationsForUser(uuid string) ([]conversation.UserConversation, server_message.Svr_message) {
	user, msg := ms.proto_users.GetUser(uuid)
	if msg.GetStatus() != 200 {
		return nil, msg
	}

	user_conversations, err := ms.dbrepo.GetUserConversationsForUser(user.Id)
	if err != nil {
		return nil, err
	}
	return user_conversations, server_message.NewCustomMessage(http.StatusOK, "user_conversation retrieved")
}
func (ms *messagingService) GetUserConversationsForConversation(uuid string) ([]conversation.UserConversation, server_message.Svr_message) {
	user_conversations, err := ms.dbrepo.GetUserConversationsForConversation(uuid)
	if err != nil {
		return nil, err
	}
	return user_conversations, server_message.NewCustomMessage(http.StatusOK, "user_conversation retrieved")
}
func (ms *messagingService) GetUserConversationByUuid(uuid string) (*conversation.UserConversation, server_message.Svr_message) {
	user_conversation, err := ms.dbrepo.GetUserConversationByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return user_conversation, server_message.NewCustomMessage(http.StatusOK, "user_conversation retrieved")

}
func (ms *messagingService) UserConversationUpdateLastAccess(uuid string, last_access_msg string) (*conversation.UserConversation, server_message.Svr_message) {
	result, err := ms.dbrepo.UserConversationUpdateLastAccess(uuid, last_access_msg)
	if err != nil {
		return nil, err
	}
	return result, server_message.NewCustomMessage(http.StatusOK, "user_conversation retrieved")
}
