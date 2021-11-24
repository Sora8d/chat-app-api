package services

import (
	"context"

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
	CreateMessage(context.Context, message.Message) (*uuids.Uuid, server_message.Svr_message)
	CreateUserConversation(context.Context, conversation.CreateUserConversationRequest) server_message.Svr_message

	GetConversationsByUser(string) (conversation.ConversationAndParticipantsSlice, server_message.Svr_message)
	GetConversationByUuid(string) (*conversation.Conversation, server_message.Svr_message)
	UpdateConversationInfo(string, conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message)

	GetMessagesByConversation(string, string) (message.MessageSlice, server_message.Svr_message)
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
	convo.SetTwilioSid(*sid)
	//---
	uuid, err := ms.dbrepo.CreateConversation(convo)
	if err != nil {
		return nil, err
	}
	return uuid, nil
}
func (ms *messagingService) CreateMessage(ctx context.Context, msg message.Message) (*uuids.Uuid, server_message.Svr_message) {
	//get a token

	user, err := ms.proto_users.GetUser(ctx, []string{msg.AuthorUuid})
	if err != nil {
		return nil, err
	}
	msg.SetAuthorId(user[0].Id)

	convo, err := ms.GetConversationByUuid(msg.ConversationUuid)
	if err != nil {
		return nil, err
	}
	msg.SetConversationId(convo.Id)
	//Twilio---
	sid, err := ms.twiorepo.CreateMessage(convo.TwilioSid, msg)
	if err != nil {
		return nil, err
	}
	msg.SetTwilioSid(*sid)
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

	conversationUuid := uuids.Uuid{Uuid: msg.ConversationUuid}
	return &conversationUuid, nil
}
func (ms *messagingService) CreateUserConversation(ctx context.Context, userconvo conversation.CreateUserConversationRequest) server_message.Svr_message {
	convo, err := ms.GetConversationByUuid(userconvo.Conversation.Uuid)
	if err != nil {
		return err
	}
	userconvo.SetConversation(*convo)

	uuids := userconvo.UserConversationSlice.GetUuidsStringSlice()
	users, err := ms.proto_users.GetUser(ctx, uuids)
	if err != nil {
		return err
	}
	userconvo.UserConversationSlice.ParseIds(users)
	for i, uc := range userconvo.UserConversationSlice {
		//Twilio---
		sid, err := ms.twiorepo.JoinParticipant(convo.TwilioSid, uc.UserUuid)
		if err != nil {
			return err
		}
		userconvo.UserConversationSlice[i].SetTwilioSid(*sid)
		//---
	}
	err = ms.dbrepo.CreateUserConversation(userconvo)
	if err != nil {
		return err
	}
	return nil
}

//

func (ms *messagingService) GetConversationsByUser(user_uuid string) (conversation.ConversationAndParticipantsSlice, server_message.Svr_message) {
	conversations, err := ms.dbrepo.GetConversationsByUser(user_uuid)
	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (ms *messagingService) GetConversationByUuid(uuid string) (*conversation.Conversation, server_message.Svr_message) {
	conversation, err := ms.dbrepo.GetConversationByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}
func (ms *messagingService) UpdateConversationInfo(uuid string, conv_info conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message) {
	convo, err := ms.dbrepo.UpdateConversationInfo(uuid, conv_info)
	if err != nil {
		return nil, err
	}
	return convo, nil

}

//

func (ms *messagingService) GetMessagesByConversation(uc_uuid string, convo_uuid string) (message.MessageSlice, server_message.Svr_message) {
	messages, err := ms.dbrepo.GetMessagesByConversation(convo_uuid)
	if err != nil {
		return nil, err
	}
	_, err = ms.dbrepo.UserConversationUpdateLastAccess(uc_uuid, messages[len(messages)-1].Uuid)
	if err != nil {
		if err.GetStatus() == 404 {
			return nil, server_message.NewBadRequestError("the given user doesnt form part of the conversation")
		}
		return nil, err
	}

	return messages, nil
}
func (ms *messagingService) UpdateMessage(uuid string, text string) (*message.Message, server_message.Svr_message) {
	message, err := ms.dbrepo.UpdateMessage(uuid, text)
	if err != nil {
		return nil, err
	}
	//Twilio
	convo, err := ms.GetConversationByUuid(message.ConversationUuid)
	if err != nil {
		return nil, err
	}
	ms.twiorepo.UpdateMessage(convo.TwilioSid, message)
	//----

	message.SetConversationUuid("")
	return message, nil

}
