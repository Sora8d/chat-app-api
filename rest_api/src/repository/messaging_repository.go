package repository

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/messaging_client"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/messaging"
)

type messagingRepository struct {
}

type MessagingRepositoryInterface interface {
}

func GetMessagingRepository() MessagingRepositoryInterface {
	return &messagingRepository{}
}

func (mr messagingRepository) CreateMessage(ctx context.Context, in *messaging.CreateMessageRequest) (*messaging.Uuid, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.CreateMessage(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Uuid, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) CreateConversation(ctx context.Context, in *messaging.Conversation) (*messaging.Uuid, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.CreateConversation(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Uuid, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) CreateUserConversation(ctx context.Context, in *messaging.CreateUserConversationRequest) server_message.Svr_message {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.CreateUserConversation(ctx, in)
	if err != nil {
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(int(response.Status), response.Message)
}

func (mr messagingRepository) GetConversationsByUser(ctx context.Context, in *messaging.Uuid) ([]*messaging.ConversationAndParticipants, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.GetConversationsByUser(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Conversations, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) GetMessagesByConversation(ctx context.Context, in *messaging.MessageRequest) ([]*messaging.Message, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.GetMessagesByConversation(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Message, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) UpdateMessage(ctx context.Context, in *messaging.Message) (*messaging.Message, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.UpdateMessage(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Message, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) UpdateConversationInfo(ctx context.Context, in *messaging.Conversation) (*messaging.Conversation, server_message.Svr_message) {
	client := messaging_client.GetMessagingClient()
	response, err := client.Client.UpdateConversationInfo(ctx, in)
	if err != nil {
		return nil, server_message.NewInternalError()
	}
	return response.Conversation, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}
